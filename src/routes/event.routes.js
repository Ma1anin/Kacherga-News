const Router = require("express");
const router = new Router();
const eventController = require("../controllers/event.controller.js");

router.post("/event", eventController.createEvent);
router.get("/event", eventController.getEvents);
router.get("/event:id", eventController.getEventById);
router.put("/event", eventController.updateEvent);
router.delete("/event:id", eventController.deleteEvent);

module.exports = router;
