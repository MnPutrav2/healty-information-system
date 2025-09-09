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

func GetLabolatoriumData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	param := r.URL.Query()
	search := "%" + param.Get("search") + "%"
	limit, _ := strconv.Atoi(param.Get("limit"))

	labRepo := repository.NewLabolatoriumRepository(sql)
	data, err := labRepo.GetLabolatoriumData(limit, search)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get labolatorium data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get labolatorium data ", path, s, 200)
}

func GetLabolatoriumTemplate(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	param := r.URL.Query()
	id := param.Get("id")

	labRepo := repository.NewLabolatoriumRepository(sql)
	data, err := labRepo.GetLabolatoriumTemplate(id)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get labolatorium template", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get labolatorium template ", path, s, 200)
}

func CreateLabolatoriumData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	lab, err := helper.GetLabolatoriumResquestBody(w, r)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	labRepo := repository.NewLabolatoriumRepository(sql)
	stat, err := labRepo.CreateLabData(lab)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed create data", err.Error(), 404, path)
		return
	}

	if stat == "duplicate" {
		helper.ResponseWarn(w, val.Id, "duplicate entry", "duplicate entry", 400, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseWarn(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "create data success", path, s, 201)
}

func CreateLabolatoriumRequest(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	lab, err := helper.GetLabolatoriumResquestBodyData(w, r)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	labRepo := repository.NewLabolatoriumRepository(sql)
	stat, err := labRepo.CreateLabolatoriumRequest(lab)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed create data", err.Error(), 404, path)
		return
	}

	if stat == "duplicate" {
		helper.ResponseWarn(w, val.Id, "duplicate entry", "duplicate entry", 400, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseWarn(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "create data success", path, s, 201)
}

func DeleteLabolatoriumData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	param := r.URL.Query()
	id := param.Get("id")

	labRepo := repository.NewLabolatoriumRepository(sql)
	if err := labRepo.DeleteLabolatoriumData(id); err != nil {
		helper.ResponseError(w, val.Id, "failed delete data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	if err != nil {
		helper.ResponseWarn(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "delete data success", path, s, 200)
}

func UpdateLabolatoriumData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	lab, err := helper.UpdateLabolatoriumResquestBody(w, r)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	labRepo := repository.NewLabolatoriumRepository(sql)
	if err := labRepo.UpdateLabolatoriumData(lab); err != nil {
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

func CreateLabolatoriumTemplateData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	lab, err := helper.GetLabolatoriumTemplateResquestBody(w, r)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	labRepo := repository.NewLabolatoriumRepository(sql)
	if err := labRepo.CreateLabTemplate(lab); err != nil {
		helper.ResponseError(w, val.Id, "failed create recipe", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseWarn(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "create data success", path, s, 201)
}

func DeleteLabolatoriumTemplate(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	param := r.URL.Query()
	id := param.Get("id")
	unit := param.Get("unit")

	labRepo := repository.NewLabolatoriumRepository(sql)
	if err := labRepo.DeleteLabolatoriumTemplate(id, unit); err != nil {
		helper.ResponseError(w, val.Id, "failed delete data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	if err != nil {
		helper.ResponseWarn(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "delete data success", path, s, 200)
}

func GetLabolatoriumRequest(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	var param = r.URL.Query()
	var date1 = param.Get("date1")
	var date2 = param.Get("date2")

	pharmacyRepo := repository.NewLabolatoriumRepository(sql)
	res, err := pharmacyRepo.GetLabolatoriumRequest(date1, date2)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get labolatorium data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get labolatorium data ", path, s, 200)
}

func UpdateLabolatoriumSample(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	param := r.URL.Query()
	id := param.Get("id")

	lab, err := helper.UpdateLabolatoriumSampleRequestBody(w, r)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	labRepo := repository.NewLabolatoriumRepository(sql)
	if err := labRepo.UpdateLabSample(lab, id); err != nil {
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

func GetLabolatoriumDataDetail(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	var param = r.URL.Query()
	var id = param.Get("id")

	pharmacyRepo := repository.NewLabolatoriumRepository(sql)
	res, err := pharmacyRepo.GetLabolatoriumDataDetail(id)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get labolatorium data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get labolatorium data ", path, s, 200)
}

func UpdateLabolatoriumValue(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	var param = r.URL.Query()
	var id = param.Get("lab_id")

	lab, err := helper.UpdateLabolatoriumValueRequestBody(w, r)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	labRepo := repository.NewLabolatoriumRepository(sql)
	if err := labRepo.UpdateLabolatoriumValue(id, lab); err != nil {
		helper.ResponseError(w, val.Id, "failed update data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "updated"})
	if err != nil {
		helper.ResponseWarn(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "update labolatorium value", path, s, 200)
}
