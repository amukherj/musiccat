package models

import (
	"time"

	"github.com/gocraft/dbr"
)

const (
	genreTbl = "genres"
)

type Genre struct {
	BaseModel
	ID        int64
	Name      string
	CreatedAt int64
	UpdatedAt int64
}

func (a *Genre) TableName() string {
	return genreTbl
}

func (g *Genre) GetID() int64 {
	return g.ID
}

func (g *Genre) Load(dbSession *dbr.Session, id int64) error {
	var r Record = g
	return loadRecord(r, dbSession, id)
}

func (g *Genre) LoadByName(dbSession *dbr.Session, name string) error {
	var r Record = g
	return loadRecordByName(r, dbSession, name)
}

func (g *Genre) Create(dbSession *dbr.Session) (int64, error) {
	g.CreatedAt = time.Now().Unix()
	g.UpdatedAt = g.CreatedAt

	columns := []string{"name", "created_at", "updated_at"}
	result, err := dbSession.InsertInto(g.TableName()).Columns(columns...).
		Record(g).Exec()
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	return id, err
}

func (g *Genre) Update(dbSession *dbr.Session) (int64, error) {
	g.UpdatedAt = time.Now().Unix()
	res, err := dbSession.Update(g.TableName()).Where("id = ?", g.ID).
		Set("name", g.Name).
		Set("updated_at", g.UpdatedAt).
		Exec()
	nrows, _ := res.RowsAffected()
	return nrows, err
}

func (g *Genre) Delete(dbSession *dbr.Session) (int64, error) {
	var r Record = g
	return deleteRecord(r, dbSession)
}
