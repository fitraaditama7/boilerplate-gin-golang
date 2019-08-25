package helper

import (
	"encoding/json"
	"golang-boilerplate/structs"
	"net/http"
)

func Responses(w http.ResponseWriter, code int, payload interface{}) {
	var result structs.Response

	if code != http.StatusOK {
		result.Error = true
		result.Code = code
		result.Message = "error"
	} else {
		result.Error = false
		result.Code = code
		result.Message = "Success"
	}
	result.Result = payload
	response, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ErrorCustomStatus(w http.ResponseWriter, code int, msg string) {
	Responses(w, code, map[string]string{"error": msg})
}
