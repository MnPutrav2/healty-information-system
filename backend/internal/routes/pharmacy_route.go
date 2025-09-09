package routes

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/controllers"
	"github.com/MnPutrav2/healty-information-system/backend/internal/pkg"
)

func PharmacyDrugData(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetDrugDatas(w, r, db, path)

	case "POST":
		controllers.CreateDrugDatas(w, r, db, path)

	case "PUT":
		controllers.UpdateDrugDatas(w, r, db, path)

	case "DELETE":
		controllers.DeleteDrugDatas(w, r, db, path)
	}

}

func Recipe(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetRecipes(w, r, db, path)

	case "POST":
		controllers.CreateRecipe(w, r, db, path)

	case "DELETE":
		controllers.DeleteDrugRecipes(w, r, db, path)
	}

}

func RecipeCompound(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetDrugRecipes(w, r, db, path)

	case "POST":
		controllers.CreateRecipeCompound(w, r, db, path)
	}

}

func Distributor(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "GET":
		controllers.GetDistributor(w, r, db, path)
	}

}

func DrugValidate(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "PUT":
		controllers.ValidateRecipe(w, r, db, path)
	}

}

func DrugHandover(w http.ResponseWriter, r *http.Request, db *sql.DB, path string) {

	if !pkg.RequestHeader(w, r) {
		return
	}

	switch r.Method {
	case "PUT":
		controllers.RecipeHandover(w, r, db, path)
	}

}
