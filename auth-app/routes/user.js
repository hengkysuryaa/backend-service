const express = require("express");
const userRouter = express.Router();
const userController = require("../controllers/user")

userRouter.post("/register", userController.RegisterUser)

module.exports = userRouter;