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
	log.Println("bro plz allocate")
         
	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
 	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   carEntryTopic,
		BatchSize:    1,
		BatchTimeout: 10000,
		// assign the logger to the writer
		Logger: l,
	})


		
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(carId),
			// create an arbitrary message payload for the value
			Value: []byte(carId +" "+ ownerName),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
		log.Println("bro plz allocate plzzz")

	
}

func DeAllocateParkingProducer(ctx context.Context,carId string) {

	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
 	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   carExitTopic,
		// assign the logger to the writer
		Logger: l,
	})

		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(carId),
			// create an arbitrary message payload for the value
			Value: []byte(carId+" is left the parking"),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

}


func AllocateParkingConsumer(ctx context.Context) (string){
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   carEntryTopic,
		GroupID: "my-group",
		StartOffset: kafka.LastOffset,
		// assign the logger to the reader
		Logger: l,
	})

		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
		r.Close()
		return string(msg.Value)

}
