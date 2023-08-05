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

var pengumpulanCollection *mongo.Collection = configs.GetCollection(configs.ConnectDB(), "pengumpulan")

func GetAllPengumpulan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	nama := c.QueryParam("nama")
	jenis := c.QueryParam("jenis")
	var allPengumpulan []models.Pengumpulan
	defer cancel()

	filter := bson.M{}

	if nama != "" {
		filter["pemohon.nama"] = nama
	}

	if jenis != "" {
		filter["jenis"] = jenis
	}

	results, err := pengumpulanCollection.Find(ctx, filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var pengumpulan models.Pengumpulan
		if err := results.Decode(&pengumpulan); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}
		allPengumpulan = append(allPengumpulan, pengumpulan)
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": allPengumpulan},
	})
}

func GetPengumpulan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var pengumpulan models.Pengumpulan
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	err := pengumpulanCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&pengumpulan)

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
		Data:    &echo.Map{"data": pengumpulan},
	})
}

func CreatePengumpulan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var pengumpulan models.Pengumpulan
	defer cancel()

	if err := c.Bind(&pengumpulan); err != nil {
		return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	newPengumpulan := models.Pengumpulan{
		Id:            primitive.NewObjectID(),
		Nama:          pengumpulan.Nama,
		NomorPengenal: pengumpulan.NomorPengenal,
		Jenis:         pengumpulan.Jenis,
		Berkas: models.BerkasPengumpulan{
			NamaBerkas: pengumpulan.Berkas.NamaBerkas,
			URLBerkas:  pengumpulan.Berkas.URLBerkas,
		},
		Tgl: time.Now().Format("2006-01-02"),
	}

	result, err := pengumpulanCollection.InsertOne(ctx, newPengumpulan)

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

func UpdatePengumpulan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var pengumpulan models.Pengumpulan
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	if err := c.Bind(&pengumpulan); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	if validationErr := validate.Struct(&pengumpulan); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": validationErr.Error()},
		})
	}

	update := bson.M{
		"nama":           pengumpulan.Nama,
		"nomor_pengenal": pengumpulan.NomorPengenal,
		"jenis":          pengumpulan.Jenis,
		"berkas": bson.M{
			"nama_berkas": pengumpulan.Berkas.NamaBerkas,
			"url_berkas":  pengumpulan.Berkas.URLBerkas,
		},
		"tgl": pengumpulan.Tgl,
	}

	result, err := pengumpulanCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	var updatedPengumpulan models.Pengumpulan
	if result.MatchedCount == 1 {
		err := pengumpulanCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedPengumpulan)
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
		Data:    &echo.Map{"data": updatedPengumpulan},
	})
}

func DeletePengumpulan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := pengumpulanCollection.DeleteOne(ctx, bson.M{"_id": objId})

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
			Data:    &echo.Map{"data": "Pengumpulan not found"},
		})
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": "Pengumpulan deleted"},
	})
}
