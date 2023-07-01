package main

import (
	"database/sql"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/borntodie-new/information/pkg/config"
	db "github.com/borntodie-new/information/pkg/db/sqlc"
	"github.com/borntodie-new/information/pkg/handler"
	"github.com/borntodie-new/information/pkg/render"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	cnf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config")
	}
	conn, err := sql.Open(cnf.DBDriver, cnf.DBDriver)
	if err != nil {
		log.Fatalln("cannot connect the db", err.Error())
	}
	app.InProduction = false

	session = scs.New()
	// 设置 session 的有效期
	session.Lifetime = time.Hour * 24
	// 设置 session
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot create template cache", err.Error())
	}
	app.TemplateCache = tc
	app.UseCache = false
	app.Session = session

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	store := db.NewStore(conn)
	handler.NewStore(store)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.Abort)
	fmt.Printf("Starting application on %s\n", cnf.HTTPServerAddress)
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    cnf.HTTPServerAddress,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()

	log.Fatal(err)
}
