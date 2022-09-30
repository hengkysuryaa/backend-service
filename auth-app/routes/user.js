const express = require("express");
const userRouter = express.Router();
const userController = require("../controllers/user")
const validator = require("../library/validator")
const middleware = require("./middleware/middleware")

userRouter.post("/register", userController.RegisterUser)
userRouter.post("/login", validator.validateLoginRequest(), userController.Login)
userRouter.get("/verify", middleware.checkAuthentication, userController.GetTokenInfo)

module.exports = userRouter;