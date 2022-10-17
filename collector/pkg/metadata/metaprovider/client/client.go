package metadataclient

import "k8s.io/client-go/tools/cache"

type Client struct {
	// tls wrap
}

func (c *Client) ListAndWatch(
	//req api.MetaDataRequest,
	handler cache.ResourceEventHandler,
) {
	


}
