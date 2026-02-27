package internal

import (
	"database/sql"
	"net/http"
	Rp "url-shortener/internal/repository"
)

type Server struct {
	Host          string         `json:"host"`
	Port          string         `json:"port"`
	DatabasePath  string         `json:"database_path"`
	Database      *sql.DB        `json:"-"`
	Router        *http.ServeMux `json:"-"`
}

func NewServer(host, port, databasePath string) *Server {
	return &Server{
		Host:         host,
		Port:         port,
		DatabasePath: databasePath,
	}
}

func (s *Server) Start() error {
	s.Database = Rp.InitDB(s.DatabasePath)
	s.Router = http.NewServeMux()
	s.routes()
	return http.ListenAndServe(s.Host+":"+s.Port, s.Router)
}

func (s *Server) routes() {
	//s.Router.HandleFunc("/api/register", s.handleRegister)
	//s.Router.HandleFunc("/api/login", s.handleLogin)
	//s.Router.HandleFunc("/", s.handleRedirect)
}
