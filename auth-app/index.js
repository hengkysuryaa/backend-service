const http = require("http");
const app = require("./app");
var process = require('process')
const server = http.createServer(app);

const { API_PORT } = process.env;
const port = process.env.PORT || API_PORT;

// server listening 
server.listen(port, () => {
  console.log(`Server running on port ${port}`);
});

// graceful shutdown
process.on('SIGINT', () => {
  console.info("Interrupted")
  process.exit(0)
})