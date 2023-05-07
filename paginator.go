package paginator

import "errors"

// currentPage this is the default page starting number
const minPage int = 1

type Paginator[T any] struct {
	results     [][]T
	currentPage int
	maxPerPage  int
	nbResults   int
	maxNbPages  int
}

// NewPaginator creates our generic paginator and accepts a slice of any type and the max items per page to paginate
func NewPaginator[T any](results []T, maxItemsPerPage int) *Paginator[T] {
	p := &Paginator[T]{
		maxPerPage:  maxItemsPerPage,
		currentPage: minPage,
		nbResults:   len(results),
	}

	p.results = chunk(results, p.nbResults, p.maxPerPage)
	p.maxNbPages = len(p.results)

	return p
}

func chunk[T any](results []T, nbResults int, maxPerPage int) [][]T {
	var chunks [][]T
	for i := 0; i < nbResults; i += maxPerPage {
		end := i + maxPerPage

		if end > nbResults {
			end = nbResults
		}

		chunks = append(chunks, results[i:end])
	}

	return chunks
}

// GetCurrentPageNumber returns the current page number from the paginator
func (p *Paginator[T]) GetCurrentPageNumber() int {
	return p.currentPage
}

// GetMaxPerPageNumber returns the max items per page from the paginator
func (p *Paginator[T]) GetMaxPerPageNumber() int {
	return p.maxPerPage
}

// GetTotalNumberOfResults returns the total number of results from the paginator
func (p *Paginator[T]) GetTotalNumberOfResults() int {
	return p.nbResults
}

// GetMaxNumberOfPages returns the maximum number of pages from the paginator
func (p *Paginator[T]) GetMaxNumberOfPages() int {
	return p.maxNbPages
}

// GetCurrentPageResults returns the current page results from the paginator
func (p *Paginator[T]) GetCurrentPageResults() []T {
	return p.results[p.currentPage-1]
}

// HasNextPage returns true|false if we have another page of results to paginate
func (p *Paginator[T]) HasNextPage() bool {
	return p.currentPage+1 <= p.maxNbPages
}

// GetNextPageResults gets the next page of results
// If no more pages are available an error is returned
func (p *Paginator[T]) GetNextPageResults() ([]T, error) {
	if p.HasNextPage() == true {
		p.currentPage += 1
		return p.GetCurrentPageResults(), nil
	}

	return nil, errors.New("no more pages available")
}

// HasPreviousPage returns true|false if we have a prior page of results to paginate
func (p *Paginator[T]) HasPreviousPage() bool {
	return p.maxNbPages > 0 && p.currentPage-1 > minPage
}

// GetPreviousPageResults gets the prior page of results
// If no prior pages are available an error is returned
func (p *Paginator[T]) GetPreviousPageResults() ([]T, error) {
	if p.HasPreviousPage() == true {
		p.currentPage -= 1
		return p.GetCurrentPageResults(), nil
	}

	return nil, errors.New("no prior pages available")
}

// SetCurrentPage jump to a pagination index of your choice
// If the page index provided is greater than max number of pages or the min pages allowed an error is returned
func (p *Paginator[T]) SetCurrentPage(page int) error {
	if page > p.maxNbPages {
		return errors.New("the page provided cannot be greater than the maximum number of pages")
	}

	if page < minPage {
		return errors.New("the page provided cannot be less than 1")
	}

	p.currentPage = page

	return nil
}
