# Slice Paginator
Takes any slice and generates a paginator

[![Go Reference](https://pkg.go.dev/badge/github.com/matt9mg/go-slice-paginator.svg)](https://pkg.go.dev/github.com/matt9mg/go-slice-paginator)


### Installation
```
go get github.com/matt9mg/go-slice-paginator
```

### Examples
```go
items := []string{"item1", "item2", "item3", "item4", "item5"}
maxItemsPerPage := 2
p := paginator.NewPaginator(items, maxItemsPerPage)
currentPage := p.GetCurrentPageResults()


items := []int{1, 2, 3, 4, 5}
maxItemsPerPage := 2
p := paginator.NewPaginator(items, maxItemsPerPage)
currentPage := p.GetCurrentPageResults()

type Items struct{
	ID int
}

items := []Items{
	{
		ID: 1,
    },
    {
        ID: 2,
    },
    {
        ID: 3,
    }
}
maxItemsPerPage := 2
p := paginator.NewPaginator(items, maxItemsPerPage)
currentPage := p.GetCurrentPageResults()
```

Other Helpful methods

```go
p.GetCurrentPageNumber()
p.GetMaxPerPageNumber()
p.GetTotalNumberOfResults()
p.GetMaxNumberOfPages()
p.HasNextPage()
p.GetNextPageResults()
p.HasPreviousPage()
p.GetPreviousPageResults()
p.SetCurrentPage()
```

### LICENSE
This project is licensed under the MIT License - see the LICENSE file for details

### Disclaimer
We take no legal responsibility for anything this code is used for.