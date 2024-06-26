const Event = require("../models/Event.js");

class EventService {
  async createEvent(event) {
    return await Event.create(event);
  }

  async getEvents() {
    return await Event.find();
  }

  async getEventById(_id) {
    if (!_id) throw new Error("The ID is not specified");
    return await Event.findById(_id);
  }

  async updateEvent(event) {
    if (!event._id) throw new Error("The ID is not specified");
    return await Event.findByIdAndUpdate(event._id, event, { new: true });
  }

  async deleteEvent(_id) {
    if (!_id) throw new Error("The ID is not specified");
    return await Event.findByIdAndDelete(_id);
  }
}

module.exports = new EventService();
