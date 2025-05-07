package docker

import (
	"github.com/docker/docker/client"
	"net/http"
	"server-aggregation/config"
	"sync"
)

var cli *client.Client

type InstanceInfo struct {
	mu sync.Mutex
}

var instanceInfo *InstanceInfo

func Init() {
	var err error
	var httpClient *http.Client
	cli, err = client.NewClientWithOpts(client.WithHost(config.GetString("docker.instance.url")), client.WithVersion(config.GetString("docker.instance.version")), client.WithHTTPClient(httpClient))
	if err != nil {
		panic(err)
	}
	instanceInfo = &InstanceInfo{}
}
