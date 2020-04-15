import { userType } from "./user"
import { GraphQLString } from "graphql"
import { validate } from "../helper/utility"
import User from "../db/model/User"

// Queries
export const login ={
    type: userType,
    args: { 
        email: { type: GraphQLString },
        password: {type: GraphQLString} 
    },
    resolve: async (parent: any, args: any) => {
        if(!validate(args.email)) {
            return {
                error: "Email is invalid"
            }
        }

        if(!validate(args.password)) {
            return {
                error: "Password is invalid"
            }
        }
        
        try {
            return await User.findOne({ email: args.email, password: args.password })
        }
        catch(e) {
            return {
                error:"Error while attempting to login"
            }
        }
    }
}

// Mutations
export const signup ={
    type: userType,
    args: { 
        name: { type: GraphQLString},
        email: { type: GraphQLString },
        password: {type: GraphQLString},
        address: { type: GraphQLString },
        contact: { type: GraphQLString },
        dateOfBirth: { type: GraphQLString },
    },
    resolve: async (parent: any, args: any) => {
        if(!validate(args.email)) {
            return { error: "Email is invalid" }
        }

        if(!validate(args.password)) {
            return { error: "Password is invalid" }
        }

        if(!validate(args.name)) {
            return { error: "Name is invalid" }
        }

        try {
            return await User.findOne({ email: args.email, password: args.password })
        }
        catch(e) {
            return { error:"Error while attempting to login" }
        }
    }
}