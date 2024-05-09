# GO SDK
## Software Development Kit For Golang


[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

GO-SDK is package that support devlopers for programing go project,
Use for both API Gateway and Servies in backend development.

## Features

- **GRPC** Conection, Protobuf generate
- Base repository for **PostgreSQL**
- Base repository for **Redis**
- Package for intergrate with **AWS (KMS, S3, ...)**
- Common struct type and util
- **RabbitMQ** 
- **Elastic Search** 

This SDK work best with microservices development that use grpc to communicate.
For example implement, see this below repositories:
- [API Gateway](https://github.com/hadanhtuan/dorm-go-gateway)
- [Services](https://github.com/hadanhtuan/dorm-go-services)

## Tech

GO-SDK uses a number of open source projects to work properly:

- [Postgre Driver](gorm.io/driver/postgres) - Driver for connect to postgreSQL
- [AWS SDK](github.com/aws/aws-sdk-go-v2) - awesome sdk prodvide AWS client and config
- [GRPC](google.golang.org/grpc) - GRPC package suppor create connection

And of course GO-SDK itself is open source with a [public repository](https://github.com/hadanhtuan/go-sdk) on GitHub.

## Installation

Install the package in your repository and start the server.

```sh
cd cmd/main.go
go get https://github.com/hadanhtuan/go-sdk
```


## Development

Want to contribute? Great!

GO-SDK uses golang for fast developing.
Make a change in your file and instantaneously see your updates!
