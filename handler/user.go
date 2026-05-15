package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saurav11sarkar/resapi/internal/model"
	"github.com/saurav11sarkar/resapi/utils"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	var body model.User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendJson(w, http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}
	body.Id = nextId
	nextId++
	users = append(users, body)

	utils.SendJson(w, http.StatusCreated, model.Response{
		Success: true,
		Message: "User create success",
		Data:    body,
	})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	utils.SendJson(w, http.StatusOK, model.Response{
		Success: true,
		Message: "User get success",
		Data:    users,
	})
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendJson(w, http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid user id",
			Data:    nil,
		})
		return
	}

	for _, user := range users {
		if user.Id == id {
			utils.SendJson(w, http.StatusOK, model.Response{
				Success: true,
				Message: "User get success",
				Data:    user,
			})
			return
		}
	}

	utils.SendJson(w, http.StatusNotFound, model.Response{
		Success: false,
		Message: "User not found",
		Data:    nil,
	})
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendJson(w, http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid user id",
			Data:    nil,
		})
		return
	}

	var body model.User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendJson(w, http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	for i, user := range users {
		if user.Id == id {
			body.Id = id
			users[i] = body
			utils.SendJson(w, http.StatusOK, model.Response{
				Success: true,
				Message: "User update success",
				Data:    users[i],
			})
			return
		}
	}

	utils.SendJson(w, http.StatusNotFound, model.Response{
		Success: false,
		Message: "User not found",
		Data:    nil,
	})
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendJson(w, http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid user id",
			Data:    nil,
		})
		return
	}

	var newUsers []model.User
	found := false
	for _, user := range users {
		if user.Id == id {
			found = true
			continue
		}
		newUsers = append(newUsers, user)
	}

	if !found {
		utils.SendJson(w, http.StatusNotFound, model.Response{
			Success: false,
			Message: "User not found",
			Data:    nil,
		})
		return
	}

	users = newUsers
	utils.SendJson(w, http.StatusOK, model.Response{
		Success: true,
		Message: "User delete success",
		Data:    nil,
	})
}
