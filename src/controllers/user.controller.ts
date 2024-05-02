import { Request, Response } from "express";
import UserService from "../services/user.service";

class UserController {
  public async createUser(req: Request, res: Response): Promise<Response> {
    try {
      const createdUser = await UserService.createUser(req.body);
      return res.json(createdUser);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async getUserById(req: Request, res: Response): Promise<Response> {
    try {
      const user = await UserService.getUserById(req.params.id);
      return res.json(user);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async updateUser(req: Request, res: Response): Promise<Response> {
    try {
      const updatedUser = await UserService.updateUser(req.body);
      return res.json(updatedUser);
    } catch (err) {
      return res.status(500).json(err);
    }
  }

  public async deleteUser(req: Request, res: Response): Promise<Response> {
    try {
      const deletedUser = await UserService.deleteUser(req.params.id);
      return res.json(deletedUser);
    } catch (err) {
      return res.status(500).json(err);
    }
  }
}

export default new UserController();
