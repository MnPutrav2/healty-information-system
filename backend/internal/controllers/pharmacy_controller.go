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

func CreateDrugDatas(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	patient, err := helper.GetRequestBodyDrugData(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	if err := pharmacyRepo.CreateDrugData(patient); err != nil {
		helper.ResponseError(w, val.Id, "failed create drug data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseError(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "create drug data", path, s, 201)
}

func GetDrugDatas(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	var search = "%" + param.Get("search") + "%"
	limit, _ := strconv.Atoi(param.Get("limit"))

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	data, err := pharmacyRepo.GetDrugData(search, limit)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get drug data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(data)
	if err != nil {
		helper.ResponseError(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "get drug data", path, s, 200)
}

func UpdateDrugDatas(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	drug, err := helper.GetRequestBodyDrugDataUpdate(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	if err := pharmacyRepo.UpdateDrugData(drug); err != nil {
		helper.ResponseError(w, val.Id, "failed create drug data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "updated"})
	if err != nil {
		helper.ResponseError(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "update drug data", path, s, 200)
}

func DeleteDrugDatas(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	var drug = param.Get("drug_id")

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	if err := pharmacyRepo.DeleteDrugData(drug); err != nil {
		helper.ResponseError(w, val.Id, "failed delete drug data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	if err != nil {
		helper.ResponseError(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "success delete drug data", path, s, 200)
}

func GetDistributor(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	data, err := pharmacyRepo.GetDistributor()
	if err != nil {
		helper.ResponseError(w, val.Id, "failed distributor data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get distributor data", path, s, 200)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	patient, err := helper.GetRequestBodyDrugRecipe(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	res, err := pharmacyRepo.CreateRecipe(patient)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed create recipe", err.Error(), 404, path)
		return
	}

	if res == "duplicate" {
		helper.ResponseWarn(w, val.Id, "duplicate entry", "duplicate entry", 500, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "created"})
	if err != nil {
		helper.ResponseError(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "create recipe success", path, s, 201)
}

func CreateRecipeCompound(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	patient, err := helper.GetRequestBodyDrugRecipeCompound(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	stat, err := pharmacyRepo.CreateRecipeCompound(patient)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed create recipe", err.Error(), 404, path)
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

	helper.ResponseSuccess(w, val.Id, "create recipe success", path, s, 201)
}

func GetRecipes(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	data, err := pharmacyRepo.GetRecipes(date1, date2)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get recipes data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get recipes data", path, s, 200)
}

func GetDrugRecipes(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	var recipe = param.Get("recipe")

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	data, err := pharmacyRepo.GetDrugRecipes(recipe)
	if err != nil {
		helper.ResponseError(w, val.Id, "failed get recipe data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	helper.ResponseSuccess(w, val.Id, "get all drug recipe data", path, s, 200)
}

func DeleteDrugRecipes(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	var recipe = param.Get("recipe_id")
	var drug = param.Get("drug_id")
	var comname = param.Get("compound_name")

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	if err := pharmacyRepo.DeleteDrugRecipes(recipe, drug, comname); err != nil {
		helper.ResponseError(w, val.Id, "failed delete drug data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	if err != nil {
		helper.ResponseError(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "success delete drug data", path, s, 200)
}

func ValidateRecipe(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	var recipe = param.Get("recipe_id")

	validate, err := helper.GetRequestBodyDrugRecipeValidate(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	if err := pharmacyRepo.ValidateRecipe(recipe, validate); err != nil {
		helper.ResponseError(w, val.Id, "failed validate recipe", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "deleted"})
	if err != nil {
		helper.ResponseError(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "success validate recipe", path, s, 200)
}

func RecipeHandover(w http.ResponseWriter, r *http.Request, sql *sql.DB, path string) {
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
	var recipe = param.Get("recipe-id")

	res, err := helper.GetRequestBodyDrugRecipeHandover(w, r, path)
	if err != nil {
		helper.ResponseError(w, val.Id, "invalid request body", err.Error(), 400, path)
		return
	}

	pharmacyRepo := repository.NewPharmacyRepository(sql)
	if err := pharmacyRepo.HandoverDrug(recipe, res); err != nil {
		helper.ResponseError(w, val.Id, "failed get recipe data", err.Error(), 404, path)
		return
	}

	s, err := json.Marshal(models.ResponseDataSuccess{Status: "success", Response: "updated"})
	if err != nil {
		helper.ResponseError(w, val.Id, "error server", err.Error(), 500, path)
		return
	}

	helper.ResponseSuccess(w, val.Id, "success handover recipe", path, s, 200)
}
