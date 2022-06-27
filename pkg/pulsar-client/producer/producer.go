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

func produceMessage(opt *types.ProducerMessageOption) {

	var auth pulsar.Authentication
	if opt.AuthType != "" && opt.AuthParams != "" {
		autht, err := pulsar.NewAuthentication(opt.AuthType, opt.AuthParams)
		if err != nil {
			log.Fatalf("New auth failed: %v", err)
		}
		auth = autht
	}

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               opt.BrokerUrl,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
		Authentication:    auth,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:                   opt.Topic,
		DisableBatching:         false,
		BatchingMaxPublishDelay: time.Second,
	})
	if err != nil {
		log.Fatalln("create.pfoducer.failed", err)
	}
	if opt.Readness {
		go types.RunReadnessAPI()
	}

	log.Println("will produce message total ", opt.MessageNum)
	for i := 0; i < int(opt.MessageNum); i++ {
		if opt.ProduceTime > 0 {
			time.Sleep(time.Millisecond * time.Duration(opt.ProduceTime))
		}
		producer.SendAsync(context.TODO(), &pulsar.ProducerMessage{
			Payload: []byte(opt.Message),
			Key: types.MessageKey,
		}, func(mid pulsar.MessageID, msg *pulsar.ProducerMessage, e error) {
			if e != nil {
				log.Println("producer.send.message.failed!", e)
			} else {
				log.Printf("producer.send.message.success! %s %d-%d-%d", opt.Topic, mid.LedgerID(), mid.EntryID(), mid.PartitionIdx())
			}
		})
	}
	log.Println("exit after 3 second")
	time.Sleep(time.Second * 3)
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
			log.Println("broker:", types.BrokerUrl)
			log.Println("topic:", types.Topic)
			log.Println("subscriptionName:", types.SubscriptionName)

			produceMessage(&types.ProducerMessageOption{
				Topic:       types.Topic,
				BrokerUrl:   types.BrokerUrl,
				MessageNum:  types.MessageNum,
				ProduceTime: types.ProduceTime,
				Readness:    types.Readness,
				AuthType:    types.AuthType,
				AuthParams:  types.AuthParams,
				Message:     types.Message,
				MessageKey:  types.MessageKey,
			})

			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&types.BrokerUrl, "broker", "", "pulsar broker url")
	cmd.PersistentFlags().StringVar(&types.AuthType, "auth-type", "", "auth type")
	cmd.PersistentFlags().StringVar(&types.AuthParams, "auth-params", "", "auth-params")
	cmd.PersistentFlags().StringVar(&types.Topic, "topic", "", "pulsar topic")
	cmd.PersistentFlags().Int64Var(&types.MessageNum, "message-num", 10000, "produce message num")
	cmd.PersistentFlags().StringVar(&types.Message, "message", "hello", "produce message")
	cmd.PersistentFlags().StringVar(&types.MessageKey, "message key", "", "message topic key")
	cmd.PersistentFlags().Int64Var(&types.ProduceTime, "produce-time", 0, "produce time for one message,0(millisecond) by default.")
	cmd.PersistentFlags().BoolVar(&types.Readness, "readness", false, "start readness api endpoint, true by default.")

	return cmd
}
