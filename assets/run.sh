#!/bin/bash

# Load environment.
export RUN_MODE=${RUN_MODE:-dev}
export CONFIG_SERVER=${CONFIG_SERVER:-http://localhost:8082}
export CONFIG_BRANCH=${CONFIG_BRANCH:-master}

# Generate app.conf
echo "Generating app.conf..."
envsubst < /.setup/app.conf > $APP_DIR/conf/app.conf

# Start app-ms-profile
cd $APP_DIR && CGO_ENABLED=0 go build && ./app-ms-profile
echo "Service started."
