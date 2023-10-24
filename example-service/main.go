package main

import (
	"context"
	"math/rand"

	"github.com/jinyao-lee/kafka-logger/example-service/internal"
	"github.com/jinyao-lee/kafka-logger/example-service/pb"
)

func main() {
	producer, err := internal.NewSaramaProducer()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 3; i++ {
		msg := pb.MySampleKafkaMessage{
			FirstName: "James",
			Contacts: []*pb.Contact{
				{
					ContactType: pb.ContactType_EMAIL,
					Value:       "james@somedomain.com",
				},
				{
					ContactType: pb.ContactType_PHONE,
					Value:       "+1(853)369-1830",
				},
			},
			LuckyNumber: rand.Uint32(),
		}
		if err := producer.ProduceMessage(context.Background(), &msg); err != nil {
			panic(err)
		}
	}

}
