package dto

type FindParams struct {
	SearchTerm     string
	CategoryIDList []int64
	Limit          int64
	Offset         int64
}
