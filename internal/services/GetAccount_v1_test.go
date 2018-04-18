package services

import (
	"os"
	"testing"
	"time"

	"github.com/eburyagin/bizone-acc/api"
	"github.com/eburyagin/bizone-acc/internal/cfg"
	"github.com/eburyagin/bizone-acc/internal/ds"
	nats "github.com/nats-io/go-nats"
)

func TestGetAccount_v1(t *testing.T) {
	cf := "../../" + os.Getenv("BZ_TEST_CONFIG_FILE")
	config, _ := cfg.Load(cf)
	ds.NewConnect(config)
	StartGetAccount_v1(config)

	nc, _ := nats.Connect(config.Bus.Urls)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	req := &api.GetAccountReq_v1{
		ID: "123",
	}
	var resp api.GetAccountResp_v1
	c.Request(api.GET_ACCOUNT_V1, req, &resp, 10*time.Millisecond)

}
