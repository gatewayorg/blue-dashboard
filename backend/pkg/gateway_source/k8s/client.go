package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"time"
)

type Client struct {
	client          *kubernetes.Clientset
	namespace, name string
}

func NewProxy(proxy string, namespace string, name string) (*Client, error) {
	client, err := kubernetes.NewForConfig(&rest.Config{
		Host: proxy,
	})

	if err != nil {
		return nil, err
	}
	return &Client{
		client:    client,
		name:      name,
		namespace: namespace,
	}, nil
}

func NewClient(namespace string, name string) (*Client, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, err
	}

	return &Client{
		client:    client,
		name:      name,
		namespace: namespace,
	}, nil
}

func (c *Client) Watch() (addrsChan <-chan []string) {
	watchChan := make(chan []string)

	go func() {
		tc := time.NewTicker(time.Second * 5)
		var (
			watchInterface watch.Interface
			err            error
		)

		for range tc.C {
			watchInterface, err = c.client.CoreV1().Endpoints(c.namespace).Watch(metav1.ListOptions{})
			if err != nil {
				continue
			}
			for e := range watchInterface.ResultChan() {
				endpoints, ok := e.Object.(*v1.Endpoints)
				if ok && endpoints.Name == c.name {
					addrs := c.getAddrs(endpoints)
					if len(addrs) != 0 {
						watchChan <- addrs
					}
				}
			}
		}
	}()
	return watchChan
}

func (c *Client) getAddrs(endpoints *v1.Endpoints) (ips []string) {
	if len(endpoints.Subsets) == 0 {
		return
	}
	ips = make([]string, 0, len(endpoints.Subsets[0].Addresses))

	for _, addrs := range endpoints.Subsets[0].Addresses {
		ips = append(ips, addrs.IP)
	}
	return ips
}
