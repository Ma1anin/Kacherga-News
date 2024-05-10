const { body } = require("express-validator");

const userDataValidateChainMethod = [
  body("login")
    .exists({ checkFalsy: true })
    .withMessage("Login is required")
    .isString()
    .withMessage("User login should be string"),
  body("password")
    .exists()
    .withMessage("Password is required")
    .isString()
    .withMessage("Password should be string")
    .isLength({ min: 8 })
    .withMessage("Password should be at least 8 characters"),
    body("fullName")
    .optional()
    .notEmpty()
    .withMessage("User name shouldn't be empty")
    .isString()
    .withMessage("User name should be string")
    .isLength({ min: 5 })
    .withMessage("Name should be at least 5 characters")
];

module.exports = userDataValidateChainMethod;
