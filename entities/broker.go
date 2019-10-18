package entities

//Broker Broker
type Broker struct {
	ID  uint
	URL string
}

type BrokerGroup struct {
	Brokers []Broker
}

func (bg BrokerGroup) URLs() []string {

	urls := make([]string, len(bg.Brokers))
	for i, x := range bg.Brokers {
		urls[i] = x.URL
	}
	return urls
}
