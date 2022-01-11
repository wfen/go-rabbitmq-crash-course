# Go RabbitMQ Crash Course!

#### Running RabbitMQ Locally with Docker

```bash
docker run -d --hostname rabbitmq --name test-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

* try opening up localhost:15672 in your browser and typing guest / guest as your username and password

#### Connecting to RabbitMQ in Go

Now that we have our RabbitMQ instance up and running, it’s time to start developing applications on top of it.

