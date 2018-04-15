package services

import (
	"log"

	"github.com/eburyagin/bizone-acc/api"
	"github.com/eburyagin/bizone-acc/internal/cfg"
	nats "github.com/nats-io/go-nats"
)

func StartGetAccount_v1(config *cfg.Config) error {
	log.Println("Запускаю ", api.GET_ACCOUNT_V1, "...")

	nc, _ := nats.Connect(config.Bus.Urls)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	c.Subscribe(api.GET_ACCOUNT_V1, func(subj, reply string, req *api.GetAccountReq_v1) {
		log.Println("Запрос: ", req)
		var resp api.GetAccountResp_v1
		getAccount_v1(req, &resp)
		c.Publish(reply, "I can help!")
		log.Println("Ответ: ", resp)
	})

	log.Println(api.GET_ACCOUNT_V1, " стартовал.")
	return nil
}

func getAccount_v1(req *api.GetAccountReq_v1, resp *api.GetAccountResp_v1) error {
	resp.Meta = api.Meta{
		Ref: req.Meta.Ref,
	}
	resp.Data = api.AccountFullData{
		ID:     req.ID,
		Number: "88347899838734",
	}
	resp.Status = api.Status{
		Code: 0,
	}
	return nil
}
