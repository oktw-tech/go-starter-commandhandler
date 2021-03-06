package configuration

import (
	"encoding/json"
	"os"
)

//JSONConfig : Json Configuration object structure
type JSONConfig struct {
	AmqpBus struct {
		AmqpURL      string `json:"amqpUrl"`
		ConsumerName string `json:"consumerName"`
		ExchangeName string `json:"exchangeName"`
		ExchangeType string `json:"exchangeType"`
		QueueName    string `json:"queueName"`
	}
}

//LoadConfiguration : Load a file configuration
func (jSONConfig *JSONConfig) LoadConfiguration(filename string) error {
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&jSONConfig)
	return err
}
