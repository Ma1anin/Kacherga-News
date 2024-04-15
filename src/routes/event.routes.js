const Router = require('express');
const router = new Router();
const eventController = require('../controllers/event.controller');

router.post('/', eventController.createEvent);
router.get('/:id', eventController.getEventById);
router.put('/', eventController.updateEvent);
router.delete('/:id', eventController.deleteEvent);
router.get('/', eventController.getEvents);

module.exports = router;