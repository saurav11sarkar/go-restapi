package hendler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/saurav11sarkar/resapi/internal/model"
	"github.com/saurav11sarkar/resapi/utils"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret string
var jwtExpires time.Duration

func ConfigureAuth(secret string, expires time.Duration) {
	jwtSecret = secret
	jwtExpires = expires
}

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

	for _, user := range users {
		if user.Email == body.Email {
			utils.SendJson(w, http.StatusConflict, model.Response{
				Success: false,
				Message: "Email already exists",
				Data:    nil,
			})
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendJson(w, http.StatusInternalServerError, model.Response{
			Success: false,
			Message: "Could not create user",
			Data:    nil,
		})
		return
	}
	body.Password = string(hashedPassword)

	body.Id = nextId
	nextId++
	users = append(users, body)

	utils.SendJson(w, http.StatusCreated, model.Response{
		Success: true,
		Message: "User create success",
		Data:    model.ToPublicUser(body),
	})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	utils.SendJson(w, http.StatusOK, model.Response{
		Success: true,
		Message: "User get success",
		Data:    model.ToPublicUsers(users),
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
				Data:    model.ToPublicUser(user),
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
			if body.Password == "" {
				body.Password = user.Password
			} else {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
				if err != nil {
					utils.SendJson(w, http.StatusInternalServerError, model.Response{
						Success: false,
						Message: "Could not update user",
						Data:    nil,
					})
					return
				}
				body.Password = string(hashedPassword)
			}
			body.Role = user.Role
			body.Id = id
			users[i] = body
			utils.SendJson(w, http.StatusOK, model.Response{
				Success: true,
				Message: "User update success",
				Data:    model.ToPublicUser(users[i]),
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

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var body model.Login

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendJson(w, http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	for _, user := range users {
		if user.Email == body.Email && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) == nil {
			token, err := utils.GenerateToken(user.Id, user.Email, user.Role, jwtSecret, jwtExpires)
			if err != nil {
				utils.SendJson(w, http.StatusInternalServerError, model.Response{
					Success: false,
					Message: "Could not create access token",
					Data:    nil,
				})
				return
			}

			utils.SendJson(w, http.StatusOK, model.Response{
				Success: true,
				Message: "User login success",
				Data: model.LoginResponse{
					AccessToken: token,
					TokenType:   "Bearer",
					ExpiresIn:   int64(jwtExpires.Seconds()),
					User:        model.ToPublicUser(user),
				},
			})
			return
		}
	}

	utils.SendJson(w, http.StatusUnauthorized, model.Response{
		Success: false,
		Message: "Invalid email or password",
		Data:    nil,
	})
}
