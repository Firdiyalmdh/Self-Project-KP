package controllers

import (
	"context"
	"golang/configs"
	"golang/models"
	"golang/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mahasiswaCollection *mongo.Collection = configs.GetCollection(configs.DB, "mahasiswa")
var validate = validator.New()

func UpdateMahasiswa(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var mahasiswa models.Mahasiswa
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	if err := c.Bind(&mahasiswa); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	if validationErr := validate.Struct(&mahasiswa); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": validationErr.Error()},
		})
	}

	update := bson.M{
		"nama":     mahasiswa.Nama,
		"nrp":      mahasiswa.NRP,
		"semester": mahasiswa.Semester,
		"email":    mahasiswa.Email,
		"password": mahasiswa.Password,
	}

	result, err := mahasiswaCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	var updatedMahasiswa models.Mahasiswa
	if result.MatchedCount == 1 {
		err := mahasiswaCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedMahasiswa)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": updatedMahasiswa},
	})
}

func GetMahasiswa(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var mahasiswa models.Mahasiswa
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	err := mahasiswaCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&mahasiswa)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": mahasiswa},
	})
}

func GetAllMahasiswa(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var allMahasiswa []models.Mahasiswa
	defer cancel()

	results, err := mahasiswaCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var mhs models.Mahasiswa
		if err := results.Decode(&mhs); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}
		allMahasiswa = append(allMahasiswa, mhs)
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": allMahasiswa},
	})
}

func CreateMahasiswa(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var mahasiswa models.Mahasiswa
	defer cancel()

	if err := c.Bind(&mahasiswa); err != nil {
		return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	newMahasiswa := models.Mahasiswa{
		Id:       primitive.NewObjectID(),
		Nama:     mahasiswa.Nama,
		NRP:      mahasiswa.NRP,
		Semester: mahasiswa.Semester,
		Email:    mahasiswa.Email,
		Password: mahasiswa.Password,
	}

	result, err := mahasiswaCollection.InsertOne(ctx, newMahasiswa)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusCreated, responses.DefaultResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    &echo.Map{"data": result},
	})
}

func DeleteMahasiswa(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := mahasiswaCollection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.DefaultResponse{
			Status:  http.StatusNotFound,
			Message: "Error",
			Data:    &echo.Map{"data": "Mahasiswa not found"},
		})
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": "Mahasiswa deleted"},
	})
}
