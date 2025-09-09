package routes

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/controllers"
	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
)

func Ambulatory(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetAmbulatoryCarePatient(w, r, db, path)

	case "POST":
		controllers.CreateAmbulatoryCarePatient(w, r, db, path)

	case "PUT":
		controllers.UpdateAmbulatoryCarePatient(w, r, db, path)

	case "DELETE":
		controllers.DeleteAmbulatoryCarePatient(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}
