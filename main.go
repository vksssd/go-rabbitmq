package main

import (
	receive "github.com/vksssd/go-rabbitmq/receive"
	send "github.com/vksssd/go-rabbitmq/send"
)


func main(){
	go send.init()
	go receive.init()
}