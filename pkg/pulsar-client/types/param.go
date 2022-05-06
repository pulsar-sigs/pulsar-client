package types

var (
	BrokerUrl        string
	Topic            string
	SubscriptionName string
	MessageNum       int64
	ConsumeTime      int64
	ProduceTime      int64
)

type ConsumerMessageOption struct {
	BrokerUrl        string
	Topic            string
	SubscriptionName string
	ConsumeTime      int64
}

type ProducerMessageOption struct {
	BrokerUrl        string
	Topic            string
	SubscriptionName string
	MessageNum       int64
	ProduceTime      int64
}
