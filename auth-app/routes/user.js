const express = require("express");
const userRouter = express.Router();
const userController = require("../controllers/user")
const validator = require("../library/validator")

userRouter.post("/register", userController.RegisterUser)
userRouter.post("/login", validator.validateLoginRequest(), userController.Login)

module.exports = userRouter;