import { GraphQLObjectType, GraphQLString, GraphQLInt } from "graphql";

const companyType = new GraphQLObjectType({
    name: "Company",
    
    fields: function () {
        return {
            id: { type: GraphQLString },
            name: { type: GraphQLString },
            description: { type: GraphQLString },
            membershipStatus: { type: GraphQLInt },
            createdAt: { type: GraphQLString },
            companyURL: { type: GraphQLString },
            companyContact: { type: GraphQLInt },
            address: { type: GraphQLString },
            error: { type: GraphQLString}
        }
    }
})

export default companyType;