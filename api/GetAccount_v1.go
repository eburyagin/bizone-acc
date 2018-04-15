package api

const GET_ACCOUNT_V1 = "bz.GetAccount_v1"

type GetAccountReq_v1 struct {
	Meta Meta
	ID   string
}

type GetAccountResp_v1 struct {
	Meta   Meta
	Data   AccountFullData
	Status Status
}
