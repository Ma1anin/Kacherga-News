const News = require("../models/News.js");

class NewsService {
  async createNews(news, picture) {
    return await News.create(news);
  }

  async getNews() {
    return await News.find();
  }

  async getNewsById(_id) {
    if (!_id) throw new Error('The ID is not specified');
    return await News.findById(_id);
  }

  async updateNews(news) {
    if (!news._id) throw new Error('The ID is not specified');
    return await News.findByIdAndUpdate(news._id, news, {new: true});
  }

  async deleteNews(news) {
    if (!news._id) throw new Error('The ID is not specified');
    return await News.findByIdAndDelete(news._id, news);
  }
}

module.exports = new NewsService();