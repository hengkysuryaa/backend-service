const userService = require("../services/user");

exports.RegisterUser = async(req, res) => {
    try {
        var registerRes = await userService.registerUser(req.body)
        if (registerRes.message != null) {
            return res.status(400).send({"message":registerRes.message})
        }

        return res.status(200).send({"user":registerRes.user})

    } catch(err) {
        return res.status(500).send({"message":err.message})
    }
}