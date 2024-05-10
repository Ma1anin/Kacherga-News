const Router = require("express");
const router = new Router();
const userController = require("../controllers/user.controller.js");
const { check } = require("express-validator");

router.post(
  "/register",
  check("password")
    .notEmpty()
    .isLength({ min: 8 })
    .withMessage("Must be at least 8 chars long"),
  userController.register
);
router.post("/login", userController.login);
router.post("/logout", userController.logout);
router.get("/user:id", userController.getUserById);
router.put("/user", userController.updateUser);
router.delete("/user:id", userController.deleteUser);

module.exports = router;
