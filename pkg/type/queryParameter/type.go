package queryParameter

import (
	"assig3/pkg/type/pagination"
	"assig3/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
}
