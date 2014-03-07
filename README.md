# ServiceStub [![Build Status](https://secure.travis-ci.org/mattlaver/servicestub.png)](http://travis-ci.org/mattlaver/servicestub)

A cross platform tool for creating dummy HTTP web services.

## How to build the code

- Clone: `git clone https://github.com/mattlaver/servicestub.git`
- Build: `go get -d && go build`


## Getting started


Say we wanted to create a fake REST API for customers:

HTTP Method: GET /customers
HTTP Method: GET /customers/:customerId




```
servicestub/
├── api/
│   ├── customers/
|   |── GET
│   |   ├── 1/
│   |   ├── GET
```


