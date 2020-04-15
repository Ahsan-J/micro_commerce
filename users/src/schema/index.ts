import { GraphQLObjectType, GraphQLSchema, GraphQLString, GraphQLID } from 'graphql';
import adminType, { adminQuery, addAdminMutation } from './admin'
import { login, signup } from './auth';
import { addVendor } from './vendor';

const RootQuery = new GraphQLObjectType({
    name: "RootQueryType",
    fields: {
        // admin
        admin:adminQuery,
        
        // authentiation
        login,
    }
})

const RootMutation = new GraphQLObjectType({
    name: "Mutation",
    fields: {
        // admin
        registerAdmin: addAdminMutation,
        registerVendor: addVendor,

        // authentication
        signup,
    }
})

export const schema =  new GraphQLSchema({
    query: RootQuery,
    mutation: RootMutation
})