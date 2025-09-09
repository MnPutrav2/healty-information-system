package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MnPutrav2/healty-information-system/backend/internal/clients/satu_sehat/services"
	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
)

func GetSatuSehatPatient(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {
	// ---- Needed for every request ---
	if !pkg.RequestHeader(w, r) {
		return
	}

	// Check Header
	auth := r.Header.Get("Authorization")
	if !pkg.CheckAuthorization(w, path, db, auth) {
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	}

	split := strings.SplitN(auth, " ", 2)

	if len(split) != 2 || split[0] != "Bearer" {
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// Check Header
	// --- ---
	var id string
	if err := db.QueryRow("SELECT users.id FROM users INNER JOIN session_token ON users.id = session_token.users_id WHERE session_token.token = $1", split[1]).Scan(&id); err != nil {
		return
	}

	token, err := pkg.CreateSatuSehatToken(db)
	if err != nil {
		helper.ResponseError(w, "", "error create token satu sehat", err.Error(), 400, path)
		return
	}

	param := r.URL.Query()

	patientService := services.NewSatuSehatPatient(db, r)
	data, err := patientService.GetDataPatientByNIK(param.Get("nik"), token)
	if err != nil {
		helper.ResponseError(w, id, "failed fetch data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: data})
	helper.ResponseSuccess(w, id, "success get patient id (satu-sehat)", "success get patient id (satu-sehat)", s, 200)
}
