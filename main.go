package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/JoshuaGanter/go-pagination-seven/pagination"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: pagination <currentPage> <totalPages>")
		return
	}

	currentPage, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	totalPages, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	paginationString, err := pagination.GetPagination(currentPage, totalPages)
	if err != nil {
		panic(err)
	}

	fmt.Println(paginationString)
}
