require("dotenv").config();
const express = require("express");
var pingRouter = require("./routes/ping");
var userRouter = require("./routes/user");

const app = express()

app.use(express.json());

/* register routes here */
app.use("/ping", pingRouter);

// User Domain
app.use("/user", userRouter)

module.exports = app;