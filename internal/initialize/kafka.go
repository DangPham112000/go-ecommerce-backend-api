package initialize

import (
	"context"
	"log"
	"time"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/segmentio/kafka-go"
)

const (
	kafkaURL   = "localhost:9092"
	kafkaTopic = "otp-auth-topic"
)

func testConn() {
	topic := "test-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, partition)
	if err != nil {
		log.Fatal("Failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("Failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("Failed to close writer:", err)
	}
}

func InitKafka() {
	testConn()

	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}

func CloseKafka() {
	err := global.KafkaProducer.Close()
	if err != nil {
		log.Fatalf("Fail to close kafka producer: %v", err)
	}
}
