import { GraphQLObjectType, GraphQLID, GraphQLString, GraphQLInt } from 'graphql';

export default new GraphQLObjectType({
    name:"User",
    fields: function() {
        return {
            id: {type: GraphQLID},
            name: {type: GraphQLString},
            email: {type: GraphQLString},
            contact: {type: GraphQLString},
            cartId: {type: GraphQLID},
            status: {type: GraphQLInt},
            adddress: {type: GraphQLString},
            dateOfBirth: {type: GraphQLString}, // To filter out the product based on age group
            createdAt: {type: GraphQLString},
            password: {type: GraphQLString},
            error: {type: GraphQLString}
        }
    },
    
})