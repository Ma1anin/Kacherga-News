const Router = require('express');
const router = new Router();
const eventController = require('../controllers/event.controller');

router.post('/create', eventController.createEvent);
router.get('/:id', eventController.getEventById);
router.put('/update:id', eventController.updateEvent);
router.delete('/delete:id', eventController.deleteEvent);
router.get('/', eventController.getEvents);

module.exports = router;