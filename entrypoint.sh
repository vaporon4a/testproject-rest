#!/bin/sh

# wait for postgres to start
sleep 5

# run migrations
./migrator -conn-string="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB" -migration-path="./migrations"

# run the app
./rest-api

# exit
exit 0