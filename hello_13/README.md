go mod init hello_13

https://github.com/jackc/pgx/wiki/Getting-started-with-pgx

The *pgx.Conn returned by pgx.Connect() represents a single connection and is not concurrency safe.  This is entirely appropriate for a simple command line example such as above. However, for many uses, such as a web application server, concurrency is required. To use a connection pool replace the import github.com/jackc/pgx/v4 with github.com/jackc/pgx/v4/pgxpool and connect with pgxpool.Connect() instead of pgx.Connect().

https://github.com/jackc/pgx/wiki/Numeric-and-decimal-support

The Go language does not have a standard decimal type. pgx supports the PostgreSQL numeric type out of the box. However, in the absence of a proper Go type it can only be used when translated to a float64 or string. This is obviously not ideal.

The recommended solution is to use the github.com/shopspring/decimal package. pgx has support for integrating with this package, but to avoid a mandatory external dependency the integration is in a separate package.

https://godoc.org/github.com/jackc/pgx/pgxpool

Establishing a Connection
The primary way of establishing a connection is with `pgxpool.Connect`.

```go
pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
```

The database connection string can be in URL or DSN format. PostgreSQL settings, pgx settings, and pool settings can be specified here. In addition, a config struct can be created by `ParseConfig` and modified before establishing the connection with `ConnectConfig`.

```go
config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
if err != nil {
    // ...
}
config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
    // do something with every new connection
}

pool, err := pgxpool.ConnectConfig(context.Background(), config)
```