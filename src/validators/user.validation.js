const { body } = require("express-validator");

const userDataValidateChainMethod = [
  body("login")
    .exists({ checkFalsy: true })
    .withMessage("Login is required")
    .isString()
    .withMessage("User name should be string"),
  body("password")
    .exists()
    .withMessage("Password is required")
    .isString()
    .withMessage("Password should be string")
    .isLength({ min: 8 })
    .withMessage("Password should be at least 8 characters"),
];

module.exports = userDataValidateChainMethod;
