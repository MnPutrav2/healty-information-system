package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
	"github.com/MnPutrav2/healty-information-system/backend/internal/repository"
)

func GetLogs(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string, m string) {
	// ---- Needed for every request ---
	if !pkg.CheckRequestHeader(w, r, sql, path, m) {
		return
	}

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
	date1 := param.Get("date1")
	date2 := param.Get("date2")

	logRepo := repository.NewLogsRepository(sql)
	dt, err := logRepo.GetLogsData(date1, date2)
	if err != nil {
		panic(err.Error())
	}

	s, _ := json.Marshal(dt)

	w.Write(s)
}
