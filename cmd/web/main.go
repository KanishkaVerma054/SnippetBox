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
	"crypto/tls"
)
type application struct {
	errorLog		*log.Logger
	infoLog			*log.Logger
	snippets 		*models.SnippetModel
	users			*models.UserModel
	templateCache 	map[string]*template.Template
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
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	/*
		// 10.2 Running a HTTPS server

		// Make sure that the Secure attribute is set on our session cookies.
		// Setting this means that the cookie will only be sent by a user's web
		// browser when a HTTPS connection is being used (and won't be sent over an
		// unsecure HTTP connection).
	*/
	// sessionManager.Cookie.Secure = true
	sessionManager.Cookie.Secure = false // to run it locally


	app := &application{
		errorLog: 		errorLog,
		infoLog: 		infoLog,
		snippets: 		&models.SnippetModel{DB: db},
		users: 			&models.UserModel{DB: db},
		templateCache: 	templateCache,
		formDecoder: 	formDecoder,
		sessionManager: sessionManager,
	}

	/*
		// 10.2 Configuring HTTPS settings

		// Initialize a tls.Config struct to hold the non-default TLS settings we
		// want the server to use. In this case the only thing that we're changing
		// is the curve preferences value, so that only elliptic curves with
		// assembly implementations are used.
	*/
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:	*addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
		TLSConfig: tlsConfig,

		/*
			10.2 Configuring HTTPS settings

			// Add Idle, Read and Write timeouts to the server.
		*/
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Staring server on %s", *addr)
	
	err = srv.ListenAndServe()
	/*
		// 10.2 Running a HTTPS server

		// Use the ListenAndServeTLS() method to start the HTTPS server. We
		// pass in the paths to the TLS certificate and corresponding private key as
		// the two parameters.
	*/
	// err = srv.ListenAndServeTLS("./tls/localhost+2.pem", "./tls/localhost+2-key.pem")

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