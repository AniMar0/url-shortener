package internal

import (
	"database/sql"
	"net/http"
)

type Server struct {
	Host         string         `json:"host"`
	Port         string         `json:"port"`
	DatabasePath string         `json:"database_path"`
	Database     *sql.DB        `json:"-"`
	Router       *http.ServeMux `json:"-"`
}
