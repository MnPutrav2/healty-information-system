package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
)

func AuthUser(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {

	// get client request body
	account, err := helper.GetRequestBodyUserAccount(w, r, path)
	if err != nil {
		helper.ResponseWarn(w, "", "invalid request body", err.Error(), 400, path)
		return
	}

	// check user
	// check user available
	// var id int
	// err = sql.QueryRow("SELECT COUNT(*) FROM users WHERE users.username = $1 AND users.password = $2", account.Username, account.Password).Scan(&id)
	// if err != nil {
	// 	helper.ResponseError(w, "", "Internal server error", err.Error(), 500, path)
	// 	return
	// }

	// // if account not available
	// if id != 1 {
	// 	helper.ResponseWarn(w, "", "Login failed : Check your username or password", "username or password error", 400, path)
	// 	return
	// }

	pepper := os.Getenv("SECRET_STRING")

	var hash string
	var id string

	err = sql.QueryRow("SELECT password, id FROM users WHERE username = $1", account.Username).Scan(&hash, &id)
	if err != nil {
		helper.ResponseError(w, "", "Internal server error", err.Error(), 500, path)
		return
	}

	if pkg.CheckPassword(hash, account.Password+pepper) {
		s, err := json.Marshal(models.AuthResponse{Status: "success", Token: pkg.SessionToken(sql, id)})

		if err != nil {
			helper.ResponseError(w, "", "Internal server error", err.Error(), 500, path)
			return
		}

		helper.ResponseSuccess(w, "", "client login", path, s, 201)
	} else {
		helper.ResponseWarn(w, "", "Login failed : Check your username or password", "username or password error", 400, path)
	}

}

func AuthPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {

	// get client request body
	account, err := helper.GetRequestBodyUserAccount(w, r, path)
	if err != nil {
		helper.ResponseWarn(w, "", "invalid request body", err.Error(), 400, path)
		return
	}

	// check user
	// check user available
	var id int
	err = sql.QueryRow("SELECT COUNT(*) FROM patient_account WHERE patient_account.username = $1 AND patient_account.password = $2", account.Username, account.Password).Scan(&id)
	if err != nil {
		helper.ResponseError(w, "", "Internal server error", err.Error(), 500, path)
		return
	}

	// if account not available
	if id != 1 {
		helper.ResponseWarn(w, "", "Login failed : Check your username or password", "username or password error", 400, path)
		return
	}

	// success
	s, err := json.Marshal(models.AuthResponse{Status: "success", Token: pkg.PatientSessionToken(sql, account.Username, account.Password)})
	if err != nil {
		helper.ResponseError(w, "", "Internal server error", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, "", "client login", path, s, 201)
}
