const User = require("../models/User.js");

class UserService {
  async createUser(user, picture) {
    return await User.create(user);
  }

  async getUserById(_id) {
    if (!_id) throw new Error('The ID is not specified');
    return await User.findById(_id);
  }

  async updateUser(user) {
    if (!user._id) throw new Error('The ID is not specified');
    return await User.findByIdAndUpdate(user._id, user);
  }

  async deleteUser(user) {
    if (!user._id) throw new Error('The ID is not specified');
    return await User.findByIdAndDelete(user._id, user);
  }
}

module.exports = new UserService();