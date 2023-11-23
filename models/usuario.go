package models

import (
	"github.com/graphql-go/graphql"
)

type Usuario struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "usuario",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
