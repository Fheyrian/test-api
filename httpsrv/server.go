package httpsrv

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	ListenIP   string
	ListenPort string
	DBconnStr  string

	router chi.Router
	db     *gorm.DB
}

func (s *Server) Init() error {
	var err error
	db, err := gorm.Open(mysql.Open(s.DBconnStr), &gorm.Config{})
	if err != nil {
		return err
	}
	s.db = db

	s.router = chi.NewRouter()
	s.router.Route("/api", func(r chi.Router) {
		r.Get("/daftar-mapel", s.listMapel)
		r.Post("/daftar-mapel", s.createMapel)
		r.Delete("/daftar-mapel/{id}", s.deleteMapel)
	})

	return nil
}

func (s *Server) Start() error {
	listenAddr := fmt.Sprintf("%s:%s", s.ListenIP, s.ListenPort)
	return http.ListenAndServe(listenAddr, s.router)
}
