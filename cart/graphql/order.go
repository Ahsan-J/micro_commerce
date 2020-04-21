package graphql

import (
	"github.com/graphql-go/graphql"
)

// OrderType is the graphQL type for Order
var OrderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Cart",
	Fields: graphql.Fields{
		"cartId": &graphql.Field{
			Type: graphql.ID,
		},
		"productId": &graphql.Field{
			Type: graphql.ID,
		},
		"itemAdded": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

/************************Query*****************************/

// GetProductIDs is a query to return product ids
var GetProductIDs = &graphql.Field{
	Type: CartType,
	Args: graphql.FieldConfigArgument{
		"cartId": &graphql.ArgumentConfig{
			Type: graphql.ID,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}

/************************Mutation***************************/

// AddProductToCart is a mutation to add an order
var AddProductToCart = &graphql.Field{
	Type: CartType,
	Args: graphql.FieldConfigArgument{
		"cartId": &graphql.ArgumentConfig{
			Type: graphql.ID,
		},

		"productId": &graphql.ArgumentConfig{
			Type: graphql.ID,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}
