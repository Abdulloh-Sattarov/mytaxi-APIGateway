package models

type OrderRes struct {
	Id        string  `json:"Id"`
	Cost      float64 `json:"Cost"`
	Status    string  `json:"Status"`
	Driver    Driver  `json:"Driver"`
	Client    Client  `json:"Client"`
	CreatedAt string  `json:"CreatedAt"`
	UpdatedAt string  `json:"UpdatedAt"`
}

type OrderReq struct {
	Id       string  `json:"Id"`
	Cost     float64 `json:"Cost"`
	Status   string  `json:"Status"`
	DriverId string  `json:"DriverId"`
	ClientId string  `json:"ClientId"`
}

type ListOrders struct {
	Orders []OrderRes `json:"Orders"`
	Count  int64      `json:"Count"`
}

type UpdateOrder struct {
	Status string `json:"Status"`
}
