const Router = require('express');
const router = new Router();
const newsController = require('@controllers/news.controller.js');

router.post('/', newsController.createNews);
router.get('/', newsController.getNews);
router.get('/:id', newsController.getNewsById);
router.put('/', newsController.updateNews);
router.delete('/:id', newsController.deleteNews);

module.exports = router;