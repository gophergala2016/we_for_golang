# 'Schema generator' for mysql
This is an attempt to generate a bootstrap file which can be used as a base for 
ORM.

## Use case
This app can be handy when we have a legacy application and wants to use ORM or
when we have a data first, this will pull out its schema in golang Struct 


## Few Examples
	Gorm

## How to Use
This is a terminal based utility which accept mysql username, password, database 
& file name to generate a Struct.

	./we_for_golang -u <USERNAME> -p <PASSWORD> -d <DATABASE> -f ~/workspace/my_project_schema.go


## Environment set up

Download the latest stable Go distribution(go1.5.3) from the golang site and follow
the instruction mentioned.
  
	https://golang.org/doc/install
	https://golang.org/dl/

## Building Source Code

Open terminal at application directory, execute "go build" command. This will build 
entire application. More instruction can be found at

	https://golang.org/cmd/go/   