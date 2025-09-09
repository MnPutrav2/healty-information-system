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

func AddExamination(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	exam, err := helper.GetExaminationRequest(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	examRepo := repository.NewExaminationRepository(sql)
	err = examRepo.AddExamination(exam)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed add data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	helper.ResponseSuccess(w, val.Id, "add examination", path, s, 201)
}

func GetExamination(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	id := param.Get("care-number")

	examRepo := repository.NewExaminationRepository(sql)
	res, err := examRepo.GetExamination(id)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed add data", err.Error(), 400, path)
		return
	}

	s, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get examination", path, s, 200)
}

func DeleteExamination(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	examRepo := repository.NewExaminationRepository(sql)
	if err := examRepo.DeleteExamination(id); err != nil {
		helper.ResponseError(w, val.Id, "failed add data", err.Error(), 400, path)
		return
	}

	s, _ := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	helper.ResponseSuccess(w, val.Id, "delete examination", path, s, 200)
}
