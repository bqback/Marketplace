package postgresql

import (
	"marketplace/internal/pkg/dto"
)

var shortIDField = "id"

var returnIDSuffix = "RETURNING " + shortIDField

var userTable = "public.user"

// public.user fields
var (
	userIdField    = "public.user.id"
	userLoginField = "public.user.login"
	userHashField  = "public.user.password_hash"
)

var (
	userShortLoginField = "login"
	userShortHashField  = "password_hash"
)

var (
	allUserSelectFields = []string{userIdField, userLoginField, userHashField}
	allUserInsertFields = []string{userShortLoginField, userShortHashField}
)

var userListingTable = "public.user_listing"

var (
	userListingFields = []string{"id_user", "id_listing"}
)

var listingTable = "public.listing"

// public.listing fields
var (
	listingIdField          = "public.listing.id"
	listingTitleField       = "public.listing.title"
	listingDescriptionField = "public.listing.description"
	listingImageLinkField   = "public.listing.image_link"
	listingPriceField       = "public.listing.price"
	listingDateCreatedField = "public.listing.date_created"
)

var (
	listingShortTitleField       = "title"
	listingShortDescriptionField = "description"
	listingShortImageLinkField   = "image_link"
	listingShortPriceField       = "price"
	listingShortDateCreatedField = "date_created"
)

var (
	allListingInsertFields = []string{listingShortTitleField, listingShortDescriptionField, listingShortImageLinkField, listingShortPriceField, listingShortDateCreatedField}
)

var SortOptionsMap = map[int]string{
	dto.DateSort:  listingDateCreatedField,
	dto.PriceSort: listingPriceField,
}

var SortOrderMap = map[int]string{
	dto.AscSort:  "ASC",
	dto.DescSort: "DESC",
}
