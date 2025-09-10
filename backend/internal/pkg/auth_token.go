package pkg

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/MnPutrav2/healty-information-system/backend/internal/repository"
	"github.com/google/uuid"
)

func SessionToken(sql *sql.DB, id string) uuid.UUID {
	authRepo := repository.NewAuthRepository(sql)
	token := authRepo.CreateSessionToken(id)

	return token
}

func PatientSessionToken(sql *sql.DB, us string, pas string) uuid.UUID {
	authRepo := repository.NewAuthRepository(sql)
	token := authRepo.CreatePatientSessionToken(us, pas)

	return token
}

func CheckAuthorization(w http.ResponseWriter, path string, db *sql.DB, auth string) bool {
	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		return false
	}

	if split[0] != "Bearer" {
		return false
	}

	token := split[1]

	authRepo := repository.NewAuthRepository(db)
	q := authRepo.CheckSessionToken(token)

	return q != 0
}

func CheckPatientAuthorization(w http.ResponseWriter, path string, db *sql.DB, auth string) bool {
	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		return false
	}

	if split[0] != "Bearer" {
		return false
	}

	token := split[1]

	authRepo := repository.NewAuthRepository(db)
	q := authRepo.CheckPatientSessionToken(token)

	return q != 0
}
