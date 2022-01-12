# Go RabbitMQ Crash Course!

#### Running RabbitMQ Locally with Docker

```bash
docker run -d --hostname rabbitmq --name test-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

* try opening up localhost:15672 in your browser and typing guest / guest as your username and password

#### Connecting to RabbitMQ in Go

Now that we have our RabbitMQ instance up and running, it’s time to start developing applications on top of it.

#### Creating Queues and Publishing Messages

Let’s expand on our RabbitMQ service and add the ability to create queues and publish some messages to these newly created queues.

#### Consuming Messages From a Queue

So, we’ve accomplished quite a lot so far, we’ve managed to set up a well structured app that is able to create a queue and subsequently publish messages to this queue.

In this video, we are going to be focused on how we can do the reverse and consume the messages that we have published to this queue.

```bash
docker exec -it <RabbitMQ Container ID> bash

# rabbitmqadmin publish exchange=amq.default routing_key="TestQueue" payload="Hello World"
Message published
```

```bash
$ go run main.go
Go RabbitMQ Crash Course
Connecting to RabbitMQ
Successfully connected to RabbitMQ
Successfully published message to queue
Received message: Hi
Received message: Hello World
```