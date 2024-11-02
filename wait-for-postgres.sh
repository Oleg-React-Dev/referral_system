#!/bin/sh
set -e

host="$1"
shift
cmd="$@"


until PGPASSWORD=$PGSQL_USERS_PASSWORD psql -h "$host" -U "$PGSQL_USERS_USER" -c '\q'; do
  >&2 echo "Postgres not available - waiting"
  sleep 1
done

>&2 echo "Postgres is available - run the command"
exec $cmd
