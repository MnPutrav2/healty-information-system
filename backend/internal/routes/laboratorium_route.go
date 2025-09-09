package routes

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/controllers"
	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
)

func LaboratoriumData(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetLabolatoriumData(w, r, db, path)

	case "POST":
		controllers.CreateLabolatoriumData(w, r, db, path)

	case "PUT":
		controllers.UpdateLabolatoriumData(w, r, db, path)

	case "DELETE":
		controllers.DeleteLabolatoriumData(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}

func LaboratoriumTemplate(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetLabolatoriumTemplate(w, r, db, path)

	case "POST":
		controllers.CreateLabolatoriumTemplateData(w, r, db, path)

	case "DELETE":
		controllers.DeleteLabolatoriumTemplate(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}

func LaboratoriumRequest(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetLabolatoriumRequest(w, r, db, path)

	case "POST":
		controllers.CreateLabolatoriumRequest(w, r, db, path)

	case "PUT":
		controllers.UpdateLabolatoriumSample(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}

func LaboratoriumDetail(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetLabolatoriumDataDetail(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}

func LaboratoriumValue(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "PUT":
		controllers.UpdateLabolatoriumValue(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}
