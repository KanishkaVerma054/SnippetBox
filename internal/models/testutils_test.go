package models

import (
	"database/sql"
	"os"
	"testing"
)

/*
	// 14.7 Integration testing: Test database setup and teardown
*/
func newTestDB(t *testing.T) *sql.DB {
	/*
		// 14.7 Integration testing: Test database setup and teardown

		// Establish a sql.DB connection pool for our test database. Because our
		// setup and teardown scripts contains multiple SQL statements, we need
		// to use the "multiStatements=true" parameter in our DSN. This instructs
		// our MySQL database driver to support executing multiple SQL statements
		// in one db.Exec() call.
	*/
	
	// db, err := sql.Open("mysql", "test_web:pass@/test_snippetbox?parseTime=true&multiStatements=true")
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// For running with docker without creating a container
	db, err := sql.Open("mysql", "test_web:pass@tcp(127.0.0.1:3306)/test_snippetbox?parseTime=true&multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}

	/*
		// 14.7 Integration testing: Test database setup and teardown

		// Read the setup SQL script from file and execute the statements.
	*/
	scripts, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(scripts))
	if err != nil {
		t.Fatal(err)
	}

	/*
		// 14.7 Integration testing: Test database setup and teardown

		// Use the t.Cleanup() to register a function *which will automatically be
		// called by Go when the current test (or sub-test) which calls newTestDB()
		// has finished*. In this function we read and execute the teardown script,
		// and close the database connection pool.
	*/
	t.Cleanup(func()  {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}
		db.Close()
	})

	/*
		// 14.7 Integration testing: Test database setup and teardown

		// Return the database connection pool.
	*/
	return db
}