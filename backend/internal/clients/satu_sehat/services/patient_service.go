package services

import (
	"database/sql"
	"encoding/json"
	"io"
	"os"

	// "io/ioutil"
	"net/http"

	"github.com/MnPutrav2/healty-information-system/backend/internal/models"
	"github.com/joho/godotenv"
)

type SatuSehatPatient interface {
	GetDataPatientByNIK(nik string, token string) (string, error)
}

type satuSehatPatient struct {
	sql *sql.DB
	r   *http.Request
}

func NewSatuSehatPatient(sql *sql.DB, r *http.Request) SatuSehatPatient {
	return &satuSehatPatient{sql, r}
}

func (q *satuSehatPatient) GetDataPatientByNIK(nik string, token string) (string, error) {
	_ = godotenv.Load()

	url := os.Getenv("SATU_SEHAT_END_POINT") + "/Patient?identifier=https://fhir.kemkes.go.id/id/nik|" + nik

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		var data models.PatientBundle
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			return "", err
		}

		return data.Entry[0].Resource.ID, nil
	} else {
		return "failed", err
	}
}
