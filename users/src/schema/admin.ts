import { GraphQLObjectType, GraphQLID, GraphQLString, GraphQLInt } from 'graphql';
import User from '../db/model/User';
import shortId from 'shortid';
import moment from 'moment';
import _ from 'lodash';
import { validate } from '../helper/utility';
import { AdminUser } from '../types/admin';
import { AccountType, AccountStatus } from '../types/user';

const adminType = new GraphQLObjectType({ // GraphQL Node definition
    name: "Admin",
    fields: function () {
        return {
            id: { type: GraphQLID },
            name: { type: GraphQLString },
            email: { type: GraphQLString },
            status: { type: GraphQLInt },
            createdAt: { type: GraphQLString },
            password: { type: GraphQLString },
            authLevel: {type: GraphQLInt},
            error: {type: GraphQLString}
        }
    },
})

export const adminQuery = {
    type: adminType,
    args: { id: { type: GraphQLID } },
    resolve: (parent: any, args: any) => {
        return {
            id: 1,
            name: "Ahsan",
            email: "email",
            status: 1,
            createdAt: "{type: GraphQLString}",
            password: "pass"
        }
    }
}

export const addAdminMutation = {
    type: adminType,
    args: {
        name: {type: GraphQLString},
        email: {type: GraphQLString},
        password: {type: GraphQLString},
        createdBy: {type: GraphQLID},
        authLevel: {type: GraphQLInt}
    },
    resolve: async (parent: any, args: any) => {
        if(!validate(args.name)) {
            console.log("Name is undefined")
            return {
                error: "Invalid or not defined name"
            }
        }
        if(!validate(args.email)) {
            console.log("Email is undefined");
            return {
                error: "Invalid or not defined email"
            }
        }
        if(!validate(args.password)) {
            console.log("Password is undefined")
            return {
                error: "Invalid or not defined password"
            }
        }
        if(!validate(args.authLevel + "")) {
            console.log("authLevel is undefined")
            return {
                error: "Invalid or not defined authLevel"
            }
        }
        if(!validate(args.createdBy + "")) {
            console.log("Created by is undefined")
            return {
                error: "Created By is missing"
            }
        }
        try {
            if(_.size(await User.find({ id: args.createdBy})) <= 0) {
                return {
                    error: "Invalid user is attempting to add Admin"
                }
            }
    
            const newAdmin = new User({
                id: shortId.generate(),
                name: args.name,
                email: args.email,
                password: args.password,
                status: AccountStatus.active,
                userType: AccountType.admin,
                createdAt: moment().toISOString(),
                authLevel: args.authLevel || 1
            } as AdminUser)
    
            return newAdmin.save();
        } catch(e) {
            return {
                error: "Error while saving a user"
            }
        }
        
    }
}

export default adminType;