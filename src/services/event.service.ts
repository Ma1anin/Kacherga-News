const EventModel = require("../models/Event.js");

interface Event {
  _id: string;
  title: string;
  content: string;
  createdAt: Date;
  authorID: string;
}

class EventService {
  public async createEvent(event: Event): Promise<Event> {
    return await EventModel.create(event);
  }

  public async getEvents(): Promise<Event[]> {
    return await EventModel.find();
  }

  public async getEventById(_id: string): Promise<Event | null> {
    if (!_id) throw new Error("The ID is not specified");
    return await EventModel.findById(_id);
  }

  public async updateEvent(event: Event): Promise<Event> {
    if (!event._id) throw new Error("The ID is not specified");
    return await EventModel.findByIdAndUpdate(event._id, event, { new: true });
  }

  public async deleteEvent(_id: string): Promise<Event | null> {
    if (!_id) throw new Error("The ID is not specified");
    return await EventModel.findByIdAndDelete(_id);
  }
}

export default new EventService();
