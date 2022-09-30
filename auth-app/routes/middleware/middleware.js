const jwt = require("../../library/jwt")

exports.checkAuthentication = async (req, res, next) => {
    try {
        const jwtToken = req.headers.authorization
        if (!jwtToken) {
            return res.status(403).send({"message":"A token is required for authentication"})
        }

        req.decodedToken = await jwt.verify(jwtToken)
        next()

    } catch(err) {
        return res.status(401).send({"message":"Invalid Token"})
    }
}