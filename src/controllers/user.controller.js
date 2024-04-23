const UserService = require("../services/user.service.js");

class UserController {
  async createUser(req, res) {
    try {
      const createdUser = await UserService.createUser(req.body);
      return res.json(createdUser);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async getUserById(req, res) {
    try {
      const user = await UserService.getUserById(req.params.id);
      return res.json(user);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async updateUser(req, res) {
    try {
      const updatedUser = await UserService.updateUser(req.body);
      return res.json(updatedUser);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async deleteUser(req, res) {
    try {
      const deletedUser = await UserService.deleteUser(req.params.id);
      return res.json(deletedUser);
    } catch (err) {
      res.status(500).json(err);
    }
  }
}

module.exports = new UserController();
