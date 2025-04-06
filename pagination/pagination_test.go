package pagination_test

import (
	"testing"

	"github.com/JoshuaGanter/go-pagination-seven/pagination"
)

func TestGetPagination(t *testing.T) {
	currentPage := 2
	totalPages := 5
	expected := "1 (2) 3 4 5"

	actual := pagination.GetPagination(currentPage, totalPages)

	if actual != expected {
		t.Errorf("GetPagination(%d, %d) returned %q, expected %q", currentPage, totalPages, actual, expected)
	}
}
