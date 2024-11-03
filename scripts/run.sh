#!/usr/bin/env sh

dbmate \
    -e "STORAGE_DSN" \
    --migrations-dir /app/migrations \
    --no-dump-schema up

exec /app/bin/svc "$1"