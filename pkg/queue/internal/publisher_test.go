package kafka

import (
	"context"
	"testing"
	"time"
)

var (
	testTopic   = "red_packet_data_topic"
	testGroup   = "test-event-group"
	testBrokers = []string{"10.10.8.16:9092"}
)

//func TestMain(m *testing.M) {
//	// to create topics when auto.create.topics.enable='true'
//	ctx := context.Background()
//	ctx, cancel := context.WithTimeout(ctx, time.Second)
//	defer cancel()
//	_, err := kafka.DialLeader(ctx, "tcp", testBrokers[0], testTopic, 0)
//	if err != nil {
//		panic(err)
//	}
//	os.Exit(m.Run())
//}

func TestPublisher(t *testing.T) {
	pub,err := NewPublisher(testBrokers)
	if err != nil {
		t.Fatal(err)
	}
	defer pub.Close()
	if err := pub.Publish(context.Background(), Event{Topic:testTopic,Key: "key1", Payload: []byte("valu31231e12222")}); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 2)
}
