# [WIP] Kafka Go client as an producer and consumer example

## Requirements

* Install Golang (If you are using brew : `brew install go` should do the trick)

## Start playing

Install the depedencies :

```
go get github.com/Shopify/sarama
```

Play with the app :

```
go run main.go -topic lab42 -value "Time is an illusion. Lunchtime doubly so." $BROKER_IP:$BROKER_PORT
```

Get the value of `$BROKER_IP` and `$BROKER_PORT` in your ZK browser.
