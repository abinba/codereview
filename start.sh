#!/bin/sh

# Download atlas script
curl -sSf https://atlasgo.sh > atlas.sh
chmod +x ./atlas.sh
./atlas.sh

# Debugging: Print environment variables
echo "DB_USER: '$DB_USER'"
echo "DB_PASSWORD: '$DB_PASSWORD'"
echo "DB_HOST: '$DB_HOST'"
echo "DB_PORT: '$DB_PORT'"
echo "DB_NAME: '$DB_NAME'"

# Construct and print the URL
DB_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"
echo "DB_URL: '$DB_URL'"

# Apply migrations
atlas migrate hash
atlas migrate apply -u "$DB_URL"

# Start the main application
./main