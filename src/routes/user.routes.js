const Router = require("express");
const router = new Router();
const userController = require("../controllers/user.controller.js");
const userDataValidateChainMethod = require("../validators/user.validation.js");

router.post("/register", userDataValidateChainMethod, userController.register);
router.post("/login", userDataValidateChainMethod, userController.login);
router.post("/logout", userController.logout);
router.get("/user:id", userController.getUserById);
router.put("/user", userController.updateUser);
router.delete("/user:id", userController.deleteUser);

module.exports = router;
