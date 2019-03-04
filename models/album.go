package models

import (
	"time"

	"github.com/gocraft/dbr"
)

const (
	albumTbl = "albums"
)

type Album struct {
	BaseModel
	ID          int64
	Name        string
	ArtisteID   int64
	ArtisteName string
	GenreID     int64
	ReleaseYear int64
	CreatedAt   int64
	UpdatedAt   int64
}

func (a *Album) TableName() string {
	return albumTbl
}

func (a *Album) GetID() int64 {
	return a.ID
}

func (a *Album) Load(dbSession *dbr.Session, id int64) error {
	var r Record = a
	return loadRecord(r, dbSession, id)
}

func (a *Album) LoadByName(dbSession *dbr.Session, name string) error {
	var r Record = a
	return loadRecordByName(r, dbSession, name)
}

func (a *Album) Create(dbSession *dbr.Session) (int64, error) {
	a.CreatedAt = time.Now().Unix()
	a.UpdatedAt = a.CreatedAt

	columns := []string{"name", "artiste_id", "genre_id",
		"release_year", "created_at", "updated_at"}
	result, err := dbSession.InsertInto(a.TableName()).Columns(columns...).
		Record(a).Exec()
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (a *Album) Update(dbSession *dbr.Session) (int64, error) {
	a.UpdatedAt = time.Now().Unix()
	res, err := dbSession.Update(a.TableName()).Where("id = ?", a.ID).
		Set("name", a.Name).
		Set("artiste_id", a.ArtisteID).
		Set("genre_id", a.GenreID).
		Set("release_year", a.ReleaseYear).
		Set("updated_at", a.UpdatedAt).
		Exec()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (a *Album) Delete(dbSession *dbr.Session) (int64, error) {
	var r Record = a
	return deleteRecord(r, dbSession)
}
