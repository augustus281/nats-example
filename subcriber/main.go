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

	subject := "order.created"
	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		var order Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Println("❌ Error decoding order:", err)
			return
		}

		fmt.Printf("💳 Processing payment for Order ID: %s, Amount: %.2f %s\n", order.OrderID, order.Amount, order.Currency)

		confirmationSubject := "order.payment.success"
		confirmationMessage := fmt.Sprintf("✅ Payment successful for Order ID: %s", order.OrderID)
		nc.Publish(confirmationSubject, []byte(confirmationMessage))
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("📡 Listening for new orders...")
	select {}
}
