package controller

import (
	"coba-gin/model"
	"coba-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *CarController) CarRoutes(r *gin.Engine) {
	r.POST("cars", c.CreateCar)
	r.GET("cars", c.GetAllCars)
	r.GET("cars/:car_id", c.GetCar)
	r.PUT("cars/:car_id", c.UpdateCar)
	r.DELETE("cars/:car_id", c.DeleteCar)
}

type CarController struct {
	service *service.CarService
}

func NewCarController() *CarController {
	return &CarController{
		service: service.NewCarService(),
	}
}

func (c *CarController) CreateCar(ctx *gin.Context) {
	var car model.Car
	err := ctx.BindJSON(&car)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		//ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := c.service.CreateCar(car)

	//car.CarID = fmt.Sprintf("c%d", len(model.CarData)+1)
	//model.CarData = append(model.CarData, car)

	//carData = append(carData, car)
	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "success",
		"data":       result,
	})
}

func (c *CarController) GetAllCars(ctx *gin.Context) {

	result := c.service.GetCars()

	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "success get all cars",
		"data":       result,
	})
}

func (c *CarController) GetCar(ctx *gin.Context) {
	carID := ctx.Param("car_id")

	for _, car := range model.CarData {
		if car.CarID == carID {
			ctx.JSON(200, gin.H{
				"statusCode": 200,
				"message":    "success get car",
				"data":       car,
			})
			return
		}
	}

	ctx.JSON(404, gin.H{
		"statusCode": 404,
		"message":    "car not found",
		"data":       nil,
	})
}

func (c *CarController) UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("car_id")

	for index, car := range model.CarData {
		if car.CarID == carID {
			var newCar model.Car
			err := ctx.BindJSON(&newCar)

			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			newCar.CarID = carID

			model.CarData[index] = newCar

			ctx.JSON(200, gin.H{
				"statusCode": 200,
				"message":    "success update car",
				"data":       newCar,
			})
			return
		}
	}

	ctx.JSON(404, gin.H{
		"statusCode": 404,
		"message":    "car not found",
		"data":       nil,
	})
}

func (c *CarController) DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("car_id")

	for index, car := range model.CarData {
		if car.CarID == carID {
			model.CarData = append(model.CarData[:index], model.CarData[index+1:]...)

			ctx.JSON(200, gin.H{
				"statusCode": 200,
				"message":    "success delete car",
				"data":       nil,
			})
			return
		}
	}

	ctx.JSON(404, gin.H{
		"statusCode": 404,
		"message":    "car not found",
		"data":       nil,
	})
}

//var cars = []Car{
//	{CarID: 1, Brand: "Toyota", Price: 100000000, Model: "Avanza"},
//	{CarID: 2, Brand: "Toyota", Price: 200000000, Model: "Innova"},
//	{CarID: 3, Brand: "Toyota", Price: 300000000, Model: "Fortuner"},
//}
