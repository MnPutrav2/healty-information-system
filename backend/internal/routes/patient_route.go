package routes

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/controllers"
	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
)

func Patient(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetPatient(w, r, db, path)

	case "POST":
		controllers.CreatePatient(w, r, db, path)

	case "PUT":
		controllers.UpdatePatientData(w, r, db, path)

	case "DELETE":
		controllers.DeletePatient(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}

func PatientStatus(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetPatientStatus(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}

func PatientAuth(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "POST":
		controllers.AuthPatient(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}
