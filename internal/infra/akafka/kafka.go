package akafka

import(
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Ler e receber dados do Kafka
func Consume(topics []string, servers string, msgChan chan *kafka.Message)  {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id": "api-golang",
		"auto.offset.reset": "earliest",
	})
	
	if err != nil {
		panic(err)
	}

	//Pegar as mensagens e jogar no canal de mensagens
	kafkaConsumer.SubscribeTopics(topics, nil)
	for{
		//-1 timeOut negativo
		if msg, err := kafkaConsumer.ReadMessage(-1); err == nil {
			msgChan <- msg
		}
	}
}