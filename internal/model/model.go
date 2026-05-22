package model

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type PublicUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type LoginResponse struct {
	AccessToken string     `json:"accessToken"`
	TokenType   string     `json:"tokenType"`
	ExpiresIn   int64      `json:"expiresIn"`
	User        PublicUser `json:"user"`
}

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Discription string `json:"discription"`
	Price       int    `json:"price"`
	Images      string `json:"images"`
	Status      int    `json:"status"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToPublicUser(user User) PublicUser {
	return PublicUser{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}
}

func ToPublicUsers(users []User) []PublicUser {
	publicUsers := make([]PublicUser, 0, len(users))
	for _, user := range users {
		publicUsers = append(publicUsers, ToPublicUser(user))
	}
	return publicUsers
}
