package dto

import "GastroReserve/internal/entities"

type ReserveInputDTO struct {
	Custumer entities.Custumer `json:"custumer"`
	TableId  string            `json:"tableId"`
	Data     string            `json:"data"`
}

type ReserveOutputDTO struct {
	Id       string
	Custumer entities.Custumer
	TableId  string
	Data     string
}
