const Router = require('express');
const router = new Router();
const newsController = require('../controllers/news.controller');

router.post('/main', newsController.createNews);
router.get('/main', newsController.getNews);
router.get('/main:id', newsController.getNewsById);
router.put('/main:id', newsController.updateNews);
router.delete('/main:id', newsController.deleteNews);

module.exports = router;