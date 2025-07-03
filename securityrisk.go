package gojsonq

import (
	"database/sql"
)

type DBHandler struct {
	db *sql.DB
}

func NewDBHandler(connStr string) (*DBHandler, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DBHandler{db: db}, nil
}

// 🚨 SQL Injection risk: unsafe string concatenation
func (h *DBHandler) GetUserByName(username string) (string, error) {
	query := "SELECT email FROM users WHERE username = '" + username + "'"
	row := h.db.QueryRow(query)
	var email string
	err := row.Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}

// Proper parameterized query for comparison
func (h *DBHandler) GetUserByNameSafe(username string) (string, error) {
	query := "SELECT email FROM users WHERE username = $1"
	row := h.db.QueryRow(query, username)
	var email string
	err := row.Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}
