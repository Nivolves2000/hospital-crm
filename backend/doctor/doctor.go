package doctor

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"../db"
	"context"
)

type Doctor struct {
	DoctorID		primitive.ObjectID `bson:"_id" json:"doctorId,omitempty"`
	Role				string `json:"role,omitempty"`
	FirebaseUID	string `json:"firebaseUID,omitempty"`     
}

func AddDoctor(c echo.Context) error {
	client, ctx := db.GetDb()
	doctor := Doctor{}

	err := c.Bind(&doctor)
  if err != nil {
    log.Printf("Failed POST doctor request: %s\n", err)
    return echo.NewHTTPError(http.StatusInternalServerError)
	}

	collection := client.Database("hospital-crm").Collection("doctors")
	insertResult, err := collection.InsertOne(context.Background(),  map[string]interface{}{
		"role": doctor.Role,
		"firebaseUID": doctor.FirebaseUID,
	})
	if err != nil {
    log.Printf("Failed POST patient request: %s\n", err)
    return echo.NewHTTPError(http.StatusInternalServerError)
	}

	defer client.Disconnect(ctx)

	return c.JSON(http.StatusOK, insertResult)
}

func GetDoctor(c echo.Context) error {
	firebaseId := c.Param("id")
	client, ctx := db.GetDb()
	collection := client.Database("hospital-crm").Collection("doctors")
	var doctor Doctor

	err := collection.FindOne(context.Background(), bson.M{"firebaseUID": firebaseId}).Decode(&doctor);
	if err != nil {
    log.Printf("Failed GET patients request: %s\n", err)
    return echo.NewHTTPError(http.StatusInternalServerError)
	}

	defer client.Disconnect(ctx)

	return c.String(http.StatusOK, doctor.DoctorID.Hex())
}
