package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Represents a Gopher =)
type Gopher struct {
	Id        primitive.ObjectID `bson:"-"`
	Name      *string            `bson:"name"`
	Color     *string            `bson:"color"`
	Status    *bool              `bson:"status"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty"`
}
