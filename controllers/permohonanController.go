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

var pmhCollection *mongo.Collection = configs.GetCollection(configs.DB, "permohonan")

func GetOnePmh(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var pmh models.Permohonan
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	err := pmhCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&pmh)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.PermohonanResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusOK, responses.PermohonanResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": pmh},
	})
}

func GetAllPmh(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var allPmh []models.Permohonan
	defer cancel()

	results, err := pmhCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.PermohonanResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var pmh models.Permohonan
		if err := results.Decode(&pmh); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.PermohonanResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}
		allPmh = append(allPmh, pmh)
	}

	return c.JSON(http.StatusOK, responses.PermohonanResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": allPmh},
	})
}

func EditAPmh(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var pmh models.Permohonan
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	if err := c.Bind(&pmh); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.PermohonanResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	if validationErr := validate.Struct(&pmh); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.PermohonanResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": validationErr.Error()},
		})
	}

	update := bson.M{
		"jenis":  pmh.Jenis,
		"status": pmh.Status,
		"pemohon": bson.M{
			"nama": pmh.Pemohon.Nama,
			"nomor_pengenal":  pmh.Pemohon.Nomor_Pengenal,
		},
		"tujuan": pmh.Tujuan,
		"berkas": bson.M{
			"nama_berkas": pmh.Berkas.NamaBerkas,
			"url_berkas":  pmh.Berkas.URLBerkas,
		},
		"tgl_masuk": pmh.TglMasuk,
	}

	result, err := pmhCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.PermohonanResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	var updatedPmh models.Permohonan
	if result.MatchedCount == 1 {
		err := pmhCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedPmh)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.PermohonanResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}
	}

	return c.JSON(http.StatusOK, responses.PermohonanResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": updatedPmh},
	})
}

func CreatePmh(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var pmh models.Permohonan
	defer cancel()

	if err := c.Bind(&pmh); err != nil {
		return c.JSON(http.StatusBadRequest, responses.PermohonanResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	newPmh := models.Permohonan{
		Id:     primitive.NewObjectID(),
		Jenis:  pmh.Jenis,
		Status: pmh.Status,
		Pemohon: models.Pemohon{
			Nama:           pmh.Pemohon.Nama,
			Nomor_Pengenal: pmh.Pemohon.Nomor_Pengenal,
		},
		Tujuan: pmh.Tujuan,
		Berkas: models.Berkas{
			NamaBerkas: pmh.Berkas.NamaBerkas,
			URLBerkas:  pmh.Berkas.URLBerkas,
		},
		TglMasuk: pmh.TglMasuk,
	}

	result, err := pmhCollection.InsertOne(ctx, newPmh)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.PermohonanResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusCreated, responses.PermohonanResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    &echo.Map{"data": result},
	})
}

func DeletePmh(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := pmhCollection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.PermohonanResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.PermohonanResponse{
			Status:  http.StatusNotFound,
			Message: "Error",
			Data:    &echo.Map{"data": "Permohonan with specified ID not found"},
		})
	}

	return c.JSON(http.StatusOK, responses.PermohonanResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": "Permohonan with specified ID successfully deleted"},
	})
}
