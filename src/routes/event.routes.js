const Router = require('express');
const router = new Router();
const eventController = require('../controllers/event.controller');

router.post('/main', eventController.createEvent);
router.get('/main', eventController.getEvents);
router.get('/main:id', eventController.getEventById);
router.put('/main:id', eventController.updateEvent);
router.delete('/main:id', eventController.deleteEvent);

module.exports = router;