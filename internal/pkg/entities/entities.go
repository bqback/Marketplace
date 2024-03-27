package entities

import (
	"marketplace/internal/pkg/dto"
)

type Listing struct {
	ID uint64 `json:"id"`
	dto.NewListingInfo
}

type User struct {
	ID    uint64 `json:"id"`
	Login string `json:"login"`
}

type Role struct {
	RoleName string
	IsAdmin  bool
}
