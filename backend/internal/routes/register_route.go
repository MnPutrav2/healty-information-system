package routes

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/controllers"
	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
)

func Registration(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetRegistrationPatient(w, r, db, path)

	case "POST":
		controllers.CreateRegistrationPatient(w, r, db, path)

	case "DELETE":
		controllers.DeleteRegistrationPatient(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}

func RegistrationUpdateStatus(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "PUT":
		controllers.UpdateStatus(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}
