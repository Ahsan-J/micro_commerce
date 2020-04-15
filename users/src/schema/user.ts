import { GraphQLObjectType, GraphQLString, GraphQLInt } from "graphql";

export const userType = new GraphQLObjectType({
    name: "User",
    fields: function () {
        return {
            id: { type: GraphQLString },
            name: { type: GraphQLString },
            email: { type: GraphQLString },
            status: { type: GraphQLInt },
            createdAt: { type: GraphQLString },
            password: { type: GraphQLString },
            userType: { type: GraphQLInt },
            cartId: { type: GraphQLString },
            address: { type: GraphQLString },
            contact: { type: GraphQLString },
            dateOfBirth: { type: GraphQLString },
            authLevel: { type: GraphQLInt },
            error: { type: GraphQLString}
        }
    }
})