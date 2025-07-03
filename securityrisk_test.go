package gojsonq

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUserByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db: %v", err)
	}
	defer db.Close()

	handler := &DBHandler{db: db}

	username := "alice"
	expectedEmail := "alice@example.com"
	mock.ExpectQuery("SELECT email FROM users WHERE username = 'alice'").
		WillReturnRows(sqlmock.NewRows([]string{"email"}).AddRow(expectedEmail))

	email, err := handler.GetUserByName(username)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if email != expectedEmail {
		t.Errorf("Got %s; want %s", email, expectedEmail)
	}
}

func TestGetUserByNameSafe(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock db: %v", err)
	}
	defer db.Close()

	handler := &DBHandler{db: db}

	username := "bob"
	expectedEmail := "bob@example.com"
	mock.ExpectQuery("SELECT email FROM users WHERE username = \\$1").
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows([]string{"email"}).AddRow(expectedEmail))

	email, err := handler.GetUserByNameSafe(username)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if email != expectedEmail {
		t.Errorf("Got %s; want %s", email, expectedEmail)
	}
}
