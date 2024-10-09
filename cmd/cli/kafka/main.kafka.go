package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "kafka:9092"
	kafkaTopic = "user_topic_vip"
)

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,              // 10KB
		MaxBytes:       10e6,              // 10MB
		CommitInterval: time.Second,       // khoảng thời gian giữa các lần commit offset là 1s
		StartOffset:    kafka.FirstOffset, // consumer mới sẽ lắng nghe từ offset đầu tiên | kafka.LastOffset: comsumer mới chỉ đọc tin nhắn mới
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua ban chung khoan
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	// create kafka message
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	// write a message by producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
	})
}

// Consumer hóng mua ATC
func RegisterConsumerATC(id int) {
	// group consumer
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer(%d) hong phien ATC::\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) error: %v\n", id, err)
		}
		fmt.Printf("Consumer(%d) hong topic: %v, partition: %v, offset: %v, time: %d \n %s = %s \n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	// đăng ký 3 user để mua stock trong phiên ATC (1), (2) and (3)
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)

	r.Run(":8999")
}
