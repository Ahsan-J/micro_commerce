var http = require('http');

var handleRequest = function(request, response) {
  console.log('App 04 Received request for URL: ' + request.url);
  response.writeHead(200);
  response.end('Hello World! from App04');
};
var www = http.createServer(handleRequest);
www.listen(8080);