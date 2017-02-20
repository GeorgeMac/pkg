// test is a set of useful test suite-like tools for Go
// It is un-opinionated and leverages/integrates with the *testing.M and TestMain function
// to provide common setup / teardown features for things like sql databases.
//
//  import (
//      "database/sql"
//
//      . "github.com/GeorgeMac/pkg/test"
//      . "github.com/GeorgeMac/pkg/test/db"
//      "github.com/GeorgeMac/pkg/psql"
//      _ "github.com/lib/pq"
//  )
//
//  func TestMain(m *testing.M) {
//      db, - := sql.Open("postgres", os.Getenv("DATABASE_URL"))
//      Suite(m, Setup(DB(db, Seed)), Teardown(DB(db, psql.TruncateAll)))
//  }
package test
