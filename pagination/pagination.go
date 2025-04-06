package pagination

import (
	"fmt"
	"strconv"
	"strings"
)

func getPaginationRange(fromPage int, toPage int, currentPage int) []string {
	result := make([]string, 0, (toPage-fromPage)+1)
	for page := fromPage; page <= toPage; page++ {
		if page == currentPage {
			result = append(result, fmt.Sprintf("(%d)", page))
			continue
		}
		result = append(result, strconv.Itoa(page))
	}
	return result
}

func GetPagination(currentPage int, totalPages int) string {

	if totalPages <= 7 {
		return strings.Join(getPaginationRange(1, totalPages, currentPage), " ")
	}

	if currentPage <= 4 {
		pagination := make([]string, 0, 7)
		pagination = append(pagination, getPaginationRange(1, 5, currentPage)...)
		pagination = append(pagination, "...", strconv.Itoa(totalPages))
		return strings.Join(pagination, " ")
	}

	if totalPages-currentPage <= 3 {
		pagination := make([]string, 0, 7)
		pagination = append(pagination, "1", "...")
		pagination = append(pagination, getPaginationRange(totalPages-4, totalPages, currentPage)...)
		return strings.Join(pagination, " ")
	}

	return fmt.Sprintf("1 ... %d (%d) %d ... %d", currentPage-1, currentPage, currentPage+1, totalPages)
}
