# gorillamux

Simple HTTP server using gorilla mux

## Usage:

* `./gorillamux`
Starts the server on `localhost:9000`

* `./gorillamux -port="8080"`
Starts the server on `localhost:8080`
_NOTE: `port` needs to be specified as a string in the arguments_

* `./gorillamux --help`
Displays the usage information

## Testing:

Use `curl` to test GET, POST and DELETE requests as follows:

* `curl localhost:9000` should return a welcome message

* `curl localhost:9000/products` should return all products stored in memory

* `curl localhost:9000/product/P001` should return a product with id `P001` if present else return a message indicating no such product present

* `curl -H "Content-Type: application/json" -X POST -d '{"name":"Prod01","desc":"Product 01","qty":20}' http://localhost:9000/product` should add a new product

* `curl -H "Content-Type: application/json" -X PUT -d '{"name":"Prod02","desc":"Product 02","qty":40}' http://localhost:9000/product/P001` should replace existing product with id `P001` if it exists else return a message indicating no such product present

* `curl -X DELETE http://localhost:9000/product/P001` should remove the product with id `P001` if present else return a message indicating no such product present

