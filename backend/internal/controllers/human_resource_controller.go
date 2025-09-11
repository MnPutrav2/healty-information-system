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

func AddEmployeeData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path, "Admin")
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	case "not_allowed":
		helper.ResponseWarn(w, "", "you are not allowed to access this resource", "resource not allowed", 400, path)
		return
	}
	// ---- Needed for every request ---

	exam, err := helper.GetHumanResourceRequest(w, r)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	examRepo := repository.NewHumanResourceRepository(sql)
	if err := examRepo.AddEmployeeData(exam); err != nil {
		helper.ResponseError(w, val.Id, "failed create data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseWarn(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "create data success", path, s, 201)
}

func GetEmployeeData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path, "User")
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	case "not_allowed":
		helper.ResponseWarn(w, "", "you are not allowed to access this resource", "resource not allowed", 400, path)
		return
	}
	// ---- Needed for every request ---

	param := r.URL.Query()
	search := "%" + param.Get("search") + "%"
	limit, _ := strconv.Atoi(param.Get("limit"))

	examRepo := repository.NewHumanResourceRepository(sql)
	res, err := examRepo.GetEmployeeData(search, limit)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed create data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get data success", path, s, 200)
}

func DeleteEmployeeData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path, "Admin")
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	case "not_allowed":
		helper.ResponseWarn(w, "", "you are not allowed to access this resource", "resource not allowed", 400, path)
		return
	}
	// ---- Needed for every request ---

	param := r.URL.Query()
	id := param.Get("id")

	examRepo := repository.NewHumanResourceRepository(sql)
	if err := examRepo.DeleteEmployeeData(id); err != nil {
		helper.ResponseError(w, val.Id, "failed create data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "delete data success", path, s, 200)
}

func UpdateEmployeeData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path, "Admin")
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	case "not_allowed":
		helper.ResponseWarn(w, "", "you are not allowed to access this resource", "resource not allowed", 400, path)
		return
	}
	// ---- Needed for every request ---

	param := r.URL.Query()
	id := param.Get("id")

	exam, err := helper.GetHumanResourceRequest(w, r)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	examRepo := repository.NewHumanResourceRepository(sql)
	if err := examRepo.UpdateEmployeeData(exam, id); err != nil {
		helper.ResponseError(w, val.Id, "failed create data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "updated"})
	if err != nil {
		helper.ResponseWarn(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "update data success", path, s, 200)
}

func AddUserLogin(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
	// ---- Needed for every request ---

	val := pkg.CheckUserLogin(w, r, sql, path, "Admin")
	switch val.Status {
	case "authorization":
		helper.ResponseWarn(w, "", "unauthorization", "unauthorization", 401, path)
		return
	case "error_format":
		helper.ResponseWarn(w, "", "unauthorization error format", "unauthorization error format", 400, path)
		return
	case "not_allowed":
		helper.ResponseWarn(w, "", "you are not allowed to access this resource", "resource not allowed", 400, path)
		return
	}
	// ---- Needed for every request ---

	exam, err := helper.GetUserRequest(w, r)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	examRepo := repository.NewHumanResourceRepository(sql)
	if err := examRepo.AddUserLogin(exam); err != nil {
		helper.ResponseError(w, val.Id, "failed create data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "add"})
	if err != nil {
		helper.ResponseWarn(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "add user login success", path, s, 201)
}
