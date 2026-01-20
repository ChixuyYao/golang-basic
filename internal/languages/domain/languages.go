package domain

import "time"

type Languages struct {
	Id        string    `json:"id"`
	ISO       string    `json:"iso"`
	Name      string    `json:"name"`
	AliasCn   string    `json:"alias_cn"`
	AliasEn   string    `json:"alias_en"`
	Script    string    `json:"script"`
	Direction string    `json:"direction"`
	Family    string    `json:"family"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
