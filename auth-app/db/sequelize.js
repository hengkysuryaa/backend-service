var Sequelize = require('sequelize');
var path = require('path')

const dbConnection = new Sequelize(
    "auth",
    null,
    null,
    {
        dialect: "sqlite",
        storage: path.resolve(__dirname, "auth.db"),
        pool:{
            idle:1000,
            max:10
        },
        define: {
            timestamps: false
        },
        sync: {
            force: true
        }
    }
)

dbConnection.authenticate().then(() => {
    console.log("Connection established")
}).catch(err => {
    console.error('Unable to connect to the database: ', err)
})

module.exports = dbConnection