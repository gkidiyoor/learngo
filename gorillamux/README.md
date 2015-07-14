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

* `curl localhost:9000/product/01` should return a product with id `01` if present else return a message indicating no such product present

* `curl -H "Content-Type: application/json" -X POST -d '{"id":"01","desc":"Product 01","qty":20}' http://localhost:9000/product/01` should add a new product (or replace existing one) with id `01`

* `curl -X DELETE http://localhost:9000/product/01` should remove the product with id `01` if present 

