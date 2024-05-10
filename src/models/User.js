const mongoose = require("mongoose");

const User = new mongoose.Schema({
  login: { type: String, required: true },
  fullName: { type: String, required: true },
  password: { type: String, required: true },
  role: { type: String, required: true, default: "user" },
  picture: { type: String, default: "null.jpg" },
});

module.exports = mongoose.model("User", User);
