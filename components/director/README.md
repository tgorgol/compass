# Director

The Director exposes GraphQL API.

## Development

After you introduce changes in the GraphQL schema, run the `gqlgen.sh` script.

To run Director with PostgreSQL container on local machine with latest DB schema, run the following command:
```bash
./run.sh
```

The GraphQL API playground is available at `localhost:3000`.

## Configuration

The Director binary allows to override some configuration parameters. You can specify following environment variables.

| ENV                         | Default        | Description                                       |
|-----------------------------|----------------|---------------------------------------------------|
| APP_ADDRESS                 | 127.0.0.1:3000 | The address and port for the service to listen on |
| APP_DB_USER                 | postgres       | Database username                                 |
| APP_DB_PASSWORD             | pgsql@12345    | Database password                                 |
| APP_DB_HOST                 | localhost      | Database host                                     |
| APP_DB_PORT                 | 5432           | Database port                                     |
| APP_DB_NAME                 | postgres       | Database name                                     |
| APP_DB_SSL                  | disable        | Database SSL mode (disable / enable)              |
| APP_API_ENDPOINT            | /graphql       | The endpoint for GraphQL API                      |
| APP_PLAYGROUND_API_ENDPOINT | /graphql       | The endpoint of GraphQL API for the Playground    |
