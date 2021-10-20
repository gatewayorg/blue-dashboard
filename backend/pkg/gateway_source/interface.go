package gateway_source

type WatchAddrs interface {
	Watch() (addrsChan <-chan []string)
}
