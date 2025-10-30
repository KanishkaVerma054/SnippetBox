package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	/*
		4.5 Designing a database model: Using the SnippetModel

		// Import the models package that we just created. You need to prefix this with
		// whatever module path you set up back in chapter 02.01 (Project Setup and Creating
		// a Module) so that the import statement looks like this:
		// "{your-module-path}/internal/models". If you can't remember what module path you
		// used, you can find it at the top of the go.mod file.
	*/
	"KanishkaVerma054/snipperBox.dev/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

/*
	4.5 Designing a database model: Using the SnippetModel

	// Add a snippets field to the application struct. This will allow us to
	// make the SnippetModel object available to our handlers.
*/
type application struct {
	errorLog	*log.Logger
	infoLog		*log.Logger
	snippets *models.SnippetModel
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP newtwork address")

	/*
		// 4.4 Creating a database connection pool

		// Define the command-line flag for mysql DSN string
	*/
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	/*
		// 4.4 Creating a database connection pool

		// To keep the main() function tidy I've put the code for creating a connection
		// pool into the separate openDB() function below. We pass openDB() the DSN
		// from the command-line flag.
	*/
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	/*
		//  4.4 Creating a database connection pool

		// We also defer a call to db.Close(), so that the connection pool is closed
		// before the main() function exits.
	*/
	defer db.Close()

	/*
		// 4.5 Designing a database model: Using the SnippetModel

		// Initialize a models.SnippetModel instance and add it to the application
		// dependencies.
	*/
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:	*addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Staring server on %s", *addr)

	// Call the ListenAndServe() method on our new http.Server struct.
	// err := srv.ListenAndServe()

	/*
		// 4.4 Creating a database connection pool

		// Because the err variable is now already declared in the code above, we need
		// to use the assignment operator = here, instead of the := 'declare and assign'
		// operator.
	*/
	err = srv.ListenAndServe()

	errorLog.Fatal(err)
}

/*
	// 4.4 Creating a database connection pool

	// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
	// for a given DSN.
*/
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}