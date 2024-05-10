const Router = require("express");
const router = new Router();
const newsController = require("../controllers/news.controller.js");

router.post("/news", newsController.createNews);
router.get("/news", newsController.getNews);
router.get("/news:id", newsController.getNewsById);
router.put("/news", newsController.updateNews);
router.delete("/news:id", newsController.deleteNews);

module.exports = router;
