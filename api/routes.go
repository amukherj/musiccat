package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gocraft/dbr"

	"http/musiccat/db"
)

// Use this type for the
// key types in a context.
type contextkey string

var once sync.Once

type APIServer struct {
	router  chi.Router
	session *dbr.Session
}

// Make the API handler DB aware by
// making the db session available
// through a key-value pair in context.
func (a *APIServer) WithDB(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, contextkey("dbsess"), a.session)
			handler.ServeHTTP(w, r.WithContext(ctx))
		})
}

func newAPIServer(dbPath string) (*APIServer, error) {
	if sess, err := db.OpenDB(dbPath); err != nil {
		return &APIServer{}, err
	} else {
		return &APIServer{
			router:  chi.NewRouter(),
			session: sess,
		}, nil
	}
}

func (s *APIServer) SetupRoutes() {
	r := s.router
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE",
			"PATCH", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         300,
	})
	r.Use(cors.Handler)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	s.setupArtisteRoutes()
	s.setupAlbumRoutes()

	if err := http.ListenAndServe(":3333", r); err != nil {
		panic(fmt.Sprintf("Failed to start API router: %v\n", err))
	}
}

func (s *APIServer) setupArtisteRoutes() {
	r := s.router
	r.Route("/artistes", func(r chi.Router) {
		// Get list of artistes
		r.Get("/", s.WithDB(getArtistes))

		// Create a new artiste
		r.Post("/", s.WithDB(createArtiste))

		r.Route("/{artisteID:\\d+}", func(r chi.Router) {
			// set the context for this request
			// by preloading the requested record
			// r.Use(artisteCtx)

			r.Get("/", s.WithDB(getArtiste))
			// Update an existing artiste
			r.Patch("/", s.WithDB(patchArtiste))
			// Delete an existing artiste
			r.Delete("/", s.WithDB(deleteArtiste))

			// Get albums for the artiste
			r.Route("/albums", func(r chi.Router) {
				r.Get("/", s.WithDB(getArtisteAlbums))
				r.Post("/", s.WithDB(createArtisteAlbum))
				r.Route("/{albumID:\\d+}", func(r chi.Router) {
					r.Patch("/", s.WithDB(patchAlbum))
					r.Delete("/", s.WithDB(deleteAlbum))
				})
			})

			r.Get("/members", s.WithDB(getActMembers))
			r.Route("/acts", func(r chi.Router) {
				r.Get("/", s.WithDB(getArtisteActs))
				r.Post("/{actID:\\d+}", s.WithDB(associateArtisteWithAct))
			})
		})
	})
}

func (s *APIServer) setupAlbumRoutes() {
	r := s.router
	r.Route("/albums", func(r chi.Router) {
		r.Get("/", s.WithDB(getAlbums))
	})
}

func StartAPIServer(dbPath string) {
	once.Do(func() {
		if s, err := newAPIServer(dbPath); err != nil {
			panic(fmt.Sprintf("Failed to start API server: %v", err))
		} else {
			s.SetupRoutes()
		}
	})
}
