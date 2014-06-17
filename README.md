# ServiceStub [![Build Status](https://secure.travis-ci.org/mattlaver/servicestub.png)](http://travis-ci.org/mattlaver/servicestub)

A cross platform tool for creating dummy HTTP web services. 

## How to build the code

- Clone: `git clone https://github.com/mattlaver/servicestub.git`
- Build: `go build github.com/mattlaver/servicestub`


## Getting started

- ServiceStub currently only supports GET methods. 
- Returned digests must be in a file called GET


# Example: customers

Suppose we wanted to fake the following API methods:

```
HTTP Method: GET /customers
HTTP Method: GET /customers/:customerId
```

Create a directory for customers and two customer id sub directories as:

```
├── customers/
|   ├── 1/
|   ├── 2/
```

Create a GET file in each of the directories and populate them with fake API data

```
├── customers/
|── GET
|   ├── 1/
|   ├── GET
|   ├── 2/
|   ├── GET
``` 


Place the servicestub at the root of the directory structure

Run servicestub

Open a browser and navigate to 

`http://localhost:3000/customers`

`http://localhost:3000/customers/1`

`http://localhost:3000/customers/2`


# Example: JSONP customers

JSONP supported, e.g. using jQuery:

```
$.ajax({
    type: "GET",
    dataType: "jsonp",
    jsonCallback: "jsonp",
    url: "http://localhost:3000/customers",
    success: function(data) {
        console.log(data);
    }
});
``` 

