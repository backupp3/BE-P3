package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username   string             `bson:"username" json:"username"`
	Password   string             `bson:"password" json:"password"`
}

type Reservasi struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama 		  string             `bson:"nama,omitempty" json:"nama,omitempty"`
	No_telp		  string             `bson:"no_telp,omitempty" json:"no_telp,omitempty"`
	TTL     	  string             `bson:"ttl,omitempty" json:"ttl,omitempty"`
	Status	      string             `bson:"status,omitempty" json:"status,omitempty"`
	Keluhan	      string             `bson:"keluhan,omitempty" json:"keluhan,omitempty"`
}

type Response struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type ReservasiResponse struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message,omitempty" bson:"message,omitempty"`
	Data    []Reservasi `json:"data" bson:"data"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Data    []User `bson:"data,omitempty" json:"data,omitempty"`
}