package dal

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

type Car struct {
	ID       int    `json:"id"`
	Year     int    `json:"year"`
	Make     string `json:"make"`	
	Model    string `json:"model"`
	Trim     string `json:"trim"`
	ImageURL string `json:"image_url"`
}

func SetCollection(c *mongo.Collection) {
	collection = c
}

func GetAllCars() ([]Car, error) {
	var cars []Car

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var car Car
		if err := cur.Decode(&car); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func CreateCar(car *Car) error {
	// Check if the car already exists
	existingCar, err := GetCar(car.ID)
	if err == nil && existingCar != nil {
		return fmt.Errorf("car with ID %d already exists", car.ID)
	}

	_, err = collection.InsertOne(context.Background(), car)
	if err != nil {
		return err
	}

	return nil
}

func GetCar(carID int) (*Car, error) {
	var car Car
	err := collection.FindOne(context.Background(), bson.M{"id": carID}).Decode(&car)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Car not found
		}
		return nil, err
	}
	return &car, nil
}

func UpdateCar(carID int, updatedCar *Car) error {
	_, err := collection.UpdateOne(context.Background(), bson.M{"id": carID}, bson.M{"$set": updatedCar})
	if err != nil {
		return err
	}
	return nil
}

func DeleteCar(carID int) error {
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": carID})
	if err != nil {
		return err
	}
	return nil
}
