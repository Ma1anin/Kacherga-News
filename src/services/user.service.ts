import FileService from "../services/file.service";
import UserModel from "../models/User";

interface User {
  _id: string;
  login: string;
  name: string;
  role: string;
  picture: string;
}

class UserService {
  public async createUser(user: User): Promise<User> {
    return await User.create({ ...user, picture: "null.jpg" });
  }

  public async getUserById(_id: string): Promise<User | null> {
    if (!_id) throw new Error("The ID is not specified");
    return await User.findById(_id);
  }

  public async updateUser(user: User): Promise<User> {
    if (!user._id) throw new Error("The ID is not specified");
    // if (picture) {
    //   const fileName = FileService.saveFile(picture);
    // }
    return await User.findByIdAndUpdate(user._id, user, { new: true });
  }

  public async deleteUser(_id: string): Promise<User> {
    if (!_id) throw new Error("The ID is not specified");
    return await User.findByIdAndDelete(_id);
  }
}

export default new UserService();
