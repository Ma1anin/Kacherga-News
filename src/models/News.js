import mongoose from 'mongoose';

const News = new mongoose.Schema({
    title: {type: String, required: true},
    content: {type: String, required: true},
    picture: {type: String},
    creratedAt: {type: Object, required: true},
    authorID: {type: String, required: true}
});

export default mongoose.model('News', News);