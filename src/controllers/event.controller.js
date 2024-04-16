const EventService = require("../services/event.service");

class EventController {
  async createEvent(req, res) {
    try {
      const createdEvent = await EventService.createEvent(req.body);
      return res.json(createdEvent);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async getEvents(req, res) {
    try {
      const events = await EventService.getEvents();
      return res.json(events);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async getEventById(req, res) {
    try {
      const event = await EventService.getEventById(req.params.id);
      return res.json(event);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async updateEvent(req, res) {
    try {
      const updatedEvent = await EventService.updateEvent(req.body);
      return res.json(updatedEvent);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async deleteEvent(req, res) {
    try {
      const deletedEvent = await EventService.deleteEvent(req.params.id);
      return res.json(deletedEvent);
    } catch (err) {
      res.status(500).json(err);
    }
  }
}

module.exports = new EventController();
