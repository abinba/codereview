#!/bin/sh

curl -sSf https://atlasgo.sh > atlas.sh
./atlas.sh

atlas migrate apply -u "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"

./main