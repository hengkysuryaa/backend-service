const joi = require("joi")

var validate = (schema) => (req, res, next) => {
    const { error }  = schema.validate(req.body)
    if (error) {
        return res.status(400).send({message:error.details[0].message})
    }

    next()
}

module.exports = {
    validateLoginRequest: () => {
        const schema = joi.object().keys({
            phone: joi.string().required(),
            password: joi.string().required()
        })
        
        return validate(schema)
    }
}