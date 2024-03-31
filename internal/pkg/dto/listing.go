package dto

import "time"

type NewListingInfo struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageLink   string    `json:"image_link" db:"date_created"`
	Price       uint      `json:"price"`
	DateCreated time.Time `json:"date_created" db:"date_created"`
}

type FeedListingInfo struct {
	ID          uint64 `json:"id"`
	Author      string `json:"author"`
	OwnedByUser bool   `json:"owned_by_user"`
	NewListingInfo
}
