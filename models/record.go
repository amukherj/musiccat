package models

import (
	"github.com/gocraft/dbr"
)

type Record interface {
	TableName() string
	GetID() int64
	Load(dbSession *dbr.Session, id int64) error
	LoadByName(dbSession *dbr.Session, name string) error
	Create(dbSession *dbr.Session) (int64, error)
	Update(dbSession *dbr.Session) (int64, error)
	Delete(dbSession *dbr.Session) (int64, error)
}

func loadRecord(r Record, dbSession *dbr.Session, id int64) error {
	_, err := dbSession.Select("*").
		From(r.TableName()).
		Where("id = ?", id).
		Load(r)

	return err
}

func loadRecordByName(r Record, dbSession *dbr.Session, name string) error {
	_, err := dbSession.Select("*").
		From(r.TableName()).
		Where("name = ?", name).
		Limit(1).
		Load(r)

	return err
}

func deleteRecord(r Record, dbSession *dbr.Session) (int64, error) {
	res, err := dbSession.DeleteFrom(r.TableName()).
		Where("id = ?", r.GetID()).Exec()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
