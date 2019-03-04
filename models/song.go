package models

import (
	"time"

	"github.com/gocraft/dbr"
)

const (
	songTbl = "songs"
)

type Song struct {
	BaseModel
	ID        int64
	Name      string
	ArtisteID int64
	AlbumID   int64
	GenreID   int64
	Credits   string
	CreatedAt int64
	UpdatedAt int64
}

func (s *Song) TableName() string {
	return songTbl
}

func (s *Song) GetID() int64 {
	return s.ID
}

func (s *Song) Load(dbSession *dbr.Session, id int64) error {
	var r Record = s
	return loadRecord(r, dbSession, id)
}

func (s *Song) LoadByName(dbSession *dbr.Session, name string) error {
	var r Record = s
	return loadRecordByName(r, dbSession, name)
}

func (s *Song) Create(dbSession *dbr.Session) (int64, error) {
	s.CreatedAt = time.Now().Unix()
	s.UpdatedAt = s.CreatedAt

	columns := []string{"name", "artiste_id", "album_id", "genre_id",
		"credits", "created_at", "updated_at"}
	result, err := dbSession.InsertInto(s.TableName()).Columns(columns...).
		Record(s).Exec()
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (s *Song) Update(dbSession *dbr.Session) (int64, error) {
	s.UpdatedAt = time.Now().Unix()
	res, err := dbSession.Update(s.TableName()).Where("id = ?", s.ID).
		Set("name", s.Name).
		Set("artiste_id", s.ArtisteID).
		Set("album_id", s.AlbumID).
		Set("genre_id", s.GenreID).
		Set("credits", s.Credits).
		Set("updated_at", s.UpdatedAt).
		Exec()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (s *Song) Delete(dbSession *dbr.Session) (int64, error) {
	var r Record = s
	return deleteRecord(r, dbSession)
}
