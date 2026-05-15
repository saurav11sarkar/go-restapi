package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/saurav11sarkar/resapi/internal/model"
)

func SendJson(w http.ResponseWriter, status int, data model.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println(err)
	}
}
