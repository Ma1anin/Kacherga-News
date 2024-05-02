import Router from "express";
import newsController from "../controllers/news.controller";

const router = new Router();

router.post("/", newsController.createNews);
router.get("/", newsController.getNews);
router.get("/:id", newsController.getNewsById);
router.put("/", newsController.updateNews);
router.delete("/:id", newsController.deleteNews);

export default router;
