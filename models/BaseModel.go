package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

func (bm *BaseModel) Create() {
	bm.CreatedAt = time.Now().UTC()
}

func (bm *BaseModel) Update() {
	bm.UpdatedAt = time.Now().UTC()
}
