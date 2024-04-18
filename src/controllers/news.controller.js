const NewsService = require("@services/news.service.js");

class NewsController {
  async createNews(req, res) {
    try {
      const createdNews = await NewsService.createNews(
        req.body,
        req.files.picture
      );
      return res.json(createdNews);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async getNews(req, res) {
    try {
      const news = await NewsService.getNews();
      return res.json(news);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async getNewsById(req, res) {
    try {
      const news = await NewsService.getNewsById(req.params.id);
      return res.json(news);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async updateNews(req, res) {
    try {
      const updatedNews = await NewsService.updateNews(req.body);
      return res.json(updatedNews);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async deleteNews(req, res) {
    try {
      const deletedNews = await NewsService.deleteNews(req.params.id);
      return res.json(deletedNews);
    } catch (err) {
      res.status(500).json(err);
    }
  }
}

module.exports = new NewsController();
