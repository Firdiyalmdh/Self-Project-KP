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

var pengumumanCollection *mongo.Collection = configs.GetCollection(configs.DB, "pengumuman")

// var validate = validator.New()

func GetPengumuman(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var pengumuman models.Pengumuman
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	err := pengumumanCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&pengumuman)

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
		Data:    &echo.Map{"data": pengumuman},
	})
}

func GetAllPengumuman(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var allPengumuman []models.Pengumuman
	defer cancel()

	results, err := pengumumanCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var pengumuman models.Pengumuman
		if err := results.Decode(&pengumuman); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}
		allPengumuman = append(allPengumuman, pengumuman)
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": allPengumuman},
	})
}

func UpdatePengumuman(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var pengumuman models.Pengumuman
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	if err := c.Bind(&pengumuman); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	if validationErr := validate.Struct(&pengumuman); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": validationErr.Error()},
		})
	}

	update := bson.M{
		"jenis": pengumuman.Jenis,
		"announcer": bson.M{
			"nama":           pengumuman.Announcer.Nama,
			"nomor_pengenal": pengumuman.Announcer.Nomor_Pengenal,
		},
		"content": bson.M{
			"tgl":  pengumuman.Content.Tgl,
			"data": pengumuman.Content.Data,
		},
	}

	result, err := pengumumanCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	var updatedPengumuman models.Pengumuman
	if result.MatchedCount == 1 {
		err := pengumumanCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedPengumuman)
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
		Data:    &echo.Map{"data": updatedPengumuman},
	})
}

func CreatePengumuman(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var pengumuman models.Pengumuman
	defer cancel()

	if err := c.Bind(&pengumuman); err != nil {
		return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	newPengumuman := models.Pengumuman{
		Id:    primitive.NewObjectID(),
		Jenis: pengumuman.Jenis,
		Announcer: models.Announcer{
			Nama:           pengumuman.Announcer.Nama,
			Nomor_Pengenal: pengumuman.Announcer.Nomor_Pengenal,
		},
		Content: models.Content{
			Tgl:  pengumuman.Content.Tgl,
			Data: pengumuman.Content.Data,
		},
	}

	result, err := pengumumanCollection.InsertOne(ctx, newPengumuman)

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

func DeletePengumuman(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := pengumumanCollection.DeleteOne(ctx, bson.M{"_id": objId})

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
			Data:    &echo.Map{"data": "Pengumuman not found"},
		})
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": "Pengumuman deleted"},
	})
}
