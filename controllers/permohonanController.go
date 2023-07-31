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

var permohonanCollection *mongo.Collection = configs.GetCollection(configs.DB, "permohonan")

func GetPermohonan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var permohonan models.Permohonan
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	err := permohonanCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&permohonan)

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
		Data:    &echo.Map{"data": permohonan},
	})
}

func GetAllPermohonan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	nama := c.QueryParam("nama")
	var allPermohonan []models.Permohonan
	defer cancel()

	filter := bson.M{}

	if nama != "" {
		filter["pemohon.nama"] = nama
	}

	results, err := permohonanCollection.Find(ctx, filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var permohonan models.Permohonan
		if err := results.Decode(&permohonan); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}
		allPermohonan = append(allPermohonan, permohonan)
	}
	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": allPermohonan},
	})
}

func UpdatePermohonan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var permohonan models.Permohonan
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	if err := c.Bind(&permohonan); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	if validationErr := validate.Struct(&permohonan); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": validationErr.Error()},
		})
	}

	update := bson.M{
		"tipe":   permohonan.Tipe,
		"status": permohonan.Status,
		"pemohon": bson.M{
			"nama":           permohonan.Pemohon.Nama,
			"nomor_pengenal": permohonan.Pemohon.Nomor_Pengenal,
		},
		"berkas": bson.M{
			"nama_berkas": permohonan.Berkas.NamaBerkas,
			"url_berkas":  permohonan.Berkas.URLBerkas,
		},
		"tgl_masuk": permohonan.TglMasuk,
	}

	result, err := permohonanCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	var updatedPermohonan models.Permohonan
	if result.MatchedCount == 1 {
		err := permohonanCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedPermohonan)
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
		Data:    &echo.Map{"data": updatedPermohonan},
	})
}

func CreatePermohonan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var permohonan models.Permohonan
	defer cancel()

	if err := c.Bind(&permohonan); err != nil {
		return c.JSON(http.StatusBadRequest, responses.DefaultResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	newPermohonan := models.Permohonan{
		Id:     primitive.NewObjectID(),
		Tipe:   permohonan.Tipe,
		Status: permohonan.Status,
		Pemohon: models.Pemohon{
			Nama:           permohonan.Pemohon.Nama,
			Nomor_Pengenal: permohonan.Pemohon.Nomor_Pengenal,
		},
		Berkas: models.BerkasPermohonan{
			NamaBerkas: permohonan.Berkas.NamaBerkas,
			URLBerkas:  permohonan.Berkas.URLBerkas,
		},
		TglMasuk: permohonan.TglMasuk,
	}

	result, err := permohonanCollection.InsertOne(ctx, newPermohonan)

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

func DeletePermohonan(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := permohonanCollection.DeleteOne(ctx, bson.M{"_id": objId})

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
			Data:    &echo.Map{"data": "Permohonan with specified ID not found"},
		})
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": "Permohonan with specified ID successfully deleted"},
	})
}
