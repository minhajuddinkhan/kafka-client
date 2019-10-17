package kafka

func (kc *client) Brokers() []string {
	return kc.brokers
}
