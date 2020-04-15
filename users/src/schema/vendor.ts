import { GraphQLObjectType, GraphQLID, GraphQLString, GraphQLInt } from 'graphql';

export default new GraphQLObjectType({
    name:"Vendor",
    fields: function() {
        return {
            id: {type: GraphQLID},
            name: {type: GraphQLString},
            email: {type: GraphQLString},
            status: {type: GraphQLInt},
            createdAt: {type: GraphQLString},
            password: {type: GraphQLString},
        }
    },
    
})