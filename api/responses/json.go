package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	resp_data := map[string]interface{}{
		"success":   true,
		"data": data,
	}
	err := json.NewEncoder(w).Encode(resp_data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
func LOGINJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
func ERRORJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	resp_data := map[string]interface{}{
		"error": data,
	}
	err := json.NewEncoder(w).Encode(resp_data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {

		ERRORJSON(w, statusCode, struct {
			Error string `json:"message"`
			Code int `json:"status_code"`
		}{
			Error: err.Error(),
			Code: statusCode,
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
