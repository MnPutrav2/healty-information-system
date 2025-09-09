package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
	"github.com/MnPutrav2/healty-information-system/backend/internal/repository"
)

func GetCurrentMR(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	common := repository.NewCommonRepository(w, r, sql)
	mr := common.GetCurrentMedicalRecord()

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: mr})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get current MR", path, s, 200)
}

func GetCurrentRegisterNum(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	date := query.Get("date")
	policlinic := query.Get("policlinic")

	common := repository.NewCommonRepository(w, r, sql)
	data, err := common.GetCurrentRegisterNumber(date, policlinic)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get data", err.Error(), 401, path)
	}

	s, err := json.Marshal(models.ResponseDataSuccessInt{Status: "success", Response: data})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get current reg number", path, s, 200)
}

func GetCurrentCareNum(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	date := query.Get("date")

	common := repository.NewCommonRepository(w, r, sql)
	data, err := common.GetCurrentCareNumber(date)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get data", err.Error(), 401, path)
	}

	s, err := json.Marshal(models.ResponseDataSuccessInt{Status: "success", Response: data})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get current care number", path, s, 200)
}

func GetPoliclinics(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	common := repository.NewCommonRepository(w, r, sql)
	status, err := common.GetPoliclinic()
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get policlinic", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(status)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get data", path, s, 200)
}

func AddRecipeNumber(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	var care = param.Get("care_number")

	common := repository.NewCommonRepository(w, r, sql)
	res, err := common.AddRecipeNumber(care)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get recipe number", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: res})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get recipe number", path, s, 200)
}

func GetCurrentRecipeNumber(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	var date = param.Get("date")

	common := repository.NewCommonRepository(w, r, sql)
	res, err := common.GetCurrentRecipeNumber(date)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get recipe number", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccessInt{Status: "success", Response: res})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get current recipe number", path, s, 200)
}

func GetCurrentLabNumber(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	var date = param.Get("date")

	common := repository.NewCommonRepository(w, r, sql)
	res, err := common.GetCurrentLabNumber(date)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get lab number", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccessInt{Status: "success", Response: res})
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get current lab number", path, s, 200)
}
