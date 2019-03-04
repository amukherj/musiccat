package main

import (
	// "fmt"
	// "os"
	// "time"

	"http/musiccat/api"
	// "http/musiccat/db"
	// "http/musiccat/models"
	// "github.com/gocraft/dbr"
	// _ "github.com/mattn/go-sqlite3"
)

const (
	DBPath = "/home/amukher1/devel/go/src/http/musiccat/data/musiccat.db"
)

func main() {
	api.StartAPIServer(DBPath)
	// get DB session
	/* dbSession, err := db.OpenDB(DBPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	if dbSession != nil {
		fmt.Printf("Valid db connection: %p\n", dbSession)
	}
	testGenre(dbSession)
	testArtiste(dbSession)
	testActArtiste(dbSession)
	testAlbum(dbSession)
	testSong(dbSession) */
}

/* func testGenre(dbSession *dbr.Session) {
	// Insert
	g := models.Genre{
		Name: "classical/western",
	}
	id, err := g.Create(dbSession)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting record: %+v. Reason: %v\n", g,
			err)
	} else {
		fmt.Printf("New genre (id = %d) added\n", id)
	}

	// Lookup by id
	var g2 models.Genre
	err = g2.Load(dbSession, 9)

	// Update
	g2.Name = "classical/persian"
	if nrows, err := g2.Update(dbSession); err != nil {
		fmt.Fprintf(os.Stderr, "Error updating record: %+v. Reason: %v\n", g2,
			err)
	} else if nrows > 0 {
		fmt.Printf("Updated genre (id = %d)\n", g2.ID)
	} else {
		fmt.Printf("No rows updated\n")
	}

	// Lookup by id
	var g3 models.Genre
	err = g3.Load(dbSession, 10)

	// Delete
	if nrows, err := g3.Delete(dbSession); err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting record: %+v. Reason: %v\n", g3,
			err)
	} else if nrows > 0 {
		fmt.Printf("Deleted genre (id = %d)\n", g3.ID)
	} else {
		fmt.Printf("No rows deleted\n")
	}
}

func testArtiste(dbSession *dbr.Session) {
	// Insert
	a := models.Artiste{
		Name:           "Mishra, Mangal",
		ArtisteType:    "Solo",
		StartYear:      time.Date(1996, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		LastActiveYear: 0,
	}
	id, err := a.Create(dbSession)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting record: %+v. Reason: %v\n", a,
			err)
	} else {
		fmt.Printf("New artiste (id = %d) added\n", id)
	}

	// Lookup by id
	var a2 models.Artiste
	err = a2.Load(dbSession, id)

	// Update
	a2.Name = "Mouskouri, Nana"
	if nrows, err := a2.Update(dbSession); err != nil {
		fmt.Fprintf(os.Stderr, "Error updating record: %+v. Reason: %v\n", a2,
			err)
	} else if nrows > 0 {
		fmt.Printf("Updated artiste (id = %d)\n", a2.ID)
	} else {
		fmt.Printf("No rows updated\n")
	}

	// Lookup by id
	var a3 models.Artiste
	err = a3.Load(dbSession, a2.ID-1)

	// Delete
	if nrows, err := a3.Delete(dbSession); err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting record: %+v. Reason: %v\n", a3,
			err)
	} else if nrows > 0 {
		fmt.Printf("Deleted artiste (id = %d)\n", a3.ID)
	} else {
		fmt.Printf("No rows deleted\n")
	}
}

func testActArtiste(dbSession *dbr.Session) {
	// Insert
	a := models.ActArtiste{
		ActID:     1,
		ArtisteID: 2,
	}
	id, err := a.Create(dbSession)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting record: %+v. Reason: %v\n", a,
			err)
	} else {
		fmt.Printf("New act-artiste relation (id = %d) added\n", id)
	}

	// Lookup by id
	var a2 models.ActArtiste
	err = a2.Load(dbSession, id)

	// Update
	a2.ArtisteID = 3
	if nrows, err := a2.Update(dbSession); err != nil {
		fmt.Fprintf(os.Stderr, "Error updating record: %+v. Reason: %v\n", a2,
			err)
	} else if nrows > 0 {
		fmt.Printf("Updated act-artiste relation (id = %d)\n", a2.ID)
	} else {
		fmt.Printf("No rows updated\n")
	}

	// Lookup by id
	var a3 models.ActArtiste
	err = a3.Load(dbSession, a2.ID-1)

	// Delete
	if nrows, err := a3.Delete(dbSession); err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting record: %+v. Reason: %v\n", a3,
			err)
	} else if nrows > 0 {
		fmt.Printf("Deleted act-artiste relation (id = %d)\n", a3.ID)
	} else {
		fmt.Printf("No rows deleted\n")
	}
}

func testAlbum(dbSession *dbr.Session) {
	// Insert
	a := models.Album{
		Name:        "Bhor bhaye",
		GenreID:     2,
		ReleaseYear: time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		ArtisteID:   3,
	}
	id, err := a.Create(dbSession)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting record: %+v. Reason: %v\n", a,
			err)
	} else {
		fmt.Printf("New album (id = %d) added\n", id)
	}

	// Lookup by id
	var a2 models.Album
	err = a2.Load(dbSession, id)

	// Update
	a2.Name = "London Calling"
	if nrows, err := a2.Update(dbSession); err != nil {
		fmt.Fprintf(os.Stderr, "Error updating record: %+v. Reason: %v\n", a2,
			err)
	} else if nrows > 0 {
		fmt.Printf("Updated album (id = %d)\n", a2.ID)
	} else {
		fmt.Printf("No rows updated\n")
	}

	// Lookup by id
	var a3 models.Album
	err = a3.Load(dbSession, a2.ID-1)

	// Delete
	if nrows, err := a3.Delete(dbSession); err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting record: %+v. Reason: %v\n", a3,
			err)
	} else if nrows > 0 {
		fmt.Printf("Deleted album (id = %d)\n", a3.ID)
	} else {
		fmt.Printf("No rows deleted\n")
	}
}

func testSong(dbSession *dbr.Session) {
	// Insert
	s := models.Song{
		Name:      "Bhor Bhaye",
		AlbumID:   1,
		ArtisteID: 2,
		GenreID:   2,
		Credits:   "Mangal Misra, Pandit Jasraj",
	}
	id, err := s.Create(dbSession)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error inserting record: %+v. Reason: %v\n", s,
			err)
	} else {
		fmt.Printf("New song (id = %d) added\n", id)
	}

	// Lookup by id
	var s2 models.Song
	err = s2.Load(dbSession, id)

	// Update
	s2.Name = "Fat Old Sun"
	if nrows, err := s2.Update(dbSession); err != nil {
		fmt.Fprintf(os.Stderr, "Error updating record: %+v. Reason: %v\n", s2,
			err)
	} else if nrows > 0 {
		fmt.Printf("Updated song (id = %d)\n", s2.ID)
	} else {
		fmt.Printf("No rows updated\n")
	}

	// Lookup by id
	var s3 models.Song
	err = s3.Load(dbSession, s2.ID-1)

	// Delete
	if nrows, err := s3.Delete(dbSession); err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting record: %+v. Reason: %v\n", s3,
			err)
	} else if nrows > 0 {
		fmt.Printf("Deleted song (id = %d)\n", s3.ID)
	} else {
		fmt.Printf("No rows deleted\n")
	}
} */
