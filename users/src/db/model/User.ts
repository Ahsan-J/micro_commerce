import mongoose from 'mongoose';

export default mongoose.model('User',new mongoose.Schema({
    id: String,
    name: String,
    email: String,
    status: Number,
    createdAt: String,
    password: String,
    userType: Number,
    cartId: String,
    address: String,
    contact: String,
    dateOfBirth: String,
    authLevel: Number,
    createdBy: String,
}))