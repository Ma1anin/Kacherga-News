const db = require("../data/db");

class EventController {
  async createEvent(req, res) {
    const { title, authorID } = req.body;
    const query = "INSERT INTO Event (title, authorID) VALUES ($1, $2) RETURNING *";

    try {
      const res = db.query(query, [title, authorID]);
      return res.rows[0];
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
  async getEvents(req, res) {
    const query = "SELECT * FROM Event";

    try {
      const res = db.query(query);
      return res.rows[0];
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
  async getEventById(req, res) {
    const id = req.params.id;
    const query = "SELECT * FROM Event WHERE ID = $1";

    try {
      const res = db.query(query, [id]);
      return res.rows[0];
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
  async updateEvent(req, res) {
    const {title, id} = req.body;
    const query = "UPDATE Event SET title = $1 WHERE ID = $2 RETURNING *";

    try {
      const res = db.query(query, [title, id]);
      return res.rows[0];
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
  async deleteEvent(req, res) {
    const id = req.params.id;
    const query = "DELETE FROM Event WHERE ID = $1";

    try {
      const res = db.query(query, [id]);
      return res.rows[0];
    } catch (err) {
      console.log(err);
      throw err;
    }
  }
}

module.exports = new EventController();
