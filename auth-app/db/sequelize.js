var Sequelize = require('sequelize');

const dbConnection = new Sequelize(
    "auth",
    null,
    null,
    {
        dialect: "sqlite",
        storage:"auth.db",
        pool:{
            idle:1000,
            max:10
        }
    }
)

dbConnection.authenticate().then(() => {
    console.log("Connection established")
}).catch(err => {
    console.error('Unable to connect to the database: ', err)
})

module.exports = dbConnection