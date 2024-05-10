const Router = require("express");
const router = new Router();
const userController = require("../controllers/user.controller.js");

router.post("/register", userController.createUser);
router.post("/login", userController.login);
//router.post("/logout", userController.logout);
router.get("/user:id", userController.getUserById);
router.put("/user", userController.updateUser);
router.delete("/user:id", userController.deleteUser);

module.exports = router;
