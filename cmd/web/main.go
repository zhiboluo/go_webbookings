package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/zhiboluo/go_webbookings/pkg/config"
	"github.com/zhiboluo/go_webbookings/pkg/handlers"
	"github.com/zhiboluo/go_webbookings/pkg/render"
)

const portNum = ":8083"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// changethis to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("starting app on port:", portNum)
	// _ = http.ListenAndServe(portNum, nil)

	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
