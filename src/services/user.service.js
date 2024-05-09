const FileService = require("../services/file.service.js");
const User = require("../models/User.js");
const argon2 = require("argon2");

class UserService {
  async createUser(user) {
    const passwordHashed = await argon2.hash(user.password);
    return await User.create({ ...user, password: passwordHashed, picture: "null.jpg" });
  }

  async getUserById(_id) {
    if (!_id) throw new Error("The ID is not specified");
    return await User.findById(_id);
  }

  async updateUser(user) {
    if (!user._id) throw new Error("The ID is not specified");
    // if (picture) {fileName = FileService.saveFile(picture);}
    return await User.findByIdAndUpdate(user._id, user, { new: true });
  }

  async deleteUser(_id) {
    if (!_id) throw new Error("The ID is not specified");
    return await User.findByIdAndDelete(_id);
  }
}

module.exports = new UserService();
