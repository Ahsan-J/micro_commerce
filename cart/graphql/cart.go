package graphql

import (
	"github.com/graphql-go/graphql"
)

// CartType is the graphQL type for Cart
var CartType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Cart",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"done": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

// GetCart is a query to return after resolving the arguments
var GetCart = &graphql.Field{
	Type: CartType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.ID,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}

// AddCart is a query to return after resolving the arguments
var AddCart = &graphql.Field{
	Type: CartType,
	Args: graphql.FieldConfigArgument{},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}
