const bcrypt = require("bcrypt")

exports.encrypt = async (plaintext) => {
    const salt = await bcrypt.genSalt(10)
    return await bcrypt.hash(plaintext, salt)
}

exports.decrypt = async (plaintext, encryptedText) => {
    const isValid = await bcrypt.compare(plaintext, encryptedText)
    return isValid
}