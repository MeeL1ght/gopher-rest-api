package controllers

import (
	"context"
	"net/http"
	"time"

	db "github.com/MeeL1ght/gopher-rest-api/database"
	model "github.com/MeeL1ght/gopher-rest-api/models"
	utils "github.com/MeeL1ght/gopher-rest-api/utilities/pointers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	validate                            = validator.New()
	gophersCollection *mongo.Collection = db.MongoCollection(
		db.NewMongoClient(),
		"gophers",
	)
)

// Create a new Gopher c:
func CreateGopher(c *fiber.Ctx) error {
	var gopher model.Gopher

	// Validation of the request body
	if err := c.BodyParser(&gopher); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			bson.M{
				"statusCode": http.StatusBadRequest,
				"message":    "error",
				"data":       err.Error(),
			},
		)
	}

	// Structure validation
	if err := validate.Struct(&gopher); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			bson.M{
				"statusCode": http.StatusBadRequest,
				"message":    "error",
				"data":       err.Error(),
			},
		)
	}

	// Checking and setting values x)
	if gopher.Status == nil {
		gopher.Status = utils.GetNewBoolPointerValue(true)
	}

	gopher.CreatedAt = time.Now()

	// Insert
	filter, err := gophersCollection.InsertOne(
		context.TODO(),
		gopher,
	)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			bson.M{
				"statusCode": http.StatusInternalServerError,
				"message":    "error",
				"data":       err.Error(),
			},
		)
	}

	// Encode
	bson.Marshal(filter)

	return c.Status(http.StatusCreated).JSON(
		bson.M{
			"statusCode": http.StatusCreated,
			"message":    "success",
			"data": bson.M{
				"name":      gopher.Name,
				"color":     gopher.Color,
				"status":    gopher.Status,
				"createdAt": gopher.CreatedAt,
			},
		},
	)
}

// Read a Gopher
func ReadGopher(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	gopherId := c.Params("id")
	var gopher model.Gopher

	defer cancel()

	var filter bson.M
	count, _ := gophersCollection.CountDocuments(context.TODO(), filter)

	if count == 0 {
		return c.Status(http.StatusNotFound).JSON(
			bson.M{
				"statusCode": http.StatusNotFound,
				"message":    "error",
				"data":       "There are no documents in the collection :(",
			},
		)
	} else {
		objId, idError := primitive.ObjectIDFromHex(gopherId)

		if idError != nil {
			c.Status(http.StatusNotFound).JSON(
				bson.M{
					"statusCode": http.StatusNotFound,
					"message":    "error",
					"data":       "id not found :(",
				},
			)
		}

		err := gophersCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&gopher)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				bson.M{
					"statusCode": http.StatusInternalServerError,
					"message":    "error",
					"data":       err.Error(),
				},
			)
		}

		return c.Status(http.StatusOK).JSON(
			bson.M{
				"statusCode": http.StatusOK,
				"message":    "success",
				"data": bson.M{
					"name":      gopher.Name,
					"color":     gopher.Color,
					"status":    gopher.Status,
					"createdAt": gopher.CreatedAt,
					"updatedAt": gopher.UpdatedAt,
				},
			},
		)
	}
}

// Read Gophers
func ReadGophers(c *fiber.Ctx) error {
	var gophers bson.A

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{}
	cur, err := gophersCollection.Find(ctx, filter)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			bson.M{
				"statusCode": http.StatusInternalServerError,
				"message":    "error",
				"data":       err.Error(),
			},
		)
	}

	for cur.Next(ctx) {
		var gopher model.Gopher

		if err = cur.Decode(&gopher); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				bson.M{
					"statusCode": http.StatusInternalServerError,
					"message":    "error",
					"data":       err.Error(),
				},
			)
		}

		gophers = append(
			gophers,
			bson.M{
				"name":      gopher.Name,
				"color":     gopher.Color,
				"status":    gopher.Status,
				"createdAt": gopher.CreatedAt,
				"updatedAt": gopher.UpdatedAt,
			},
		)
	}

	return c.Status(http.StatusOK).JSON(
		bson.M{
			"statusCode": http.StatusOK,
			"message":    "success",
			"data":       gophers,
		},
	)
}

// Update a Gopher
func UpdateGopher(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var filter bson.M
	count, _ := gophersCollection.CountDocuments(context.TODO(), filter)

	if count == 0 {
		return c.Status(http.StatusNotFound).JSON(
			bson.M{
				"statusCode": http.StatusNotFound,
				"message":    "error",
				"data":       "There are no documents in the collection :(",
			},
		)
	} else {
		var currentGopher model.Gopher
		gopherId := c.Params("id")
		objId, _ := primitive.ObjectIDFromHex(gopherId)

		// Find and get
		err := gophersCollection.FindOne(
			ctx,
			bson.M{"_id": objId},
		).Decode(&currentGopher)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				bson.M{
					"statusCode": http.StatusInternalServerError,
					"message":    "error",
					"data":       err.Error(),
				},
			)
		}

		var newGopher model.Gopher

		// Validation of the request body
		if err := c.BodyParser(&newGopher); err != nil {
			return c.Status(http.StatusBadRequest).JSON(
				bson.M{
					"statusCode": http.StatusBadRequest,
					"message":    "error",
					"data":       err.Error(),
				},
			)
		}

		// Structure validation
		if err := validate.Struct(&newGopher); err != nil {
			return c.Status(http.StatusBadRequest).JSON(
				bson.M{
					"statusCode": http.StatusBadRequest,
					"message":    "error",
					"data":       err.Error(),
				},
			)
		}

		// Checking and setting values x)

		// Name
		if newGopher.Name == nil {
			newGopher.Name = utils.GetNewStringPointerValue(*currentGopher.Name)
		}

		// Color
		if newGopher.Color == nil {
			newGopher.Color = utils.GetNewStringPointerValue(*currentGopher.Color)
		}

		// Status
		if newGopher.Status == nil {
			newGopher.Status = utils.GetNewBoolPointerValue(*currentGopher.Status)
		}

		// Filter
		filter = bson.M{"_id": objId}

		update := bson.M{
			"$set": bson.M{
				"name":      newGopher.Name,
				"color":     newGopher.Color,
				"status":    newGopher.Status,
				"updatedAt": time.Now(),
			},
		}

		result, newErr := gophersCollection.UpdateOne(ctx, filter, update)

		if newErr != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				bson.M{
					"statusCode": http.StatusInternalServerError,
					"message":    "error",
					"data":       newErr.Error(),
				},
			)
		}

		var updatedGopher model.Gopher

		if result.MatchedCount == 1 {
			err := gophersCollection.FindOne(ctx, filter).Decode(&updatedGopher)

			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(
					bson.M{
						"statusCode": http.StatusInternalServerError,
						"message":    "error",
						"data":       err.Error(),
					},
				)
			}
		}

		return c.Status(http.StatusOK).JSON(
			bson.M{
				"statusCode": http.StatusOK,
				"message":    "success",
				"data":       updatedGopher,
			},
		)
	}
}

// Delete Gopher :x
func DeleteGopher(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	gopherId := c.Params("id")

	defer cancel()

	var filter bson.M
	count, _ := gophersCollection.CountDocuments(context.TODO(), filter)

	if count == 0 {
		return c.Status(http.StatusNotFound).JSON(
			bson.M{
				"statusCode": http.StatusNotFound,
				"message":    "error",
				"data":       "There are no documents in the collection :(",
			},
		)
	} else {
		objId, _ := primitive.ObjectIDFromHex(gopherId)
		result, err := gophersCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				bson.M{
					"statusCode": http.StatusInternalServerError,
					"message":    "error",
					"data":       err.Error(),
				},
			)
		}

		if result.DeletedCount < 1 {
			return c.Status(http.StatusNotFound).JSON(
				bson.M{
					"statusCode": http.StatusNotFound,
					"message":    "error",
					"data":       "id not found :(",
				},
			)
		}

		return c.Status(http.StatusOK).JSON(
			bson.M{
				"statusCode": http.StatusOK,
				"message":    "success",
				"data":       "Gopher successfully removed!",
			},
		)
	}
}
