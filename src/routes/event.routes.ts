import Router from "express";
import eventController from "../controllers/event.controller";

const router = new Router();

router.post("/", eventController.createEvent);
router.get("/:id", eventController.getEventById);
router.put("/", eventController.updateEvent);
router.delete("/:id", eventController.deleteEvent);
router.get("/", eventController.getEvents);

export default router;
