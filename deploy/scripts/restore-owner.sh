#!/usr/bin/env bash

echo '### Setting owner'
sudo chown weighttracker:weighttracker /opt/weight-tracker/weight-tracker
echo '### Setting permissions'
sudo chmod 755 /opt/weight-tracker
sudo chmod 755 /opt/weight-tracker/weight-tracker
echo '### Restoring SE Linux labels'
sudo restorecon -RFv /opt/weight-tracker
echo '### Done'
