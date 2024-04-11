import User from "../models/User.js";

class UserService {
  async createUser(user, picture) {
    return await User.create(user);
  }

  async getUserById(id) {
    return await User.findById(id);
  }

  async updateUser(id, user) {
    return await User.findByIdAndUpdate(id, user);
  }

  async updateUserRole(id, role) {
    return await User.findByIdAndUpdate(id, role);
  }

  async deleteUser(id, user) {
    return await User.findByIdAndDelete(id, user);
  }
}

export default new UserService();