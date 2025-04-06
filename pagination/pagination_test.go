package pagination_test

import (
	"testing"

	"github.com/JoshuaGanter/go-pagination-seven/pagination"
)

func TestGetPagination(t *testing.T) {
	basicTests := []struct {
		currentPage int
		totalPages  int
		expected    string
	}{
		{2, 5, "1 (2) 3 4 5"},
		{3, 7, "1 2 (3) 4 5 6 7"},
	}

	for _, testCase := range basicTests {
		actual := pagination.GetPagination(testCase.currentPage, testCase.totalPages)

		if actual != testCase.expected {
			t.Errorf("GetPagination(%d, %d) returned %q, expected %q", testCase.currentPage, testCase.totalPages, actual, testCase.expected)
		}
	}
}
