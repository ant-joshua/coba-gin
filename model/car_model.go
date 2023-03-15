package model

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Price int    `json:"price"`
	Model string `json:"model"`
}

var CarData []Car
