import { GraphQLObjectType, GraphQLID, GraphQLString, GraphQLInt } from 'graphql';
import companyType from './company';
import { validate } from '../helper/utility';
import Company from '../db/model/Company';

const vendorType = new GraphQLObjectType({
    name:"Vendor",
    fields: function() {
        return {
            id: {type: GraphQLID},
            name: {type: GraphQLString},
            email: {type: GraphQLString},
            status: {type: GraphQLInt},
            createdAt: {type: GraphQLString},
            password: {type: GraphQLString},
            company: {
                type: companyType,
                resolve: async (parent, args) => {
                    if(validate("" + parent.companyId)) {
                        try {
                            return await Company.find({ id : parent.companyId})
                        } catch(e) {
                            return { error : "No Company registered"}
                        }
                    }
                    return {
                        error: "companyId missing"
                    }
                }
            }
        }
    },
    
})

export const addVendor = {
    type: vendorType,
    args: {
        name: {type: GraphQLString},
        email: {type: GraphQLString},
        password: {type: GraphQLString}
    }

}

export default vendorType