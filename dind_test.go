package dind

import (
	"fmt"
	"testing"

	"github.com/ory/dockertest"
	"github.com/streadway/amqp"
	"gotest.tools/assert"
)

func TestAll(t *testing.T) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatal(err)
	}

	res, err := pool.Run("rabbitmq", "management", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Purge(res)

	url := fmt.Sprintf("amqp://guest:guest@localhost:%s/", res.GetPort("5672/tcp"))

	var (
		conn *amqp.Connection
		ch   *amqp.Channel
		q    amqp.Queue
	)

	if err := pool.Retry(func() error {
		var err error
		conn, err = amqp.Dial(url)
		if err != nil {
			return err
		}

		ch, err = conn.Channel()
		if err != nil {
			return err
		}

		q, err = ch.QueueDeclare("hello", false, false, false, false, nil)
		return err
	}); err != nil {
		t.Fatal(err)
	}

	defer conn.Close()
	defer ch.Close()

	if err := ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello, World!"),
	}); err != nil {
		t.Fatal(err)
	}

	msgs, err := ch.Consume("", q.Name, true, false, false, false, nil)
	if err != nil {
		t.Fatal(err)
	}

	msg := <-msgs
	assert.Equal(t, string(msg.Body), "Hello, World!")
}
