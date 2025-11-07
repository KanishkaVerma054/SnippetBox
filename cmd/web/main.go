package main

import (
	"KanishkaVerma054/snipperBox.dev/internal/models"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/go-playground/form"
	_ "github.com/go-sql-driver/mysql"
	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
)

/*
	// 9.2 Setting up the session manager

	// Add a new sessionManager field to the application struct.
*/
type application struct {
	errorLog		*log.Logger
	infoLog			*log.Logger
	snippets 		*models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder		*form.Decoder
	sessionManager  *scs.SessionManager
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP newtwork address")

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}
	formDecoder := form.NewDecoder()

	/*
		// 9.2 Setting up the session manager

		// Use the scs.New() function to initialize a new session manager. Then we
		// configure it to use our MySQL database as the session store, and set a
		// lifetime of 12 hours (so that sessions automatically expire 12 hours
		// after first being created).
	*/
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour


	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder: formDecoder,

		/*
			// 9.2 Setting up the session manager

			// And add the session manager to our application dependencies.
		*/
		sessionManager: sessionManager,
	}

	srv := &http.Server{
		Addr:	*addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Staring server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

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