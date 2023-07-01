package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/borntodie-new/information/pkg/config"
	"github.com/borntodie-new/information/pkg/handler"
	"github.com/borntodie-new/information/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
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

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.Abort)
	fmt.Printf("Starting application on port %s\n", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()

	log.Fatal(err)
}
