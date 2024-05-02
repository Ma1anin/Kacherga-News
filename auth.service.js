const argon2 = require('argon2');
const jwt = require('jsonwebtoken');
const User = require("../models/User.js");

class AuthService {
    async register(login, password, name) {
        const passwordHashed = await argon2.hash(password);

        const userRecord = await User.create({
            password: passwordHashed,
            login,
            name,
            role: 'user',
            picture: 'null.jpg'
        });

        return {
            login: userRecord.login,
            name: userRecord.name
        }
    }

    async login(login, password) {
        const userRecord = await User.findOne(login);
        
        if (!userRecord) {
            throw new Error('User not found');
        }
        else {
            const correctPassword = await argon2.verify(userRecord.password, password);
            if (!correctPassword) throw new Error('Incorrect password')
        }

        return {
            login: userRecord.login,
            name: userRecord.name
        }
        // GENERATE JWT TOKEN
    }

    generateToken(user) {
        const data = {
            _id: user._id,
            login: user.login,
            name: user.name
        }

        const signature = 'null';
        const expiration = '0h';

        return jwt.sign({ data }, signature, {expiresIn: expiration});
    }
}

module.exports = new AuthService();
