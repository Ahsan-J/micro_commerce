import mongoose from 'mongoose';

export default mongoose.model('Company',new mongoose.Schema({
    id: String,
    name: String,
    companyURL: String,
    description: String,
    companyContact: String,
    createdAt: String,
    address: String,
    membershipStatus: Number,
}))