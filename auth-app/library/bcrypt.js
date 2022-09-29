const bcrypt = require("bcrypt")

exports.encrypt = async (plaintext) => {
    const salt = await bcrypt.genSalt(10)
    return await bcrypt.hash(plaintext, salt)
}