package models

import (
	"time"

	"github.com/gocraft/dbr"
)

const (
	actArtisteTbl = "acts_artistes"
)

type ActArtiste struct {
	BaseModel
	ID          int64
	ActID       int64
	ActName     string
	ArtisteID   int64
	ArtisteName string
	CreatedAt   int64
	UpdatedAt   int64
}

func (a *ActArtiste) TableName() string {
	return actArtisteTbl
}

func (a *ActArtiste) GetID() int64 {
	return a.ID
}

func (a *ActArtiste) Load(dbSession *dbr.Session, id int64) error {
	var r Record = a
	return loadRecord(r, dbSession, id)
}

// This method will fail because acts_artistes table does not have
// a column called `name`. It's implemented only for compatibility
// with the Record interface
func (a *ActArtiste) LoadByName(dbSession *dbr.Session, name string) error {
	var r Record = a
	return loadRecordByName(r, dbSession, name)
}

func (a *ActArtiste) Create(dbSession *dbr.Session) (int64, error) {
	a.CreatedAt = time.Now().Unix()
	a.UpdatedAt = a.CreatedAt

	columns := []string{"act_id", "artiste_id", "created_at", "updated_at"}
	result, err := dbSession.InsertInto(a.TableName()).Columns(columns...).
		Record(a).Exec()
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (a *ActArtiste) Update(dbSession *dbr.Session) (int64, error) {
	a.UpdatedAt = time.Now().Unix()
	res, err := dbSession.Update(a.TableName()).Where("id = ?", a.ID).
		Set("act_id", a.ActID).
		Set("artiste_id", a.ArtisteID).
		Set("updated_at", a.UpdatedAt).
		Exec()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (a *ActArtiste) Delete(dbSession *dbr.Session) (int64, error) {
	var r Record = a
	return deleteRecord(r, dbSession)
}
