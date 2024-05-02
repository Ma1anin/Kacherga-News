import Router from "express";
import userController from "../controllers/user.controller";

const router = new Router();

router.post("/", userController.createUser);
router.get("/:id", userController.getUserById);
router.put("/", userController.updateUser);
router.delete("/:id", userController.deleteUser);

export default router;
