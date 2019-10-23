# Webserve-API
This starting API service is quite straight forward.
To build the binary,

1. pull the code
    ```
        git clone git@github.com:goyeh/gowebserve.git
        or
        git pull git@github.com:goyeh/gowebserve.git       
    ```
1. inside the home fold run the make program.
    ```
        make
    ```
    you should get something the looks like the following. the http frame work and the env file reader
    ```
    make
    go clean
    rm -f 
    go get github.com/valyala/fasthttp
    go get github.com/joho/godotenv
    go build -o connect-api -v
    ```
1. simple run the binary file.
    ```
    connect-api
    ```
1. test the program
    ```
    curl localhost:8000/login

    ```
    you should see something like:
    ```
    Request method is "GET"
    RequestURI is "/test"
    Requested path is "/test"
    Host is "localhost:8000"
    Query string is ""
    User-Agent is "curl/7.54.0"
    Connection has been established at 2019-10-23 12:16:23.595999 +0800 HKT m=+75.066635474
    Request has been started at 2019-10-23 12:16:23.596017 +0800 HKT m=+75.066653094
    Serial request number for the current connection is 1
    Your ip is "127.0.0.1"

    ```
1. another test with paramaters
    ```
    curl localhost:8000/test?fred=bob
    ```
    the result should present the method you can use to pass paramters to the web site.
    ```
    Request method is "GET"
    RequestURI is "/test?fred=bob"
    Requested path is "/test"
    Host is "localhost:8010"
    Query string is "fred=bob"
    User-Agent is "curl/7.54.0"
    Connection has been established at 2019-10-23 12:17:59.502669 +0800 HKT m=+170.970428060
    Request has been started at 2019-10-23 12:17:59.502684 +0800 HKT m=+170.970443046
    Serial request number for the current connection is 1
    Your ip is "127.0.0.1"

    ```
    
## The API documentation follows

### Login

### Register

### Upload/KYC

### Other stuff.


## Questions or comments
Feel free to message me, or send email to tipton[at]mysecretroom.com
I am stationed in Hong Kong, and enjoy programming challenges.