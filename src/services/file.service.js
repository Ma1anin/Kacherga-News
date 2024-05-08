const path = require("path");
const uuid = require("uuid");

class FileService {
  saveFile(file) {
    try {
      const fileName = uuid.v4() + ".jpg";
      const filePath = path.resolve("src", "static", fileName);
      file.mv(filePath);
      return fileName;
    } catch (err) {
      console.log(err);
    }
  }
}

module.exports = new FileService();
