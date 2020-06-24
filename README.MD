# RabbitMQ and Go Services

## Docs
Produce event or message using rabbitMQ server https://www.rabbitmq.com/

## Install
#### A. Requirement
1. RabbitMQ Server Installed or on Cloud Server (install on local https://www.rabbitmq.com/download.html)
2. Go Version => 1.14
3. go mod Dependency Manager Tool Golang

#### B. Installation
##### 1. clone this repo
```$xslt
    git clone https://github.com/huzairuje/go-rabbitmq
```

##### 2. cd to the directory
```$xslt
    cd ~/GO/SRC/DIRECTORY/go-rabbitmq
```

#### C. Run The Server
##### 1. make sure the RabbitMQ Server is running
or you can test it with access the admin API with
```$xslt
    curl -i -X POST \
    --url http://localhost:15672/#/
```

or you can start the rabbitMQ server with
 ```$xslt
     sudo systemctl start rabbitmq-server
 ```
##### 2. run the publisher
cd to user service
```$xslt
    cd ~/GO/SRC/DIRECTORY/go-rabbitmq/publisher
```
and run the service with
```$xslt
    go run main.go
```

##### 3. run the consumer
cd to user service
```$xslt
    cd ~/GO/SRC/DIRECTORY/go-rabbitmq/publisher
```
and run the service with
```$xslt
    go run main.go
```