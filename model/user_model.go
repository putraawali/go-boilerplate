package model

type UserModel struct {
	Id            int64 `json:"id"`
	TransporterId int64 `json:"transporter_id"`
	ShipperId     int64 `json:"shipper_id"`
}
