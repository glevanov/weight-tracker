package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	defaultMongoDB            = "weight-tracker"
	defaultUsersCollection    = "users"
	defaultWeightsCollection  = "weight"
	defaultBatchSize          = 1000
	maxLoggedValidationErrors = 20
)

type migrationConfig struct {
	mongoURI          string
	mongoDB           string
	pgURI             string
	usersCollection   string
	weightsCollection string
	truncate          bool
	dryRun            bool
	batchSize         int
}

type collectionStats struct {
	scanned  int
	valid    int
	imported int
	skipped  int
}

type userRow struct {
	username string
	password string
	salt     string
}

type weightRow struct {
	username  string
	weight    float64
	timestamp time.Time
}

func main() {
	cfg := loadConfig()

	if cfg.batchSize < 1 {
		log.Fatalf("invalid --batch-size=%d, must be >= 1", cfg.batchSize)
	}

	ctx := context.Background()

	mongoClient, err := connectMongo(ctx, cfg.mongoURI)
	if err != nil {
		log.Fatalf("mongo connection failed: %v", err)
	}
	defer func() {
		disconnectCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := mongoClient.Disconnect(disconnectCtx); err != nil {
			log.Printf("warning: mongo disconnect failed: %v", err)
		}
	}()

	pgPool, err := connectPostgres(ctx, cfg.pgURI)
	if err != nil {
		log.Fatalf("postgres connection failed: %v", err)
	}
	defer pgPool.Close()

	if cfg.truncate && !cfg.dryRun {
		if err := truncateTables(ctx, pgPool); err != nil {
			log.Fatalf("failed to truncate postgres tables: %v", err)
		}
		log.Print("truncated postgres tables users and weights")
	}

	db := mongoClient.Database(cfg.mongoDB)

	usersStats, userErrors, err := migrateUsers(ctx, db.Collection(cfg.usersCollection), pgPool, cfg)
	if err != nil {
		log.Fatalf("users migration failed: %v", err)
	}

	weightsStats, weightErrors, err := migrateWeights(ctx, db.Collection(cfg.weightsCollection), pgPool, cfg)
	if err != nil {
		log.Fatalf("weights migration failed: %v", err)
	}

	printSummary(cfg, usersStats, weightsStats, userErrors, weightErrors)
}

func loadConfig() migrationConfig {
	mongoURI := getEnv("MONGO_URI", "")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is required")
	}

	pgURI := getEnv("PG_URI", getEnv("DATABASE_URL", ""))
	if pgURI == "" {
		log.Fatal("PG_URI or DATABASE_URL is required")
	}

	mongoDBDefault := getEnv("MONGO_DB", defaultMongoDB)
	usersCollectionDefault := getEnv("MONGO_USERS_COLLECTION", defaultUsersCollection)
	weightsCollectionDefault := getEnv("MONGO_WEIGHTS_COLLECTION", defaultWeightsCollection)

	cfg := migrationConfig{}
	flag.StringVar(&cfg.mongoURI, "mongo-uri", mongoURI, "MongoDB connection URI (defaults to MONGO_URI)")
	flag.StringVar(&cfg.mongoDB, "mongo-db", mongoDBDefault, "MongoDB database name")
	flag.StringVar(&cfg.pgURI, "pg-uri", pgURI, "PostgreSQL connection URI (defaults to PG_URI or DATABASE_URL)")
	flag.StringVar(&cfg.usersCollection, "users-collection", usersCollectionDefault, "Mongo users collection name")
	flag.StringVar(&cfg.weightsCollection, "weights-collection", weightsCollectionDefault, "Mongo weights collection name")
	flag.BoolVar(&cfg.truncate, "truncate", true, "Truncate users and weights in PostgreSQL before import")
	flag.BoolVar(&cfg.dryRun, "dry-run", false, "Validate and count records without writing to PostgreSQL")
	flag.IntVar(&cfg.batchSize, "batch-size", defaultBatchSize, "Insert/upsert batch size")
	flag.Parse()

	return cfg
}

func connectMongo(ctx context.Context, uri string) (*mongo.Client, error) {
	connCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(connCtx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	pingCtx, pingCancel := context.WithTimeout(ctx, 10*time.Second)
	defer pingCancel()

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, err
	}

	return client, nil
}

func connectPostgres(ctx context.Context, uri string) (*pgxpool.Pool, error) {
	connCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	pool, err := pgxpool.New(connCtx, uri)
	if err != nil {
		return nil, err
	}

	pingCtx, pingCancel := context.WithTimeout(ctx, 10*time.Second)
	defer pingCancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func truncateTables(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "TRUNCATE TABLE weights, users RESTART IDENTITY")
	return err
}

func migrateUsers(
	ctx context.Context,
	collection *mongo.Collection,
	pool *pgxpool.Pool,
	cfg migrationConfig,
) (collectionStats, []string, error) {
	stats := collectionStats{}
	validationErrors := make([]string, 0)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return stats, validationErrors, err
	}
	defer cursor.Close(ctx)

	buffer := make([]userRow, 0, cfg.batchSize)

	for cursor.Next(ctx) {
		stats.scanned++

		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			stats.skipped++
			validationErrors = appendValidationError(validationErrors, fmt.Sprintf("users doc #%d decode failed: %v", stats.scanned, err))
			continue
		}

		username, err := getRequiredString(doc, "username")
		if err != nil {
			stats.skipped++
			validationErrors = appendValidationError(validationErrors, fmt.Sprintf("users doc #%d invalid username: %v", stats.scanned, err))
			continue
		}

		password, err := getRequiredString(doc, "password")
		if err != nil {
			stats.skipped++
			validationErrors = appendValidationError(validationErrors, fmt.Sprintf("users doc #%d (%s) invalid password: %v", stats.scanned, username, err))
			continue
		}

		salt, err := getRequiredString(doc, "salt")
		if err != nil {
			stats.skipped++
			validationErrors = appendValidationError(validationErrors, fmt.Sprintf("users doc #%d (%s) invalid salt: %v", stats.scanned, username, err))
			continue
		}

		stats.valid++
		buffer = append(buffer, userRow{username: username, password: password, salt: salt})

		if len(buffer) >= cfg.batchSize {
			if err := flushUsers(ctx, pool, buffer, cfg.dryRun); err != nil {
				return stats, validationErrors, err
			}
			stats.imported += len(buffer)
			buffer = buffer[:0]
		}
	}

	if err := cursor.Err(); err != nil {
		return stats, validationErrors, err
	}

	if len(buffer) > 0 {
		if err := flushUsers(ctx, pool, buffer, cfg.dryRun); err != nil {
			return stats, validationErrors, err
		}
		stats.imported += len(buffer)
	}

	return stats, validationErrors, nil
}

func flushUsers(ctx context.Context, pool *pgxpool.Pool, rows []userRow, dryRun bool) error {
	if dryRun || len(rows) == 0 {
		return nil
	}

	batch := &pgx.Batch{}
	for _, row := range rows {
		batch.Queue(
			"INSERT INTO users (username, password, salt) VALUES ($1, $2, $3) ON CONFLICT (username) DO UPDATE SET password = EXCLUDED.password, salt = EXCLUDED.salt",
			row.username,
			row.password,
			row.salt,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range rows {
		if _, err := br.Exec(); err != nil {
			return err
		}
	}

	return nil
}

func migrateWeights(
	ctx context.Context,
	collection *mongo.Collection,
	pool *pgxpool.Pool,
	cfg migrationConfig,
) (collectionStats, []string, error) {
	stats := collectionStats{}
	validationErrors := make([]string, 0)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return stats, validationErrors, err
	}
	defer cursor.Close(ctx)

	buffer := make([]weightRow, 0, cfg.batchSize)

	for cursor.Next(ctx) {
		stats.scanned++

		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			stats.skipped++
			validationErrors = appendValidationError(validationErrors, fmt.Sprintf("weight doc #%d decode failed: %v", stats.scanned, err))
			continue
		}

		username, err := getUsername(doc)
		if err != nil {
			stats.skipped++
			validationErrors = appendValidationError(validationErrors, fmt.Sprintf("weight doc #%d invalid user: %v", stats.scanned, err))
			continue
		}

		weightValue, err := getRequiredNumber(doc, "weight")
		if err != nil {
			stats.skipped++
			validationErrors = appendValidationError(validationErrors, fmt.Sprintf("weight doc #%d (%s) invalid weight: %v", stats.scanned, username, err))
			continue
		}

		timestamp, err := getRequiredTime(doc, "timestamp")
		if err != nil {
			stats.skipped++
			validationErrors = appendValidationError(validationErrors, fmt.Sprintf("weight doc #%d (%s) invalid timestamp: %v", stats.scanned, username, err))
			continue
		}

		stats.valid++
		buffer = append(buffer, weightRow{username: username, weight: weightValue, timestamp: timestamp})

		if len(buffer) >= cfg.batchSize {
			if err := flushWeights(ctx, pool, buffer, cfg.dryRun); err != nil {
				return stats, validationErrors, err
			}
			stats.imported += len(buffer)
			buffer = buffer[:0]
		}
	}

	if err := cursor.Err(); err != nil {
		return stats, validationErrors, err
	}

	if len(buffer) > 0 {
		if err := flushWeights(ctx, pool, buffer, cfg.dryRun); err != nil {
			return stats, validationErrors, err
		}
		stats.imported += len(buffer)
	}

	return stats, validationErrors, nil
}

func flushWeights(ctx context.Context, pool *pgxpool.Pool, rows []weightRow, dryRun bool) error {
	if dryRun || len(rows) == 0 {
		return nil
	}

	batch := &pgx.Batch{}
	for _, row := range rows {
		batch.Queue(
			"INSERT INTO weights (username, weight, timestamp) VALUES ($1, $2, $3)",
			row.username,
			row.weight,
			row.timestamp,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range rows {
		if _, err := br.Exec(); err != nil {
			return err
		}
	}

	return nil
}

func getUsername(doc bson.M) (string, error) {
	if value, ok := doc["user"]; ok {
		username, err := toNonEmptyString(value)
		if err == nil {
			return username, nil
		}
	}

	if value, ok := doc["username"]; ok {
		username, err := toNonEmptyString(value)
		if err == nil {
			return username, nil
		}
	}

	return "", errors.New("missing or empty user/username")
}

func getRequiredString(doc bson.M, key string) (string, error) {
	value, ok := doc[key]
	if !ok {
		return "", errors.New("field is missing")
	}

	return toNonEmptyString(value)
}

func toNonEmptyString(value any) (string, error) {
	s, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("expected string but got %T", value)
	}

	s = strings.TrimSpace(s)
	if s == "" {
		return "", errors.New("value is empty")
	}

	return s, nil
}

func getRequiredNumber(doc bson.M, key string) (float64, error) {
	value, ok := doc[key]
	if !ok {
		return 0, errors.New("field is missing")
	}

	return toFloat64(value)
}

func toFloat64(value any) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case primitive.Decimal128:
		f, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return 0, fmt.Errorf("invalid decimal128 value: %w", err)
		}
		return f, nil
	default:
		return 0, fmt.Errorf("expected numeric value but got %T", value)
	}
}

func getRequiredTime(doc bson.M, key string) (time.Time, error) {
	value, ok := doc[key]
	if !ok {
		return time.Time{}, errors.New("field is missing")
	}

	return toTime(value)
}

func toTime(value any) (time.Time, error) {
	switch v := value.(type) {
	case time.Time:
		return v, nil
	case primitive.DateTime:
		return v.Time(), nil
	case string:
		t, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return time.Time{}, fmt.Errorf("invalid timestamp string: %w", err)
		}
		return t, nil
	default:
		return time.Time{}, fmt.Errorf("expected time/date but got %T", value)
	}
}

func appendValidationError(errors []string, message string) []string {
	if len(errors) < maxLoggedValidationErrors {
		return append(errors, message)
	}
	return errors
}

func printSummary(
	cfg migrationConfig,
	usersStats collectionStats,
	weightsStats collectionStats,
	userErrors []string,
	weightErrors []string,
) {
	action := "imported"
	if cfg.dryRun {
		action = "would import"
	}

	fmt.Printf("\nMongo -> Postgres migration complete (%s)\n", action)
	fmt.Printf("- Mongo DB: %s\n", cfg.mongoDB)
	fmt.Printf("- Users collection: %s\n", cfg.usersCollection)
	fmt.Printf("- Weights collection: %s\n", cfg.weightsCollection)
	fmt.Printf("- Truncate enabled: %t\n", cfg.truncate)
	fmt.Printf("- Dry run: %t\n\n", cfg.dryRun)

	printCollectionStats("users", usersStats, cfg.dryRun)
	printCollectionStats("weights", weightsStats, cfg.dryRun)

	printValidationErrors("users", usersStats, userErrors)
	printValidationErrors("weights", weightsStats, weightErrors)
}

func printCollectionStats(name string, stats collectionStats, dryRun bool) {
	actionLabel := "Imported"
	if dryRun {
		actionLabel = "Would import"
	}

	fmt.Printf("%s:\n", strings.ToUpper(name))
	fmt.Printf("- Scanned: %d\n", stats.scanned)
	fmt.Printf("- Valid: %d\n", stats.valid)
	fmt.Printf("- %s: %d\n", actionLabel, stats.imported)
	fmt.Printf("- Skipped: %d\n\n", stats.skipped)
}

func printValidationErrors(name string, stats collectionStats, validationErrors []string) {
	if stats.skipped == 0 {
		return
	}

	fmt.Printf("%s validation errors (showing up to %d):\n", strings.ToUpper(name), maxLoggedValidationErrors)
	for _, issue := range validationErrors {
		fmt.Printf("- %s\n", issue)
	}
	if stats.skipped > len(validationErrors) {
		fmt.Printf("- ... and %d more\n", stats.skipped-len(validationErrors))
	}
	fmt.Println()
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}
