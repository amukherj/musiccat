package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/amukherj/musiccat/models"
)

func artisteCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// Artiste APIs
func getArtistes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	// query and render
	var result []models.Artiste
	dbSess.Select("*").From("artistes").Load(&result)
	resp := []render.Renderer{}
	for _, a := range result {
		resp = append(resp, a)
	}
	render.RenderList(w, r, resp)
}

func getArtiste(w http.ResponseWriter, r *http.Request) {
	artst := chi.URLParam(r, "artisteID")
	artisteID, _ := strconv.Atoi(artst)
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	// query and render
	var result models.Artiste
	dbSess.Select("*").From("artistes").Where("id = ?", artisteID).
		Load(&result)
	render.Render(w, r, result)
}

func createArtiste(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	var payload struct {
		Artiste models.Artiste
	}
	req, err := ioutil.ReadAll(r.Body)
	if err = json.Unmarshal(req, &payload); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, &models.APIFailure{
			Message:   "Could not parse the request body",
			ErrorCode: 1,
			HTTPCode:  http.StatusBadRequest,
			Time:      time.Now().Unix(),
		})
		return
	}

	if payload.Artiste.Name == "" {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, &models.APIFailure{
			Message:   "Artiste name cannot be blank",
			ErrorCode: 1,
			HTTPCode:  http.StatusBadRequest,
			Time:      time.Now().Unix(),
		})
		return
	}

	payload.Artiste.ID = 0
	payload.Artiste.CreatedAt = time.Now().Unix()
	payload.Artiste.UpdatedAt = payload.Artiste.CreatedAt

	id, err := payload.Artiste.Create(dbSess)
	sendResponse(w, r, id, 1, err, "insert", "artiste")
}

func patchArtiste(w http.ResponseWriter, r *http.Request) {
	artst := chi.URLParam(r, "artisteID")
	artisteID, _ := strconv.Atoi(artst)
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	var payload struct {
		Artiste models.Artiste
	}
	req, err := ioutil.ReadAll(r.Body)
	if err = json.Unmarshal(req, &payload); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, &models.APIFailure{
			Message:   "Could not parse the request body",
			ErrorCode: 1,
			HTTPCode:  http.StatusBadRequest,
			Time:      time.Now().Unix(),
		})
		return
	}
	payload.Artiste.ID = int64(artisteID)
	payload.Artiste.UpdatedAt = time.Now().Unix()
	nrows, err := payload.Artiste.Update(dbSess)
	sendResponse(w, r, int64(artisteID), nrows, err, "update", "artiste")
}

func deleteArtiste(w http.ResponseWriter, r *http.Request) {
	artst := chi.URLParam(r, "artisteID")
	artisteID, _ := strconv.Atoi(artst)
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)
	artiste := models.Artiste{ID: int64(artisteID)}
	nrows, err := artiste.Delete(dbSess)
	sendResponse(w, r, int64(artisteID), nrows, err, "delete", "artiste")
}

func getArtisteAlbums(w http.ResponseWriter, r *http.Request) {
	artst := chi.URLParam(r, "artisteID")
	artisteID, _ := strconv.Atoi(artst)
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	// query and render
	var result []models.Album
	_, err := dbSess.Select("artistes.Name as artiste_name, albums.*").
		From("artistes").
		Join("albums", "albums.artiste_id = artistes.id").
		Where("artistes.id = ?", artisteID).
		Load(&result)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		var response []render.Renderer
		for _, res := range result {
			response = append(response, &res)
		}
		render.RenderList(w, r, response)
	}
}

func createArtisteAlbum(w http.ResponseWriter, r *http.Request) {
	artst := chi.URLParam(r, "artisteID")
	artisteID, _ := strconv.Atoi(artst)
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	var payload struct {
		Album models.Album
	}
	req, err := ioutil.ReadAll(r.Body)
	if err = json.Unmarshal(req, &payload); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, &models.APIFailure{
			Message:   "Could not parse the request body",
			ErrorCode: 1,
			HTTPCode:  http.StatusBadRequest,
			Time:      time.Now().Unix(),
		})
		return
	}
	payload.Album.ArtisteID = int64(artisteID)
	payload.Album.CreatedAt = time.Now().Unix()
	payload.Album.UpdatedAt = payload.Album.CreatedAt

	id, err := payload.Album.Create(dbSess)
	sendResponse(w, r, id, 1, err, "insert", "album")
}

func patchAlbum(w http.ResponseWriter, r *http.Request) {
	artst := chi.URLParam(r, "artisteID")
	albm := chi.URLParam(r, "albumID")
	artisteID, _ := strconv.Atoi(artst)
	albumID, _ := strconv.Atoi(albm)
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	var payload struct {
		Album models.Album
	}
	req, err := ioutil.ReadAll(r.Body)
	if err = json.Unmarshal(req, &payload); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, &models.APIFailure{
			Message:   "Could not parse the request body",
			ErrorCode: 1,
			HTTPCode:  http.StatusBadRequest,
			Time:      time.Now().Unix(),
		})
		return
	}
	payload.Album.ID = int64(albumID)
	payload.Album.ArtisteID = int64(artisteID)
	payload.Album.UpdatedAt = payload.Album.CreatedAt

	nrows, err := payload.Album.Update(dbSess)
	sendResponse(w, r, int64(albumID), nrows, err, "update", "album")
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	// artst := chi.URLParam(r, "artisteID")
	albm := chi.URLParam(r, "albumID")
	// artisteID, _ := strconv.Atoi(artst)
	albumID, _ := strconv.Atoi(albm)
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	var album models.Album
	album.ID = int64(albumID)

	nrows, err := album.Delete(dbSess)
	sendResponse(w, r, int64(albumID), nrows, err, "delete", "album")
}

func getActMembers(w http.ResponseWriter, r *http.Request) {
	artst := chi.URLParam(r, "artisteID")
	artisteID, _ := strconv.Atoi(artst)
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	var assocs []models.ActArtiste
	_, err := dbSess.SelectBySql(
		`SELECT aa.*, a.Name as act_name, b.Name as artiste_name
		FROM acts_artistes aa INNER JOIN artistes a ON a.id = aa.act_id
		INNER JOIN artistes b ON b.id = aa.artiste_id
		WHERE aa.act_id = ?`, int64(artisteID)).Load(&assocs)

	if err != nil {
		// FIXME
		fmt.Printf("Error: %v\n", err)
	} else {
		response := []render.Renderer{}
		for _, a := range assocs {
			response = append(response, a)
		}
		render.RenderList(w, r, response)
	}
}

func getArtisteActs(w http.ResponseWriter, r *http.Request) {
	artst := chi.URLParam(r, "artisteID")
	artisteID, _ := strconv.Atoi(artst)
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	var assocs []models.ActArtiste
	_, err := dbSess.SelectBySql(
		`SELECT aa.*, a.name as artiste_name, b.name as act_name
		FROM acts_artistes aa INNER JOIN artistes a ON a.id = aa.artiste_id
		INNER JOIN artistes b ON b.id = aa.act_id
		WHERE aa.artiste_id = ?`, int64(artisteID)).Load(&assocs)

	if err != nil {
		// FIXME
		fmt.Printf("Error: %v\n", err)
	} else {
		response := []render.Renderer{}
		for _, a := range assocs {
			response = append(response, a)
		}
		render.RenderList(w, r, response)
	}
}

func associateArtisteWithAct(w http.ResponseWriter, r *http.Request) {
	act := chi.URLParam(r, "actID")
	mmbr := chi.URLParam(r, "artisteID")
	actID, _ := strconv.Atoi(act)
	memberID, _ := strconv.Atoi(mmbr)
	fmt.Printf("Associating member %d with act %d\n", memberID, actID)
	ctx := r.Context()
	// get db session from context
	dbSess := mustGetDBSession(ctx)

	var assoc models.ActArtiste
	assoc.ActID = int64(actID)
	assoc.ArtisteID = int64(memberID)
	assoc.CreatedAt = time.Now().Unix()
	assoc.UpdatedAt = assoc.CreatedAt
	id, err := assoc.Create(dbSess)
	sendResponse(w, r, id, 1, err, "create", "artiste act association")
}
