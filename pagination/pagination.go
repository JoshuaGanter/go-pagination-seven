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
	} else {
		return "1 ... 41 (42) 43 ... 100"
	}

	return strings.Join(pagination, " ")
}
