# Safe Zone 
Safe Zone is service like simple calculator for you hard to use.

## Quick Start
### Installation Required 
### Install Golang and Dependency
I use Golang go `go1.11.4 darwin/amd64` you must be install Go in your PC. 

In go use Dependency for package manager or ignore if you using go > v11.** please see https://github.com/golang/dep.

#### Install Docker (if you using Docker)
I Use Docker for deployment my apps and integration app. Docker is very simple for running this app.

So please install docker into your PC https://docs.docker.com/install

### Deploy and run 
#### Run With Golang 

Build go Apps:
    
    $ go build 
    
 
Running for sum *note(--x for variable x and --y for variable y this variable is required) :

    $ ./safeZone sum --x=200000 --y=1000000 
    
Running for multiple *note(--x for variable x and --y for variable y this variable is required) :

    $ ./safeZone multiple --x=200000 --y=1000000 
   
Running for prime number *note(--n for variable n this variable is required) :

    $ ./safeZone prime --n=4

Running for fibonacci *note(--n for variable n this variable is required) :

    $ ./safeZone sum --n=4
    
##### Run with Docker 
So Easy and you does'n effort for running this apps if using docker. 

Build Docker:

    $ docker build -t safe_zone . 
    
Run Docker:

    $  docker run safe_zone sum --x 100 --y 200
    
  *note: command `sum --x 100 --y 200` following flags running with Go. 
    
### Unit Test and Mock
To run unit test for the project :

    $ go test $(go list ./... | grep -v /vendor/) -cover
        
### Code Style

In golang they already define the convention style https://golang.org/doc/effective_go.html