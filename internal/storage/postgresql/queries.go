package postgresql

import (
	"marketplace/internal/pkg/dto"
	"strings"
)

var shortIDField = "id"

var userTable = "public.user"

// public.user fields
var (
	userIdField    = "public.user.id"
	userLoginField = "public.user.login"
	userHashField  = "public.user.password_hash"
)

var (
	allUserSelectFields = []string{userIdField, userLoginField, userHashField, userIsAdminField}
)

var movieUpdateReturnSuffix = "RETURNING " + strings.Join(allMovieSelectFields, ", ")

var (
	actorMovieFields = []string{"id_actor", "id_movie"}
)

var SortOptionsMap = map[int]string{
	dto.DateSort:    movieTitleField,
	dto.PriceSort:   movieRatingField,
	dto.ReleaseSort: movieReleaseField,
}

var SortOrderMap = map[int]string{
	dto.AscSort:  "ASC",
	dto.DescSort: "DESC",
}
