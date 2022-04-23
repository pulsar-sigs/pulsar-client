package producer

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/pulsar-sigs/pulsar-client/pkg/pulsar-client/types"
	"github.com/spf13/cobra"
)

func consumeMessage(url, topic, subscriptionName string) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		log.Fatalln("create.pfoducer.failed", err)
	}

	for {
		_, err := producer.Send(context.TODO(), &pulsar.ProducerMessage{
			Payload: []byte("hello"),
		})
		if err != nil {
			log.Println("producer.send.message.failed!", err)
		}
		time.Sleep(time.Second * 1)

	}
}

func NewProducerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "producer",
		Short:   "producer",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if types.BrokerUrl == "" {
				cmd.Help()
				return errors.New("brokerUrl is empty")
			}
			if types.Topic == "" {
				cmd.Help()
				return errors.New("topic is empty")
			}
			if types.SubscriptionName == "" {
				cmd.Help()
				return errors.New("subscriptionName is empty")
			}
			log.Println("broker:", types.BrokerUrl)
			log.Println("topic:", types.Topic)
			log.Println("subscriptionName:", types.SubscriptionName)

			consumeMessage(types.BrokerUrl, types.Topic, types.SubscriptionName)

			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&types.BrokerUrl, "broker", "", "pulsar broker url")
	cmd.PersistentFlags().StringVar(&types.Topic, "topic", "", "pulsar topic")
	cmd.PersistentFlags().StringVar(&types.SubscriptionName, "subscription-name", "", "pulsar consumer subscriptionName")

	return cmd
}
