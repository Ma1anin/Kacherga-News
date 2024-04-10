const db = require("../data/db");

class NewsController {
  async createNews(req, res) {
    const {title, imageUrl, authorID} = req.body;
    const query = 'INSERT INTO News (title, imageUrl, authorID) VALUES ($1, $2, $3) RETURNING *';

    try {
      const res = await db.query(query, [title, imageUrl, authorID]);
      return res.rows[0];
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
  async getNews(req, res) {
    const query = 'SELECT * FROM News';

    try {
      const res = await db.query(query);
      return res.rows;
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
  async getNewsById(req, res) {
    const id = req.params.id;
    const query = 'SELECT * FROM News WHERE ID = $1'

    try {
      const res = await db.query(query, [id]);
      return res.rows[0];
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
  async updateNews(req, res) {
    const {title, imageUrl, id} = req.body;
    const query = 'UPDATE Event SET title = $1, imageUrl = $2 WHERE ID = $3 RETURNING *'

    try {
      const res = await db.query(query, [title, imageUrl, id]);
      return res.rows[0];
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
  async deleteNews(req, res) {
    const id = req.params.id;
    const query = 'DELETR FROM News WHERE ID = $1'

    try {
      const res = await db.query(query, [id]);
      return res.rows[0];
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
}

module.exports = new NewsController();
