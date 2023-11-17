package domain

type SubmitOrderDelayResp struct {
	Message         string `json:"message"`
	NewDeliveryTime int    `json:"new_delivery_time"`
}
