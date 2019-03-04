package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/gocraft/dbr"

	"http/musiccat/models"
)

func mustGetDBSession(ctx context.Context) *dbr.Session {
	dbSess, ok := ctx.Value(contextkey("dbsess")).(*dbr.Session)
	if !ok {
		panic(fmt.Sprintf("Could not retrieve database session from context"))
	}
	return dbSess
}

func sendResponse(w http.ResponseWriter, r *http.Request, id int64,
	nrows int64, err error, opName string, objType string) {
	if err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.Render(w, r, &models.APIFailure{
			Message:   fmt.Sprintf("The %s operation failed", opName),
			ErrorCode: 1,
			HTTPCode:  http.StatusUnprocessableEntity,
			Time:      time.Now().Unix(),
		})
	} else if nrows > 0 {
		render.Render(w, r, &models.APISuccess{
			Message: fmt.Sprintf("The %s operation succeeded: %s id=%d",
				opName, objType, id),
			ID:   id,
			Time: time.Now().Unix(),
		})
	} else {
		render.Render(w, r, &models.APISuccess{
			Message: fmt.Sprintf("The %s operation did not affect any records",
				opName),
			ID:   id,
			Time: time.Now().Unix(),
		})
	}
}
