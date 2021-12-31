package util

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model interface {
	Created()
	Updated()
	SetId(id interface{})
	SetLoaded(loaded bool)
}

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`

	R        Repository `bson:"-" json:"-"`
	isLoaded bool
}

func (bm *BaseModel) Created() {
	bm.CreatedAt = time.Now().UTC()
	bm.UpdatedAt = time.Now().UTC()
}

func (bm *BaseModel) Updated() {
	bm.UpdatedAt = time.Now().UTC()
}

func (bm *BaseModel) SetId(id interface{}) {
	bm.ID = id.(primitive.ObjectID)
}

func (bm *BaseModel) SetLoaded(loaded bool) {
	bm.isLoaded = loaded
}

func (bm *BaseModel) Loaded() bool {
	return bm.isLoaded
}
