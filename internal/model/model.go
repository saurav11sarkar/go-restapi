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
