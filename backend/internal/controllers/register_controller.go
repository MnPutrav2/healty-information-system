package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
	"github.com/MnPutrav2/healty-information-system/backend/internal/repository"
)

func CreateRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path)
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// ---- Needed for every request ---

	// get client request body
	body, err := helper.GetRequestBodyRegisterData(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	err = registerRepo.CreateRegistrationData(body, path)
	if err != nil {
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	helper.ResponseSuccess(w, val.Id, "create patient registration", path, s, 201)
}

func DeleteRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path)
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// ---- Needed for every request ---

	query := r.URL.Query().Get("care-number")

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	err := registerRepo.DeleteRegistrationData(query)
	if err != nil {
		helper.ResponseWarn(w, val.Id, "failed delete data", "failed delete data", 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	helper.ResponseSuccess(w, val.Id, "delete registration", path, s, 200)
}

func GetRegistrationPatient(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	// Check Header
	val := pkg.CheckUserLogin(w, r, sql, path)
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}

	query := r.URL.Query()

	date1 := query.Get("date1")
	date2 := query.Get("date2")
	limit := query.Get("limit")
	search := "%" + query.Get("search") + "%"
	lim, _ := strconv.Atoi(limit)

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	data, err := registerRepo.GetRegistrationData(date1, date2, lim, search)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get data", err.Error(), 401, path)
		return
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get patient data", path, s, 200)
}

func UpdateStatus(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path)
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	}
	// ---- Needed for every request ---

	query := r.URL.Query()
	cn := query.Get("care-number")

	status, err := helper.GetRequestBodyRegisterUpdateStatus(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid json format", err.Error(), 400, path)
		return
	}

	registerRepo := repository.NewRegisterRepository(sql, w, r)
	if err := registerRepo.UpdateStatus(cn, status.Status); err != nil {
		helper.ResponseError(w, val.Id, "failed get data", err.Error(), 401, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	helper.ResponseSuccess(w, val.Id, "delete registration", path, s, 200)
}
