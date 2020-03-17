package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
	"github.com/ihaveint/snippetbox/pkg/models/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
	session       *sessions.Session
	users         *mysql.UserModel
}

func main() {
	port := flag.String("port", ":4000", "HTTP port")
	dsn := flag.String("dsn", fmt.Sprintf("web:%s@/snippetbox?parseTime=true", os.Getenv("MySQLPass")), "MySQL DSN")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
		session:       session,
		users:         &mysql.UserModel{DB: db},
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *port,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Startingg server on %s", *port)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
