const User = require("../models/User");
const argon2 = require("argon2");
const { validationResult } = require("express-validator");

class UserController {
  async login(req, res) {
    const { login, password } = req.body;
    try {
      if (req.session.userId) {
        return res.status(400).json({
          success: false,
          message: "You already authorized",
        });
      }

      const user = await User.findOne({ login: login });

      if (!user) {
        return res.status(401).json({
          success: false,
          message: "User not found",
        });
      }

      if (await argon2.verify(user.password, password)) {
        req.session.userId = user.id;
        return res.status(200).json({
          success: true,
          data: {
            login: user.login,
            fullName: user.fullName,
            picture: user.picture,
          },
          message: "Successful account login",
        });
      } else {
        return res.status(401).json({
          success: false,
          message: "Incorrect password",
        });
      }
    } catch (err) {
      console.log(err);
      res.status(500).json({
        success: false,
        message: "Server Error Please reload page",
      });
    }
  }

  async logout(req, res) {
    if (req.session.userId) {
      req.session.destroy();
      res.status(200).json({
        success: true,
        message: "Successful account logout",
      });
    } else {
      res.status(200).json({
        success: false,
        message: "Unauthorized",
      });
    }
  }

  async register(req, res) {
    const { fullName, login, password } = req.body;
    try {
      const errors = validationResult(req);

      if (!errors.isEmpty()) {
        return res.status(401).json({
          success: false,
          errors: errors.array(),
        });
      }

      const user = await User.findOne({ login: login });

      if (user) {
        return res.status(401).json({
          success: false,
          message: "The user already exists",
        });
      }

      const hashedPassword = await argon2.hash(password);
      const newUser = {
        login: login,
        password: hashedPassword,
        fullName: fullName,
      };

      await User.create(newUser);

      return res.status(200).json({
        success: true,
        message: "Successful registration",
      });
    } catch (err) {
      res.status(500).json({
        success: false,
        message: "Server Error Please reload page",
      });
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
