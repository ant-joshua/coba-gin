package service

import (
	"coba-gin/model"
	"fmt"
)

// CarService is a service for managing cars.

type CarService struct {
}

func NewCarService() *CarService {
	return &CarService{}
}

func (s *CarService) GetCar(carId string) model.Car {
	var carModel model.Car

	for _, car := range model.CarData {
		if car.CarID == carId {
			carModel = car
		}
	}

	return carModel
}

func (s *CarService) GetCars() []model.Car {
	return model.CarData
}

func (s *CarService) CreateCar(request model.Car) model.Car {

	var carId = fmt.Sprintf("c%d", len(model.CarData)+1)

	request.CarID = carId

	model.CarData = append(model.CarData, request)

	return request
}

func (s *CarService) UpdateCar(carId string, request model.Car) model.Car {
	var updateCar model.Car

	for index, car := range model.CarData {
		if car.CarID == carId {
			request.CarID = carId
			model.CarData[index] = request
			updateCar = model.CarData[index]
		}
	}

	return updateCar
}

func (s *CarService) DeleteCar() bool {
	return true
}
