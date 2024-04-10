const Router = require('express');
const router = new Router();
const userController = require('../controllers/user.controller');

router.post('/create', userController.createUser);
router.get('/:id', userController.getUserByLogin);
router.put('/update:id', userController.updateUser);
router.delete('/delete:id', userController.deleteUser);

module.exports = router;