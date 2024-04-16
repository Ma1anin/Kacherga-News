const mongoose = require('mongoose');

const News = new mongoose.Schema({
    title: {type: String, required: true},
    content: {type: String, required: true},
    picture: {type: String},
    createdAt: {type: Object, required: true},
    authorID: {type: String, required: true}
});

module.exports = mongoose.model('News', News);