import News from "../models/News.js";

class NewsService {
  async createNews(news, picture) {
    return await News.create(news);
  }

  async getNews() {
    return await News.find();
  }

  async getNewsById(id) {
    return await News.findById(id);
  }

  async updateNews(id, news) {
    return await News.findByIdAndUpdate(id, news);
  }

  async deleteNews(id, news) {
    return await News.findByIdAndDelete(id, news);
  }
}

export default new NewsService();