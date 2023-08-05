package middlewares

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

var mahasiswaCollection *mongo.Collection = configs.GetCollection(configs.DB, "mahasiswa")
var dosenCollection *mongo.Collection = configs.GetCollection(configs.DB, "dosen")
var adminCollection *mongo.Collection = configs.GetCollection(configs.DB, "admin")
var sessionCollection *mongo.Collection = configs.GetCollection(configs.DB, "session")

func LoginHandler(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	email := c.FormValue("email")
	pass := c.FormValue("pass")
	var user models.User
	var mahasiswa models.Mahasiswa
	var dosen models.Dosen
	var admin models.Admin
	defer cancel()

	filter := bson.M{
		"email":    email,
		"password": pass,
	}

	err := mahasiswaCollection.FindOne(ctx, filter).Decode(&mahasiswa)

	if err != nil {
		err := dosenCollection.FindOne(ctx, filter).Decode(&dosen)

		if err != nil {
			err := adminCollection.FindOne(ctx, filter).Decode(&admin)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
					Status:  http.StatusInternalServerError,
					Message: "Error",
					Data:    &echo.Map{"data": err.Error()},
				})
			}

			user.Id = primitive.NewObjectID()
			user.IdUser = admin.Id.Hex()
			user.Email = admin.Email
			user.Nama = admin.Nama
			user.NomorPengenal = "#"
			user.Role = "admin"

			result, err := sessionCollection.InsertOne(ctx, user)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
					Status:  http.StatusInternalServerError,
					Message: "Error",
					Data:    &echo.Map{"data": err.Error()},
				})
			}

			return c.JSON(http.StatusOK, responses.DefaultResponse{
				Status:  http.StatusOK,
				Message: "DSN",
				Data:    &echo.Map{"data": result},
			})
		}

		user.Id = primitive.NewObjectID()
		user.IdUser = dosen.Id.Hex()
		user.Email = dosen.Email
		user.Nama = dosen.Nama
		user.NomorPengenal = dosen.NIP
		user.Role = "dosen"
		user.Login = time.Now().Format("2006-01-02 15:04:05")

		result, err := sessionCollection.InsertOne(ctx, user)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error",
				Data:    &echo.Map{"data": err.Error()},
			})
		}

		return c.JSON(http.StatusOK, responses.DefaultResponse{
			Status:  http.StatusOK,
			Message: "DSN",
			Data:    &echo.Map{"data": result},
		})
	}

	user.Id = primitive.NewObjectID()
	user.IdUser = mahasiswa.Id.Hex()
	user.Nama = mahasiswa.Nama
	user.Email = mahasiswa.Email
	user.NomorPengenal = mahasiswa.NRP
	user.Role = "mahasiswa"
	user.Login = time.Now().Format("2006-01-02 15:04:05")

	result, err := sessionCollection.InsertOne(ctx, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.DefaultResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error",
			Data:    &echo.Map{"data": err.Error()},
		})
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "MHS",
		Data:    &echo.Map{"data": result},
	})
}

func LogoutHandler(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.QueryParam("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := dosenCollection.DeleteOne(ctx, bson.M{"_id": objId})

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
			Data:    &echo.Map{"data": "Session ID not found"},
		})
	}

	return c.JSON(http.StatusOK, responses.DefaultResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    &echo.Map{"data": "Session deleted"},
	})
}

func GetSession(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	id := c.QueryParam("id")
	defer cancel()

	err := sessionCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

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
		Data:    &echo.Map{"data": user},
	})
}
