package cryo

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/taudelta/talky/internal/dto"
)

func (a *ApiV1) Find(w http.ResponseWriter, r *http.Request) {
	var params dto.FindParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println(err)
		errorResponse(w, err)
		return
	}

	q := a.database.Query()

	phrase, err := a.phraseRepo.Find(q, params)
	if err != nil {
		log.Println(err)
		errorResponse(w, err)
		return
	}

	successResponse(w, phrase)
}
