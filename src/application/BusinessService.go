package application

import (
	"encoding/json"
	"fmt"
	crosscutting "go-starter-commandhandler/src/crosscutting/error"
	"go-starter-commandhandler/src/crosscutting/messaging"
	"go-starter-commandhandler/src/domain/event"
	"go-starter-commandhandler/src/domain/model"

	"time"
)

//BusinessService : Business service
func BusinessService(message string) {

	fmt.Println(model.DomainModel{Propertie: message})
	raisedDomainEvent := event.DomainEvent{DomainModel: model.DomainModel{Propertie: message}, PublicationDateTimeUnix: time.UnixDate}
	body, err := json.Marshal(raisedDomainEvent)
	rabbitMqProvider := messaging.RabbitMqProvider{}
	rabbitMqProvider.PublishOneMessage("raisedDomainEvent", []byte(body))
	crosscutting.RaiseError("Failed to publish a message", err)
}
