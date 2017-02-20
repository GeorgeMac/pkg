My Go Packages
==============

# pkg/test

## suite

### example usage

```go
import (
    "database/sql"

    . "github.com/GeorgeMac/pkg/test"
    . "github.com/georgemac/pkg/test/db"
    "github.com/GeorgeMac/pkg/psql"
    _ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
    db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))

    // test package exposes a DSL for performing actions
    // before and after tests.
    Suite(m, Setup(DB(db, Seed)), Teardown(DB(db, psql.TruncateAll)))
}
```

# pkg/psql

common Postgres related functions.

1. `psql.TruncateAll` truncates all Postgres tables apart from a table called `schema_migrations`
