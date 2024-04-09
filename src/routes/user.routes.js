const Router = require('express');
const router = new Router();
const userController = require('../controllers/user.controller');

router.post('/main', userController.createUser);
router.get('/main:id', userController.getUserById);
router.put('/main:id', userController.updateUser);
router.delete('/main:id', userController.deleteUser);

module.exports = router;