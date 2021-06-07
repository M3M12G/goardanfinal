// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/m3m12g/goardanfinal/business/auth"
	"github.com/m3m12g/goardanfinal/business/mid"
	"log"
	"net/http"
	"os"

	"github.com/m3m12g/goardanfinal/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth, db *sqlx.DB) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics(log))

	check := checkGroup{
		db: db,
	}

	app.Handle(http.MethodGet, "/readiness", check.readiness, mid.Authenticate(a), mid.Authorize(auth.RoleAdmin))

	return app
}
