package pagination

import (
	"fmt"
	"strconv"
	"strings"
)

func GetPagination(currentPage int, totalPages int) string {
	pagination := make([]string, 0, 7)

	if totalPages <= 7 {
		for page := 1; page <= totalPages; page++ {
			if page == currentPage {
				pagination = append(pagination, fmt.Sprintf("(%d)", page))
				continue
			}
			pagination = append(pagination, strconv.Itoa(page))
		}
	} else if currentPage <= 4 {
		for page := 1; page <= 5; page++ {
			if page == currentPage {
				pagination = append(pagination, fmt.Sprintf("(%d)", page))
				continue
			}
			pagination = append(pagination, strconv.Itoa(page))
		}
		pagination = append(pagination, "...", strconv.Itoa(totalPages))
	} else {
		return fmt.Sprintf("1 ... %d (%d) %d ... %d", currentPage-1, currentPage, currentPage+1, totalPages)
	}

	return strings.Join(pagination, " ")
}
