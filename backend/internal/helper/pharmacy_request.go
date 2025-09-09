package helper

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
)

func GetRequestBodyDrugData(w http.ResponseWriter, r *http.Request, path string) (models.RequestBodyDrugData, error) {
	var drug models.RequestBodyDrugData

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RequestBodyDrugData{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugDataUpdate(w http.ResponseWriter, r *http.Request, path string) (models.RequestBodyDrugDataUpdate, error) {
	var drug models.RequestBodyDrugDataUpdate

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RequestBodyDrugDataUpdate{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugRecipe(w http.ResponseWriter, r *http.Request, path string) (models.RecipeRequest, error) {
	var drug models.RecipeRequest

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RecipeRequest{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugRecipeCompound(w http.ResponseWriter, r *http.Request, path string) (models.RecipeCompoundRequest, error) {
	var drug models.RecipeCompoundRequest

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RecipeCompoundRequest{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugRecipeValidate(w http.ResponseWriter, r *http.Request, path string) (models.ValidateRecipe, error) {
	var drug models.ValidateRecipe

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.ValidateRecipe{}, err
	}

	return drug, nil
}

func GetRequestBodyDrugRecipeHandover(w http.ResponseWriter, r *http.Request, path string) (models.RecipeHandover, error) {
	var drug models.RecipeHandover

	// Buat decoder dan disallow field yang tidak dikenal
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode JSON langsung ke struct
	err := decoder.Decode(&drug)
	if err != nil {
		return models.RecipeHandover{}, err
	}

	return drug, nil
}
