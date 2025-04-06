package pagination_test

import (
	"testing"

	"github.com/JoshuaGanter/go-pagination-seven/pagination"
)

func TestGetPagination(t *testing.T) {
	t.Run("total pages less than or equal to 7", func(t *testing.T) {
		basicTests := []struct {
			currentPage int
			totalPages  int
			expected    string
		}{
			{1, 1, "(1)"},
			{2, 5, "1 (2) 3 4 5"},
			{1, 7, "(1) 2 3 4 5 6 7"},
			{3, 7, "1 2 (3) 4 5 6 7"},
			{7, 7, "1 2 3 4 5 6 (7)"},
		}

		for _, testCase := range basicTests {
			actual := pagination.GetPagination(testCase.currentPage, testCase.totalPages)

			if actual != testCase.expected {
				t.Errorf("GetPagination(%d, %d) returned %q, expected %q", testCase.currentPage, testCase.totalPages, actual, testCase.expected)
			}
		}
	})

	t.Run("with ellipsis", func(t *testing.T) {
		testsWithEllipsis := []struct {
			currentPage int
			totalPages  int
			expected    string
		}{
			{1, 9, "(1) 2 3 4 5 ... 9"},
			{4, 9, "1 2 3 (4) 5 ... 9"},
			{5, 9, "1 ... 4 (5) 6 ... 9"},
			{6, 9, "1 ... 5 (6) 7 8 9"},
			{9, 9, "1 ... 5 6 7 8 (9)"},
			{42, 100, "1 ... 41 (42) 43 ... 100"},
		}

		for _, testCase := range testsWithEllipsis {
			actual := pagination.GetPagination(testCase.currentPage, testCase.totalPages)

			if actual != testCase.expected {
				t.Errorf("GetPagination(%d, %d) returned %q, expected %q", testCase.currentPage, testCase.totalPages, actual, testCase.expected)
			}
		}
	})
}
