package handlers

import (
	"context"
	"errors"
	"github.com/m3m12g/goardanfinal/foundation/web"
	"log"
	"math/rand"
	"net/http"
)

type check struct {
	logger *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	if n := rand.Intn(100); n%2 == 0 {
		return errors.New("untrusted error")
	}

	status := struct {
		Status string
	}{
		Status: "Ok",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
