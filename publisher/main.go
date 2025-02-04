package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

type Order struct {
	OrderID   string  `json:"order_id"`
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Timestamp string  `json:"timestamp"`
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	order := Order{
		OrderID:   "ORD12345",
		UserID:    "USR5678",
		Amount:    49.99,
		Currency:  "USD",
		Timestamp: "2025-02-04T10:00:00Z",
	}

	orderData, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}

	subject := "order.created"
	err = nc.Publish(subject, orderData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Order published:", order.OrderID)

	nc.Flush()
}
