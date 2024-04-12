const NewsService = require("../services/news.service");

class NewsController {
  async createNews(req, res) {
    try {
      const createdNews = await NewsService.createNews(req.body);
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
      const news = NewsService.getNewsById(req.params.id);
      return res.json(news);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async updateNews(req, res) {
    try {
      const updatedNews = NewsService.updateNews(req.body);
      return res.json(updatedNews);
    } catch (err) {
      res.status(500).json(err);
    }
  }

  async deleteNews(req, res) {
    try {
      const deletedNews = NewsService.deleteNews(req.params.id);
      return res.json(deletedNews);
    } catch (err) {
      res.status(500).json(err);
    }
  }
}

module.exports = new NewsController();
