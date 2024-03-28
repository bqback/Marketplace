package dto

type FeedOptions struct {
	Sort  *SortOptions
	Page  *PaginationOptions
	Price *PriceRange
}

type SortOptions struct {
	Type  int
	Order int
}

const (
	DateSort int = iota
	PriceSort
)

const (
	AscSort int = iota
	DescSort
)

type PaginationOptions struct {
	Page    *uint64
	PerPage *uint64
}

type PriceRange struct {
	Min *uint64
	Max *uint64
}

var SortTypeMap = map[string]int{
	"":      DateSort,
	"date":  DateSort,
	"price": PriceSort,
}

var SortOrderMap = map[string]int{
	"asc":  AscSort,
	"desc": DescSort,
}

var DefaultSort = map[int]int{
	DateSort:  DescSort,
	PriceSort: AscSort,
}
