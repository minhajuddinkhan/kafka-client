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
	for _, x := range bg.Brokers {
		urls = append(urls, x.URL)
	}
	return urls
}
