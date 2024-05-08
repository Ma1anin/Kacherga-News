const Router = require("express");
const router = new Router();
const userController = require("../controllers/user.controller.js");

router.post("/", userController.createUser);
router.get("/:id", userController.getUserById);
router.put("/", userController.updateUser);
router.delete("/:id", userController.deleteUser);

module.exports = router;
