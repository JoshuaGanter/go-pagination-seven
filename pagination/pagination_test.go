package pagination_test

import (
	"testing"

	"github.com/JoshuaGanter/go-pagination-seven/pagination"
	"github.com/stretchr/testify/assert"
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
			actual, err := pagination.GetPagination(testCase.currentPage, testCase.totalPages)

			assert.NoError(t, err)
			assert.Equal(t, actual, testCase.expected)
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
			{4, 100, "1 2 3 (4) 5 ... 100"},
			{42, 100, "1 ... 41 (42) 43 ... 100"},
			{97, 100, "1 ... 96 (97) 98 99 100"},
		}

		for _, testCase := range testsWithEllipsis {
			actual, err := pagination.GetPagination(testCase.currentPage, testCase.totalPages)

			assert.NoError(t, err)
			assert.Equal(t, actual, testCase.expected)
		}
	})

	t.Run("errors out on invalid input", func(t *testing.T) {
		testsWithErrors := []struct {
			currentPage   int
			totalPages    int
			expectedError string
		}{
			{-1, -1, "parameter \"totalPages\" must be greater than 0, actual: -1"},
			{1, -1, "parameter \"totalPages\" must be greater than 0, actual: -1"},
			{1, 0, "parameter \"totalPages\" must be greater than 0, actual: 0"},
			{-1, 9, "parameter \"currentPage\" must be in range [1..9], actual: -1"},
			{0, 9, "parameter \"currentPage\" must be in range [1..9], actual: 0"},
			{6, 5, "parameter \"currentPage\" must be in range [1..5], actual: 6"},
			{42, 10, "parameter \"currentPage\" must be in range [1..10], actual: 42"},
		}
		for _, testCase := range testsWithErrors {
			_, actualError := pagination.GetPagination(testCase.currentPage, testCase.totalPages)

			assert.EqualError(t, actualError, testCase.expectedError)
		}
	})
}
