// Code generated by github.com/trevor-scheer/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/trevor-scheer/gqlgen/_examples/scalars/external"
)

type Address struct {
	ID       external.ObjectID `json:"id"`
	Location *Point            `json:"location,omitempty"`
}

type Query struct {
}
