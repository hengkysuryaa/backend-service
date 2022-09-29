var express = require("express");
var pingRouter = express.Router();

pingRouter.get("/", (req, res) => {
    res.status(200).send({"message":"Hello world!"});
});

module.exports = pingRouter;