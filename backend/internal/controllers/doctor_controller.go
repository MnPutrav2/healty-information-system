package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
	"github.com/MnPutrav2/healty-information-system/backend/internal/repository"
)

func GetDoctorSchedule(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	id := param.Get("id")

	doctorRepo := repository.NewDoctorRepository(sql)
	status, err := doctorRepo.GetDoctorSchedule(id)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get doctor schedule", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(status)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get schedule", path, s, 200)
}

func GetDoctorData(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	if _, ok := param["specialis"]; !ok {
		doctorRepo := repository.NewDoctorRepository(sql)
		status, err := doctorRepo.GetDoctorsAll()
		if err != nil {
			helper.ResponseError(w, val.Id, "failed get doctor data", err.Error(), 404, path)
			return
		}

		s, err := json.Marshal(status)
		if err != nil {
			panic(err.Error())
		}

		helper.ResponseSuccess(w, val.Id, "get data", path, s, 200)
	} else {
		id := param.Get("specialis")

		doctorRepo := repository.NewDoctorRepository(sql)
		status, err := doctorRepo.GetDoctors(id)
		if err != nil {
			helper.ResponseError(w, val.Id, "failed get doctor data", err.Error(), 404, path)
			return
		}

		s, err := json.Marshal(status)
		if err != nil {
			panic(err.Error())
		}

		helper.ResponseSuccess(w, val.Id, "get data", path, s, 200)
	}
}
