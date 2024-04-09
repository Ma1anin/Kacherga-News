const db = require('../data/db');

class UserController {
  async createUser(req, res) {
    const {login, name, role} = req.body;
    const query = `INSERT INTO User (login, name, role) VALUES ($1, $2, $3) RETURNING *`;

    try {
      const res = await db.query(query, [login, name, role]);
      return res.rows[0];
    } catch (err) {
      console.error(err);
      throw err;
    }
  }
  async getUserByLogin(req, res) {
    const login = req.params.login;
    const query = `SELECT * FROM User WHERE login = $1`;

    try {
      const res = await db.query(query, [login]);
      return res.rows[0];
    } catch (err) {
      console.error(err);
      throw err;
    }
  }
  async updateUser(req, res) {
    const {login, fullName, avatarUrl} = req.body;
    const query = `UPDATE User SET login = $1, fullName = $2, avatarUrl = $3 WHERE login = $1 RETURNING *`;

    try {
      const res = await db.query(query, [login, fullName, avatarUrl]);
      return res.rows[0];
    } catch (err) {
      console.error(err);
      throw err;
    }
  }
  async deleteUser(req, res) {
    const login = req.params.login;
    const query = `DELETE FROM User WHERE login = $1`;

    try {
        const res = await db.query(query, [login]);
        return res.rows[0];
    } catch (err) {
        console.log(err);
        throw err;
    }
  }
}

module.exports = new UserController();
