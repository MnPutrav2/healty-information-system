package routes

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/controllers"
	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
)

func Examination(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetExaminationData(w, r, db, path)

	case "POST":
		controllers.CreateExaminationData(w, r, db, path)

	case "PUT":
		controllers.UpdateExaminationData(w, r, db, path)

	case "DELETE":
		controllers.DeleteExaminationData(w, r, db, path)

	default:
		helper.ResponseError(w, "", "method not allowed", "method not allowed : 400", 400, path)
	}

}
