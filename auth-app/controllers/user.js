const userService = require("../services/user")

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

exports.Login = async (req, res) => {
    try {
        var loginRes = await userService.login(req.body)
        if (loginRes.error_message != null) {
            return res.status(400).send({"message":loginRes.error_message})
        }

        return res.status(200).send({"token":loginRes.token})

    } catch(err) {
        return res.status(500).send({"message":err.message})
    }
}