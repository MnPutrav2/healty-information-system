package routes

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/controllers"
	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
)

func Employee(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetEmployeeData(w, r, db, path)

	case "POST":
		controllers.AddEmployeeData(w, r, db, path)

	case "PUT":
		controllers.UpdateEmployeeData(w, r, db, path)

	case "DELETE":
		controllers.DeleteEmployeeData(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}

func Users(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "POST":
		controllers.AddUserLogin(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}
