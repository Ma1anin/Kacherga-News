const mongoose = require('mongoose');

const Event = new mongoose.Schema({
    title: {type: String, required: true},
    content: {type: String, required: true},
    creratedAt: {type: Object, required: true},
    authorID: {type: String, required: true}
});

module.exports = mongoose.model('Event', Event);