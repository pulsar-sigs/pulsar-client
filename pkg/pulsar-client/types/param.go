package types

var (
	BrokerUrl        string
	Topic            string
	SubscriptionName string
	MessageNum       int64
)

type ProducerMessageOption struct {
	BrokerUrl        string
	Topic            string
	SubscriptionName string
	MessageNum       int64
}
