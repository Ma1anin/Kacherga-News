const FileService = require('@services/file.service.js');
const User = require("@models/User.js");

class UserService {
  async createUser(user, picture) {
    const fileName = FileService.saveFile(picture);
    return await User.create({ ...user, picture: fileName });
  }

  async getUserById(_id) {
    if (!_id) throw new Error("The ID is not specified");
    return await User.findById(_id);
  }

  async updateUser(user) {
    if (!user._id) throw new Error("The ID is not specified");
    return await User.findByIdAndUpdate(user._id, user, { new: true });
  }

  async deleteUser(_id) {
    if (!_id) throw new Error("The ID is not specified");
    return await User.findByIdAndDelete(_id);
  }
}

module.exports = new UserService();
