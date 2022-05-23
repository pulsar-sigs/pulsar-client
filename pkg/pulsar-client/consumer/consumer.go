package consumer

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/pulsar-sigs/pulsar-client/pkg/pulsar-client/types"
	"github.com/spf13/cobra"
)

func consumeMessage(opt *types.ConsumerMessageOption) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               opt.BrokerUrl,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	log.Println(opt.Topics)

	pulsarconsumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            opt.Topic,
		SubscriptionName: opt.SubscriptionName,
		Type:             pulsar.Shared,
		TopicsPattern:    opt.TopicsPattern,
		Topics:           opt.Topics,
	})
	if err != nil {
		log.Fatalf("Could not create Pulsar consumer: %v", err)
	}
	defer pulsarconsumer.Close()

	if opt.Readness {
		go types.RunReadnessAPI()
	}

	for {
		msg, err := pulsarconsumer.Receive(context.TODO())
		if err != nil {
			log.Println("receive message failed!", err)
			continue
		}
		if opt.ConsumeTime > 0 {
			time.Sleep(time.Millisecond * time.Duration(opt.ConsumeTime))
		}
		log.Printf("consume message: topic is %s , topic key is :%s and payload is :%s \n", msg.Topic(), msg.Key(), msg.Payload())
		pulsarconsumer.Ack(msg)
	}
}

func NewConsumerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "consumer",
		Short:   "consumer",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if types.BrokerUrl == "" {
				cmd.Help()
				return errors.New("brokerUrl is empty")
			}
			if types.Topic == "" && len(types.Topics) == 0 && types.TopicsPattern == "" {
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

			consumeMessage(&types.ConsumerMessageOption{
				BrokerUrl:        types.BrokerUrl,
				Topic:            types.Topic,
				SubscriptionName: types.SubscriptionName,
				ConsumeTime:      types.ConsumeTime,
				Readness:         types.Readness,
				TopicsPattern:    types.TopicsPattern,
				Topics:           types.Topics,
			})

			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&types.BrokerUrl, "broker", "", "pulsar broker url")
	cmd.PersistentFlags().StringVar(&types.Topic, "topic", "", "pulsar topic")
	cmd.PersistentFlags().StringVar(&types.TopicsPattern, "topic-pattern", "", "pulsar topic parttern")
	cmd.PersistentFlags().StringVar(&types.SubscriptionName, "subscription-name", "", "pulsar consumer subscriptionName")
	cmd.PersistentFlags().Int64Var(&types.ConsumeTime, "consume-time", 0, "consume time (millisecond) for one message, 0 by default")
	cmd.PersistentFlags().BoolVar(&types.Readness, "readness", false, "start readness api endpoint, true by default.")
	cmd.PersistentFlags().StringArrayVar(&types.Topics, "topics", []string{}, "topics")

	return cmd
}
