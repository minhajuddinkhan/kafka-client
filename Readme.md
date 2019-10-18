# kafka-client


```
NAME:
   Apache Kafka Client CLI - CLI application for interacting with Apache Kafka

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     consume  consumes a message on given topic, partition and offset
     produce  publishes given message on given topic
     brokers  saves broker urls to storage
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --brokers value  Saves broker urls
   --help, -h       show help
   --version, -v    print the version
```


### Install

```bash
$ git clone https://github.com/minhajuddinkhan/kafka-client
$ cd kafka-client
```

### Set brokers

```bash
$ go run bin/kafka-client/main.go brokers set localhost:9092
```

### Consume messages from Kafka
```bash
$ go run bin/kafka-client/main.go consume --topic topic-name --brokers localhost:9092 --partition 0
```
### Publish topics to kafka
```bash
go run bin/kafka-client/main.go produce --topic 123 --value {"msg: "Hello Kafka"} --brokers localhost:9092,ec2:9092
```


## Run Kafka Locally with zoo-keeper out of the box
```bash
$ sudo docker run -p 2181:2181 -p 9092:9092 --env ADVERTISED_HOST=localhost --env ADVERTISED_PORT=9092 spotify/kafka
```