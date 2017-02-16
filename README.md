My Go Packages
==============

# pkg/test

## suite

### example usage

```go
import (
    . "github.com/GeorgeMac/pkg/test"
    "github.com/GeorgeMac/pkg/psql"
)

func TestMain(m *testing.M) {
    db, - := os.Open(os.Getenv("DATABASE_URL"))
    Suite(m, DB(db, Setup(Seed), Teardown(psql.TruncateAll)))
}
```

# pkg/psql

common Postgres related functions.

1. `psql.TruncateAll` truncates all Postgres tables apart from a table called `schema_migrations`
