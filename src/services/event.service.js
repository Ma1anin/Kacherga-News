import Event from "../models/Event.js";

class EventService {
  async createEvent(event) {
    return await Event.create(event);
  }

  async getEvent() {
    return await Event.find();
  }

  async getEventById(id) {
    return await Event.findById(id);
  }

  async updateEvent(id, event) {
    return await Event.findByIdAndUpdate(id, event);
  }

  async deleteEvent(id, event) {
    return await Event.findByIdAndDelete(id, event);
  }
}

export default new EventService();