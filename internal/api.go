package cryo

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taudelta/talky/internal/storage"
)

type Dependencies struct {
	Database   *storage.PostgreSQLStorage
	PhraseRepo storage.PhraseRepository
}

type ApiV1 struct {
	database   *storage.PostgreSQLStorage
	phraseRepo storage.PhraseRepository
}

type ApiError struct {
	Error string
}

func errorResponse(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")

	apiError := ApiError{
		Error: err.Error(),
	}

	_ = json.NewEncoder(w).Encode(&apiError)
}

func successResponse(w http.ResponseWriter, s interface{}) {
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(s)
}

type HealthcheckResult struct {
	Database string
}

type HealthcheckResponse struct {
	Healthy bool
	Result  HealthcheckResult
}

func (a *ApiV1) Healthcheck(w http.ResponseWriter, r *http.Request) {

	var databaseError error
	healthy := true

	result := HealthcheckResult{}

	databaseError = a.database.Ping()
	if databaseError != nil {
		healthy = false
		result.Database = databaseError.Error()
	}

	response := HealthcheckResponse{
		Healthy: healthy,
		Result:  result,
	}

	_ = json.NewEncoder(w).Encode(&response)
}

func InitApiV1(r *mux.Router, deps Dependencies) *ApiV1 {

	api := &ApiV1{
		database:   deps.Database,
		phraseRepo: deps.PhraseRepo,
	}

	r.HandleFunc("/healthcheck", api.Healthcheck)
	r.HandleFunc("/v1/phrases", api.Find).Methods("POST")

	return api
}
