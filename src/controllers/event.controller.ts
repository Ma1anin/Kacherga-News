import { Request, Response } from "express";
import EventService from "../services/event.service";

class EventController {
  public async createEvent(req: Request, res: Response): Promise<Response> {
    try {
      const createdEvent = await EventService.createEvent(req.body);
      return res.json(createdEvent);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async getEvents(req: Request, res: Response): Promise<Response> {
    try {
      const events = await EventService.getEvents();
      return res.json(events);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async getEventById(req: Request, res: Response): Promise<Response> {
    try {
      const event = await EventService.getEventById(req.params.id);
      return res.json(event);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async updateEvent(req: Request, res: Response): Promise<Response> {
    try {
      const updatedEvent = await EventService.updateEvent(req.body);
      return res.json(updatedEvent);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async deleteEvent(req: Request, res: Response): Promise<Response> {
    try {
      const deletedEvent = await EventService.deleteEvent(req.params.id);
      return res.json(deletedEvent);
    } catch (err) {
      return res.status(500).json(err);
    }
  }
}

export default new EventController();
