package application

import (
	"fmt"
	"go-starter-commandhandler/src/crosscutting/messaging"
)

//CommandListenerService : Command listener service
func CommandListenerService(commandName string) {

	fmt.Println("Listening : {0}", commandName)

	rabbitMqProvider := messaging.RabbitMqProvider{}
	rabbitMqProvider.ListenExchange("#")
}
