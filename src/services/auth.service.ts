const argon2 = require("argon2");
const jwt = require("jsonwebtoken");
const User = require("../models/User.js");

class AuthService {
  public async register(
    login: string,
    password: string,
    name: string
  ): Promise<any> {
    const passwordHashed = await argon2.hash(password);

    const userRecord = await User.create({
      password: passwordHashed,
      login,
      name,
      role: "user",
      picture: "null.jpg",
    });

    return {
      user: {
        login: userRecord.login,
        name: userRecord.name,
      },
    };
  }

  public async login(login: string, password: string): Promise<any> {
    const userRecord = await User.findOne(login);

    if (!userRecord) {
      throw new Error("User not found");
    } else {
      const correctPassword = await argon2.verify(
        userRecord.password,
        password
      );
      if (!correctPassword) throw new Error("Incorrect password");
    }

    return {
      user: {
        login: userRecord.login,
        name: userRecord.name,
      },
      token: this.generateToken(userRecord),
    };
  }

  private generateToken(user) {
    const data = {
      _id: user._id,
      login: user.login,
      name: user.name,
    };

    const signature = "Joylov";
    const expiration = "6h";

    return jwt.sign({ data }, signature, { expiresIn: expiration });
  }
}

module.exports = new AuthService();
