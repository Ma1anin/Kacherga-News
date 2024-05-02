// import FileService from "../services/file.service";
import UserModel from "../models/User";
import User from "../interfaces/user.interface";

class UserService {
  public async createUser(user: User): Promise<User> {
    return await UserModel.create({ ...user, picture: "null.jpg" });
  }

  public async getUserById(_id: string): Promise<User | null> {
    if (!_id) throw new Error("The ID is not specified");
    return await UserModel.findById(_id);
  }

  public async updateUser(user: User): Promise<User> {
    if (!user._id) throw new Error("The ID is not specified");
    // if (picture) {
    //   const fileName = FileService.saveFile(picture);
    // }
    return await UserModel.findByIdAndUpdate(user._id, user, { new: true });
  }

  public async deleteUser(_id: string): Promise<User> {
    if (!_id) throw new Error("The ID is not specified");
    return await UserModel.findByIdAndDelete(_id);
  }
}

export default new UserService();
