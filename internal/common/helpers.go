package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		RespondWithJSON(w, http.StatusBadRequest, fmt.Sprint("Error marshalling JSON: - ", err))
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, err string) {
	RespondWithJSON(w, code, map[string]string{"error": fmt.Sprintf("%+v", err)})
}

func GetParamsFromRequestBody[T interface{}](structBody T, r *http.Request) (T, error) {
	decoder := json.NewDecoder(r.Body)
	params := structBody
	err := decoder.Decode(&params)
	if err != nil {
		return structBody, NewError("decoding error", err)
	}
	return params, err
}

func NewError(s string, err error) error {
	return fmt.Errorf("%s - %s", s, err.Error())
}
