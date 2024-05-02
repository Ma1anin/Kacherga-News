import path from "path";
import { v4 as uuidv4 } from "uuid";

class FileService {
  saveFile(file) {
    try {
      const fileName = uuidv4.v4() + ".jpg";
      const filePath = path.resolve("src", "static", fileName);
      file.mv(filePath);
      return fileName;
    } catch (err) {
      console.log(err);
    }
  }
}

export default new FileService();
