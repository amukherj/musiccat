package models

import (
	"net/http"
)

type BaseModel struct{}

func (BaseModel) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
