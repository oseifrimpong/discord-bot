#!/bin/sh
# This is the first script of the entrypoint call sequence
# It is responsible for building the .env file for the Go Service
echo "Validating OS Env Existence..."
ENV=/run/secrets/DISCUSSION_BOT_ENV_FILE
if test -f "$FILE"; then
  source /run/secrets/DISCUSSION_BOT_ENV_FILE
fi

echo APP_PORT="$APP_PORT"

echo "Reading OS Environment Variables & Writing to .env"
printenv > .env

echo "Starting $@..."
exec "$@"