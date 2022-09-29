const { Sequelize } = require("sequelize")
const User = require("../models/user")
const crypto = require("crypto")
const bcrypt = require("../library/bcrypt")

exports.registerUser = async (req) => {
    try {
        // create a random password, and encrypt it
        var password = crypto.randomBytes(2).toString('hex')
        var encryptedPassword = await bcrypt.encrypt(password)

        const [user, created] = await User.findOrCreate({
            where:{ phone:req.phone },
            defaults: {
                name:req.name,
                phone:req.phone,
                role:req.role,
                password:encryptedPassword,
                created_at:Sequelize.literal('CURRENT_TIMESTAMP')
            }
        })

        if (!created) {
            return {"user":null, "message":"Phone Number already registered in the App. Please login"}
        }

        var userDetail = {
            "name":user.dataValues.name,
            "phone":user.dataValues.phone,
            "role":user.dataValues.role,
            "password":password
        }

        return {"user":userDetail, "message":null}
        
    } catch(err) {
        throw err
    }
}