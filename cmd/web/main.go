package main

import (
	"encoding/gob"
	"fmt"
	"github.com/aaboemira/bookings/internal/config"
	"github.com/aaboemira/bookings/internal/handlers"
	"github.com/aaboemira/bookings/internal/models"
	"github.com/aaboemira/bookings/internal/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	fmt.Println(fmt.Sprintf("Starting application on port   %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
func run() error {
	gob.Register(models.Reservation{})

	//change this to true when in production

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	tc, err := render.TemplateCreate()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false
	fmt.Println(app.TemplateCache)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	return nil

}
