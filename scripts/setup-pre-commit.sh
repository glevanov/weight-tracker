#!/bin/sh

set -eu

HOOK_PATH=".git/hooks/pre-commit"

cat > "$HOOK_PATH" <<'EOF'
#!/bin/sh

set -eu

make pre-commit
EOF

chmod +x "$HOOK_PATH"

echo "Installed pre-commit hook at $HOOK_PATH"
