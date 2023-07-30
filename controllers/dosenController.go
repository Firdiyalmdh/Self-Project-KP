package controllers

import (
	"context"
	"golang/configs"
	"golang/models"
	"golang/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dsnCollection *mongo.Collection = configs.GetCollection(configs.DB, "dosen")

// var validate = validator.New()

func GetOneDsn(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var dsn models.Dosen
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	err := dsnCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&dsn)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DosenResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusOK, responses.DosenResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": dsn},
	})
}

func GetAllDsn(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var allDsn []models.Dosen
	defer cancel()

	results, err := dsnCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DosenResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var dsn models.Dosen
		if err := results.Decode(&dsn); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DosenResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}
		allDsn = append(allDsn, dsn)
	}

	return c.JSON(http.StatusOK, responses.DosenResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": allDsn},
	})
}

func EditADsn(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var dsn models.Dosen
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	if err := c.Bind(&dsn); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DosenResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	if validationErr := validate.Struct(&dsn); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.DosenResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": validationErr.Error()},
		})
	}

	update := bson.M{
		"nama":     dsn.Nama,
		"nrp":      dsn.NIP,
		"email":    dsn.Email,
		"password": dsn.Password,
	}

	result, err := dsnCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DosenResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	var updatedDsn models.Dosen
	if result.MatchedCount == 1 {
		err := dsnCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedDsn)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DosenResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}
	}

	return c.JSON(http.StatusOK, responses.DosenResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": updatedDsn},
	})
}

func CreateDsn(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var dsn models.Dosen
	defer cancel()

	if err := c.Bind(&dsn); err != nil {
		return c.JSON(http.StatusBadRequest, responses.DosenResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	newDsn := models.Dosen{
		Id:       primitive.NewObjectID(),
		Nama:     dsn.Nama,
		NIP:      dsn.NIP,
		Email:    dsn.Email,
		Password: dsn.Password,
	}

	result, err := dsnCollection.InsertOne(ctx, newDsn)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DosenResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusCreated, responses.DosenResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    &echo.Map{"data": result},
	})
}

func DeleteDsn(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := dsnCollection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DosenResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.DosenResponse{
			Status:  http.StatusNotFound,
			Message: "Error",
			Data:    &echo.Map{"data": "Dosen with specified ID not found"},
		})
	}

	return c.JSON(http.StatusOK, responses.DosenResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": "Dosen with specified ID successfully deleted"},
	})
}
