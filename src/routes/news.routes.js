const Router = require('express');
const router = new Router();
const newsController = require('../controllers/news.controller');

router.post('/create', newsController.createNews);
router.get('/:id', newsController.getNewsById);
router.put('/update:id', newsController.updateNews);
router.delete('/delete:id', newsController.deleteNews);
router.get('/', newsController.getNews);

module.exports = router;