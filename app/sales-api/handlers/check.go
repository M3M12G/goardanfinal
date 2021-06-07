package handlers

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/m3m12g/goardanfinal/foundation/database"
	"github.com/m3m12g/goardanfinal/foundation/web"
	"net/http"
)

type checkGroup struct {
	db *sqlx.DB
}

func (c checkGroup) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := "ok"
	statusCode := http.StatusOK
	if err := database.StatusCheck(ctx, c.db); err != nil {
		status = "db not ready"
		statusCode = http.StatusInternalServerError
	}

	health := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	return web.Respond(ctx, w, health, statusCode)

}
