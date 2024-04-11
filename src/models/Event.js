import mongoose from 'mongoose';

const Event = new mongoose.Schema({
    title: {type: String, required: true},
    content: {type: String, required: true},
    creratedAt: {type: Object, required: true},
    authorID: {type: String, required: true}
});

export default mongoose.model('Event', Event);