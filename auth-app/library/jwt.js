const jwt = require("jsonwebtoken")
const jwtSecret = process.env.JWT_SECRET

exports.sign = async (data) => {
    return jwt.sign(data, jwtSecret)
}

exports.verify = async (token) => {
    return jwt.verify(token, jwtSecret)
}