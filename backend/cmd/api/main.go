package main

import (
	"database/sql"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/config"
	"github.com/MnPutrav2/healty-information-system/backend/internal/controllers"
	"github.com/MnPutrav2/healty-information-system/backend/internal/helper"
	"github.com/MnPutrav2/healty-information-system/backend/internal/routes"
)

func main() {
	db := config.SqlDb()

	defer db.Close()

	// User API
	handlerPath("/user/auth", routes.AuthUser, db)
	handlerPath("/user/status", routes.UserStatus, db)
	handlerPath("/user/pages", routes.UserPages, db)
	handlerPath("/user/logout", routes.UserLogout, db)

	// Patient API
	handlerPath("/patient", routes.Patient, db)
	handlerPath("/patient/status", routes.PatientStatus, db)
	handlerPath("/patient/auth", routes.PatientAuth, db)

	// Registration API
	handlerPath("/registration", routes.Registration, db)
	handlerPath("/registration/update-status", routes.RegistrationUpdateStatus, db)

	// Finance API
	handlerPath("/finance/examination-cost", routes.Examination, db)

	// Ambulatory API
	handlerPath("/ambulatory", routes.Ambulatory, db)

	// Laboratorium
	handlerPath("/laboratorium/data", routes.LaboratoriumData, db)
	handlerPath("/laboratorium/template", routes.LaboratoriumTemplate, db)
	handlerPath("/laboratorium/request", routes.LaboratoriumRequest, db)
	handlerPath("/laboratorium/detail", routes.LaboratoriumDetail, db)
	handlerPath("/laboratorium/value", routes.LaboratoriumValue, db)

	// Pharmacy API
	handlerPath("/pharmacy/drug-data", routes.PharmacyDrugData, db)
	handlerPath("/pharmacy/distributor", routes.Distributor, db)
	handlerPath("/pharmacy/validate", routes.DrugValidate, db)
	handlerPath("/pharmacy/handover", routes.DrugHandover, db)
	handlerPath("/pharmacy/recipe", routes.Recipe, db)
	handlerPath("/pharmacy/recipe-compound", routes.RecipeCompound, db)

	// Common API
	handlerPath("/common/current-medical-record", routes.CurrentMedicalRecord, db)
	handlerPath("/common/current-register-number", routes.CurrentRegisterNumber, db)
	handlerPath("/common/current-care-number", routes.CurrentCareNumber, db)
	handlerPath("/common/policlinic", routes.Policlinic, db)
	handlerPath("/common/recipe-number", routes.RecipeNumber, db)
	handlerPath("/common/current-recipe-number", routes.CurrentRecipeNumber, db)
	handlerPath("/common/current-laboratorium-number", routes.CurrentRecipeNumber, db)

	// Doctor API
	handlerPath("/doctor", routes.DoctorData, db)
	handlerPath("/doctor/schedule", routes.DoctorSchedule, db)

	// Examination API
	handlerPath("/examination", routes.ExaminationCare, db)

	// Satu Sehat API
	handlerPath("/satu-sehat/patient", routes.SatuSehatPatient, db)
	handlerPath("/satu-sehat/encounter", routes.SatuSehatEncounter, db)
	handlerPath("/satu-sehat/condition", routes.SatuSehatCondition, db)
	handlerPath("/satu-sehat/care-plan", routes.SatuSehatCarePlan, db)
	handlerPath("/satu-sehat/clinical-impression", routes.SatuSehatClinicalImpression, db)
	handlerPath("/satu-sehat/observation", routes.SatuSehatObservation, db)

	// Logs
	handler("GET", "/logs", controllers.GetLogs, db)

	helper.LogWorker("[INFO] server runing in port 8080")
	http.ListenAndServe(":8080", nil)
}

func handlerPath(path string, handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, path string), db *sql.DB) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db, path)
	})
}

func handler(m string, path string, handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, path string, m string), db *sql.DB) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db, path, m)
	})
}
