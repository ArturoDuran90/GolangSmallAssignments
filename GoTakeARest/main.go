package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"carsApi.com/dal"
)

var port = ":8000"
var collection *mongo.Collection

type CarResponse struct {
	Message string `json:"message"`
}

type Car dal.Car

var uri = "mongodb+srv://aduran:Nu191036673@cluster0.lutxvb3.mongodb.net/"

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("carsdb")
	collection = db.Collection("cars")

	dal.SetCollection(collection)

	app := echo.New()

	app.Static("/images", "images")

	app.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, CarResponse{Message: "Welcome to my API. Try our routes"})
	})
	app.POST("/car/create", createCar)
	app.GET("/car/:id", getCar)
	app.GET("/car", getAllCars)
	app.PUT("/car/:id", updateCar)
	app.DELETE("/car/:id", deleteCar)

	app.Logger.Fatal(app.Start(port))
}

func getAllCars(ctx echo.Context) error {
	cars, err := dal.GetAllCars()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, CarResponse{Message: "Error retrieving cars"})
	}
	return ctx.JSON(http.StatusOK, cars)
}

func createCar(ctx echo.Context) error {
	car := new(Car)
	if err := ctx.Bind(car); err != nil {
		return ctx.JSON(http.StatusBadRequest, CarResponse{Message: "Invalid request payload"})
	}

	err := dal.CreateCar((*dal.Car)(car))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, CarResponse{Message: "Error creating car"})
	}

	return ctx.JSON(http.StatusCreated, car)
}

func getCar(ctx echo.Context) error {
	carID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, CarResponse{Message: "Invalid Car ID"})
	}

	car, err := dal.GetCar(carID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, CarResponse{Message: "Car not found"})
	}

	return ctx.JSON(http.StatusOK, car)
}

func updateCar(ctx echo.Context) error {
	carID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, CarResponse{Message: "Invalid Car ID"})
	}

	updatedCar := new(dal.Car)
	if err := ctx.Bind(updatedCar); err != nil {
		return ctx.JSON(http.StatusBadRequest, CarResponse{Message: "Invalid request payload"})
	}

	err = dal.UpdateCar(carID, updatedCar)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, CarResponse{Message: "Error updating car"})
	}

	return ctx.JSON(http.StatusOK, updatedCar)
}

func deleteCar(ctx echo.Context) error {
	carID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, CarResponse{Message: "Invalid Car ID"})
	}

	err = dal.DeleteCar(carID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, CarResponse{Message: "Error deleting car"})
	}

	return ctx.JSON(http.StatusOK, CarResponse{Message: "Car deleted"})
}
