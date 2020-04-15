import mongoose from 'mongoose';

export default mongoose.model('Company',new mongoose.Schema({
    id: String,
    name: String,
    description: String,
    companyContact: String,
    createdAt: String,
    address: String,
    contact: String,
    membershipStatus: Number,
}))