package messaging

import (
	crosscutting "go-starter-commandhandler/src/crosscutting/error"
	"go-starter-commandhandler/src/infrastructure/configuration"
	"log"

	"github.com/streadway/amqp"
)

//RabbitMqProvider : Provider to interact with a RabbitMq backbone
type RabbitMqProvider struct {
	AmqpURL, ExchangeName string
}

//GetConnection : Return an amqp connection from url
func (RabbitMqProvider) GetConnection() *amqp.Connection {
	jSONConfig := configuration.JSONConfig{}
	err := jSONConfig.LoadConfiguration("config.json")
	AmqpURL := jSONConfig.AmqpBus.AmqpURL

	connection, err := amqp.Dial(AmqpURL)
	crosscutting.RaiseError("Failed to connect to RabbitMQ", err)
	return connection
}

//PublishOneMessage : Publish one message on the message bus
func (rabbitMqProvider *RabbitMqProvider) PublishOneMessage(messageName string, messageBody []byte) {
	jSONConfig := configuration.JSONConfig{}
	err := jSONConfig.LoadConfiguration("config.json")
	ExchangeName := jSONConfig.AmqpBus.ExchangeName

	connection := rabbitMqProvider.GetConnection()
	defer connection.Close()

	channel, err := connection.Channel()
	crosscutting.RaiseError("Failed to open a channel", err)
	defer channel.Close()

	err = channel.Publish(
		ExchangeName, // exchange
		messageName,  // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        messageBody,
		})
	crosscutting.RaiseError("Failed to publish a message", err)

	log.Printf(" [x] Sent %s", messageBody)
}

//DeclareExchangeOnChannel : Declare a new exchange
func (rabbitMqProvider *RabbitMqProvider) DeclareExchangeOnChannel(nameExchange string, typeExchange string, durable bool, autoDeleted bool, internal bool, noWait bool, arguments amqp.Table) {
	connection := rabbitMqProvider.GetConnection()
	channel, err := connection.Channel()

	crosscutting.RaiseError("Failed to open a channel", err)

	defer connection.Close()
	defer channel.Close()

	err = channel.ExchangeDeclare(
		nameExchange, // name
		typeExchange, // type
		durable,      // durable
		autoDeleted,  // auto-deleted
		internal,     // internal
		noWait,       // no-wait
		arguments,    // arguments
	)
	crosscutting.RaiseError("Failed to declare an exchange", err)
}
