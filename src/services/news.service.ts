import FileService from "../services/file.service";
import NewsModel from "../models/News";
import News from "../interfaces/news.interface";

class NewsService {
  public async createNews(news: News, picture: string): Promise<News> {
    const fileName = FileService.saveFile(picture);
    return await NewsModel.create({ ...news, picture: fileName });
  }

  public async getNews(): Promise<News[]> {
    return await NewsModel.find();
  }

  public async getNewsById(_id: string): Promise<News | null> {
    if (!_id) throw new Error("The ID is not specified");
    return await NewsModel.findById(_id);
  }

  public async updateNews(news: News): Promise<News> {
    if (!news._id) throw new Error("The ID is not specified");
    return await NewsModel.findByIdAndUpdate(news._id, news, { new: true });
  }

  public async deleteNews(_id: string): Promise<News | null> {
    if (!_id) throw new Error("The ID is not specified");
    return await NewsModel.findByIdAndDelete(_id);
  }
}

export default new NewsService();
