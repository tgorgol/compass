#!/usr/bin/env bash

# This script is responsible for validating if migrations scripts are correct.
# It starts Postgres, executes UP and DOWN migrations.
# This script requires `compass-schema-migrator` Docker image.

RED='\033[0;31m'
GREEN='\033[0;32m'
INVERTED='\033[7m'
NC='\033[0m' # No Color

set -e

readonly SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

function cleanup() {
    echo -e "${GREEN}Cleanup Postgres container and network${NC}"
    docker-compose down
}

trap cleanup EXIT

cd "${SCRIPT_DIR}/../.."

echo -e "${GREEN}Start Postgres in detached mode${NC}"
docker-compose up -d database

echo -e "${GREEN}Run UP migrations ${NC}"
docker-compose run schema-migrator

echo -e "${GREEN}Show schema_migrations table after UP migrations${NC}"
docker exec compass_database_1 psql -U usr compass -c "select * from schema_migrations"

echo -e "${GREEN}Run DOWN migrations ${NC}"
DB_SCHEMA_MIGRATION_DIRECTION=down docker-compose run schema-migrator

echo -e "${GREEN}Show schema_migrations table after DOWN migrations${NC}"
docker exec compass_database_1 psql -U usr compass -c "select * from schema_migrations"