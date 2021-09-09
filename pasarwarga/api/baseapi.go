package api

import (
"encoding/json"
"log"
"net/http"
)

type BaseApi struct {
}

func (baseApi BaseApi) Error(w http.ResponseWriter, err error) {

	log.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	data := map[string]string{
		"error": err.Error(),
	}

	err1 := json.NewEncoder(w).Encode(data)
	if err1 != nil {
		log.Println(err1)
	}
}

func (baseApi BaseApi) QueryParam(r *http.Request, key string) string {
	q := r.URL.Query()[key]
	if len(q) > 0 {
		return q[0]
	}
	return ""
}

func (baseApi BaseApi) Json(w http.ResponseWriter, data interface{}, status int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
		baseApi.Error(w, err)
	}
}

func (baseApi BaseApi) Empty(w http.ResponseWriter, status int) {

	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(status)
	w.Write([]byte{})
}

