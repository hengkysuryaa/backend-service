require("dotenv").config();
const express = require("express");
var pingRouter = require("./routes/ping");

const app = express()

app.use(express.json());

// register routes here
app.use("/ping", pingRouter);

module.exports = app;