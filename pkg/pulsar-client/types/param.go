package types

var (
	BrokerUrl        string
	Topic            string
	SubscriptionName string
	MessageNum       int64
	ConsumeTime      int64
	ProduceTime      int64
	Readness         bool

	Topics        []string
	TopicsPattern string

	SubscriptionType string

	ReadCompacted bool
)

type ConsumerMessageOption struct {
	BrokerUrl        string
	Topic            string
	Topics           []string
	SubscriptionName string
	ConsumeTime      int64
	Readness         bool
	TopicsPattern    string
	SubscriptionType string
	ReadCompacted    bool
}

type ProducerMessageOption struct {
	BrokerUrl        string
	Topic            string
	SubscriptionName string
	MessageNum       int64
	ProduceTime      int64
	Readness         bool
}
