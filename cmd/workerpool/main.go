package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	id    int
	Title string
	Text  string
}

type IPool interface {
	Create()
	Handle(Message)
	Wait()
	Stats()
}

var maxMessages = 10
var messagesCounter = 0

func getMessages() []Message {
	messagesCount := rand.Intn(maxMessages)

	messages := make([]Message, 0, messagesCount)

	for range messagesCount {
		messagesCounter++
		messages = append(messages, Message{id: messagesCounter})
	}

	return messages
}

func processMessage(workerId int, message Message) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("Worker %d processed message %d\n", workerId, message.id)
}

func main() {
	var pool IPool

	pool = NewPool(processMessage)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
		}

		messages := getMessages()

		pool.Create()

		for _, message := range messages {
			pool.Handle(message)
		}

		pool.Wait()
	}

	pool.Stats()
}
