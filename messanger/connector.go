package messanger

import(
	"context"
	"fmt"
	"log"
	"os"
	"github.com/segmentio/kafka-go"
)

const (
	carEntryTopic = "carEntry"
	carExitTopic = "carExit"
	brokerAddress = "localhost:9092"
)


func AllocateParkingProducer(ctx context.Context,carId string,ownerName string) {
         
	l := log.New(os.Stdout, "kafka writer: ", 0)
 	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   carEntryTopic,
		BatchSize:    1,
		BatchTimeout: 10000,
		Logger: l,
	})

		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(carId),
			Value: []byte(carId +" "+ ownerName),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

	
}

func DeAllocateParkingProducer(ctx context.Context,carId string) {

	l := log.New(os.Stdout, "kafka writer: ", 0)
 	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   carExitTopic,
		Logger: l,
	})

		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(carId),
			Value: []byte(carId+" is left the parking"),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

}


func AllocateParkingConsumer(ctx context.Context) (string){
	l := log.New(os.Stdout, "kafka reader: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   carEntryTopic,
		GroupID: "my-group",
		StartOffset: kafka.LastOffset,
		Logger: l,
	})

		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
		r.Close()
		return string(msg.Value)

}
