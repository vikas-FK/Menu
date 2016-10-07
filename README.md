# Menu
For understanding Golang, playing around

## Refrences for GoLang
* https://github.com/dariubs/GoBooks#books
* http://thenewstack.io/make-a-restful-json-api-go/


## Setup GoLang on your local machine
* Download go from [link](https://golang.org/dl/)
* Follow the installation instruction https://golang.org/doc/install#install
* Add following lines to your bash_profile (~/.bash_profile)

Note: Ignore if done already as part of Golang installation

~~~ shell
export GOROOT="/usr/local/go"
export GOBIN=$GOPATH/bin
~~~
Add /usr/local/go/bin to you $PATH
~~~ shell
export PATH=$PATH:/usr/local/go/bin
~~~

## Get the code base from github
~~~ shell
go get https://github.com/vikas-FK/Menu
~~~
above command will fetch the code repository from git and place it in $GOPATH/src/github.com/vikas-FK/Menu

## Resolve all the GoLang dependencies
Go to your project root `$GOPATH/src/github.com/vikas-FK/Menu` and run below command to resolve all the dependencies
~~~ shell
go get ...
~~~

