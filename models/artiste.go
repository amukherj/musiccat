package models

import (
	"time"

	"github.com/gocraft/dbr"
)

const (
	artisteTbl = "artistes"
)

type Artiste struct {
	BaseModel
	ID             int64
	Name           string
	ArtisteType    string
	StartYear      int64
	LastActiveYear int64
	CreatedAt      int64
	UpdatedAt      int64
}

func (a *Artiste) TableName() string {
	return artisteTbl
}

func (a *Artiste) GetID() int64 {
	return a.ID
}

func (a *Artiste) Load(dbSession *dbr.Session, id int64) error {
	var r Record = a
	return loadRecord(r, dbSession, id)
}

func (a *Artiste) LoadByName(dbSession *dbr.Session, name string) error {
	var r Record = a
	return loadRecordByName(r, dbSession, name)
}

func (a *Artiste) Create(dbSession *dbr.Session) (int64, error) {
	a.CreatedAt = time.Now().Unix()
	a.UpdatedAt = a.CreatedAt

	columns := []string{"name", "artiste_type", "start_year",
		"last_active_year", "created_at", "updated_at"}
	result, err := dbSession.InsertInto(a.TableName()).Columns(columns...).
		Record(a).Exec()
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (a *Artiste) Update(dbSession *dbr.Session) (int64, error) {
	updStmt := dbSession.Update(a.TableName()).Where("id = ?", a.ID)
	count := 0

	if a.Name != "" {
		updStmt = updStmt.Set("name", a.Name)
		count += 1
	}

	if a.ArtisteType != "" {
		updStmt = updStmt.Set("artiste_type", a.ArtisteType)
		count += 1
	}

	if a.StartYear > 0 {
		updStmt = updStmt.Set("start_year", a.StartYear)
		count += 1
	}

	if a.LastActiveYear > 0 {
		updStmt = updStmt.Set("last_active_year", a.LastActiveYear)
		count += 1
	}

	if count > 0 {
		a.UpdatedAt = time.Now().Unix()
		res, err := updStmt.Set("updated_at", a.UpdatedAt).Exec()
		if err != nil {
			return 0, err
		}
		return res.RowsAffected()
	} else {
		// no real updates to perform
		return 0, nil
	}
}

func (a *Artiste) Delete(dbSession *dbr.Session) (int64, error) {
	var r Record = a
	return deleteRecord(r, dbSession)
}
