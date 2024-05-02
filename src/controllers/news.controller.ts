import { Request, Response } from "express";
import NewsService from "../services/news.service";

class NewsController {
  public async createNews(req: Request, res: Response): Promise<Response> {
    try {
      const createdNews = await NewsService.createNews(
        req.body,
        req.files.picture
      );
      return res.json(createdNews);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async getNews(req: Request, res: Response): Promise<Response> {
    try {
      const news = await NewsService.getNews();
      return res.json(news);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async getNewsById(req: Request, res: Response): Promise<Response> {
    try {
      const news = await NewsService.getNewsById(req.params.id);
      return res.json(news);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async updateNews(req: Request, res: Response): Promise<Response> {
    try {
      const updatedNews = await NewsService.updateNews(req.body);
      return res.json(updatedNews);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async deleteNews(req: Request, res: Response): Promise<Response> {
    try {
      const deletedNews = await NewsService.deleteNews(req.params.id);
      return res.json(deletedNews);
    } catch (err) {
      return res.status(500).json(err);
    }
  }
}

export default new NewsController();
