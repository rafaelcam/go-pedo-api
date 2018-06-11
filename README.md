# Go PeDo API

<a href="https://codeclimate.com/github/rafaelcam/go-pedo-api/maintainability"><img src="https://api.codeclimate.com/v1/badges/1600c70ad7db756736f5/maintainability" /></a>

API in GO that upon receiving a single request, returns the following information:

* Print the numbers from 1 to 100. 
* For multiples of 3, print 'Pé' instead of the number. 
* For multiples of 5, print 'Do' instead of the number. 
* For multiples of both 3 and 5, print 'PéDo'.

## Requirements

* [Golang](https://golang.org/)
* [Dep](https://github.com/golang/dep)

## Create the Package Structure
Create the required package structure by running:
```bash
$ mkdir $GOPATH/src/github.com/rafaelcam/
```
Now unzip the `.zip` file sent to this directory and everything should be ready to run the application.

## Download Dependencies

The project uses [Dep](https://github.com/golang/dep) to manage their dependencies.
Run the following command to download them.

```bash
#!$GOPATH/src/github.com/rafaelcam/go-pedo-api
$ dep ensure
```
If all goes well you should see a `vendor` folder with the project dependencies.

## Running the Tests
Run the following command to run the tests.

```bash
#!$GOPATH/src/github.com/rafaelcam/go-pedo-api
$ go test ./...
```

## Running the Application
With the correct folder structure in your `$GOPATH` and downloaded dependencies
run the command below to run the application:

```bash
#!$GOPATH/src/github.com/rafaelcam/go-pedo-api
$ go run src/main.go
```

If all goes well you will see the message `Running server!` on the console.

## Check Application Health

**ENDPOINT**

```
GET http://localhost:3000/health
```

**RESPONSE**

Code | Response
------------ | -------------
`200` | `{"status":"OK"}`

## API

**REQUEST**

```
GET http://localhost:3000/api/v1/pedo?start={start}&end={end}
```

Example request: `http://localhost:3000/api/v1/pedo?start=1&end=100`

**REQUEST PARAMETERS**

Param | Description
--------- | -------------
`start`   | Start value of range. Must be an Integer.
`end`     | End value of range. Must be an Integer.

**RESPONSES**

Code | Response
------------ | -------------
`200` | `List of Values ​​as for example: ["1","2","Pé","4","Do"]` 
`400` | `Only positive values ​​are accepted.`
`400` | `Invalid start value. Try again.`
`400` | `Invalid end value. Try again.`
`400` | `Start value can not be greater than the End value.`