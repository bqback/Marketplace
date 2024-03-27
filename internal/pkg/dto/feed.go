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

var ParamMap = map[string]int{
	"asc":   AscSort,
	"desc":  DescSort,
	"date":  DateSort,
	"price": PriceSort,
}
