package hendler

import (
	"net/http"

	"github.com/saurav11sarkar/resapi/internal/model"
	"github.com/saurav11sarkar/resapi/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.SendJson(w, http.StatusOK, model.Response{
		Success: true,
		Message: "Hello World",
		Data:    nil,
	})
}
