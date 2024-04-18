const FileService = require("../services/file.service.js");
const News = require("../models/News.js");

class NewsService {
  async createNews(news, picture) {
    const fileName = FileService.saveFile(picture);
    return await News.create({ ...news, picture: fileName });
  }

  async getNews() {
    return await News.find();
  }

  async getNewsById(_id) {
    if (!_id) throw new Error("The ID is not specified");
    return await News.findById(_id);
  }

  async updateNews(news) {
    if (!news._id) throw new Error("The ID is not specified");
    return await News.findByIdAndUpdate(news._id, news, { new: true });
  }

  async deleteNews(_id) {
    if (!_id) throw new Error("The ID is not specified");
    return await News.findByIdAndDelete(_id);
  }
}

module.exports = new NewsService();
