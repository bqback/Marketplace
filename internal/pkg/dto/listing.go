package dto

type NewListingInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
	Price       uint   `json:"price"`
}

type FeedListingInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
	Price       uint   `json:"price"`
	Author      string `json:"author"`
	OwnedByUser bool   `json:"owned_by_user"`
}
