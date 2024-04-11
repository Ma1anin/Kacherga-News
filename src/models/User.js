import mongoose from 'mongoose';

const User = new mongoose.Schema({
    login: {type: String, required: true},
    fullName: {type: String, required: true},
    role: {type: String, required: true},
    avatar: {type: String}
});

export default mongoose.model('User', User);