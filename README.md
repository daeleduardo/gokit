### __Gokit__

A simple Golang skeleton project, for Web applications  and Microservices (over TCP/IP), using:
* Gin Gonic Framework
* MVC Archteture.
* Cache (Redis).
* JWT (Service and Middleware)
* ORM (Gorm)
* Logger
* Shell Script handler

#### __Folder Structure__
* Config _(Global configurations variables, can be use a default value or set in flag when run the app )_
* Controllers _(controllers to request handle )_
* Lib _(Libraries to use over the project, like constants, logger and other utils functions )_
* Middleware 
* Models  _(Gorm Models)_
* Public _(Template, CSS, JS , HTML and others assets files)_
* Repository _(Database and Cache connections)_
* Routes 
* Services _(To handle the bussiness logic)_

#### __How to use__
Run _main.go_  file with flags parameters, example:
``bash
go run main.go --gin-mode=dev
``
#### __Example App__
The repository have a demonstration app, can be start with command:
``bash
go run main.go
``
#### __Todo__
*   Dockerfile
*   Tests