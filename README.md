# 'Schema generator' for mysql
This is an attempt to generate a bootstrap file which can be used as a base for 
ORM.

## Use case
This app can be handy when we have a legacy application and wants to use ORM or
when we have a data first, this will pull out its schema in golang Struct 


## Few Examples
	Gorm

On executing the application a http server will be launched at port :8000

## Environment set up

Download the latest stable Go distribution(go1.5.3) from the golang site and follow
the instruction mentioned.
  
	https://golang.org/doc/install
	https://golang.org/dl/

## Building Source Code

Open terminal at application directory, execute "go build" command. This will build 
entire application. More instruction can be found at

	https://golang.org/cmd/go/   