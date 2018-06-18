#!/bin/bash

# Load environment.
export RUN_MODE=${RUN_MODE:-dev}
export CONFIG_SERVER=${CONFIG_SERVER:-http://localhost:8082}
export CONFIG_BRANCH=${CONFIG_BRANCH:-master}

# Generate app.conf
if [ ! -f /conf/app.conf ]; then
	echo "Generating app.conf..."
	envsubst < /.setup/conf/app.conf > /conf/app.conf
fi

# Start app-ms-profile
go build && ./main
echo "Service started."
