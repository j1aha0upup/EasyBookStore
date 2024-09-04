package model

type Pages struct {
	EveryPageRecordCount int
	PreviousPage         int
	NextPage             int
	CurrentPage          int
	TotalRecords         int
	TotalPages           int
	Min                  float64
	Max                  float64
}

func (page *Pages) GetPages() {
	// every page record count setting for six
	page.TotalPages = page.TotalRecords / page.EveryPageRecordCount
	if page.TotalRecords%page.EveryPageRecordCount != 0 {
		page.TotalPages++
	}
	if page.CurrentPage != 0 {
		page.PreviousPage = page.CurrentPage
	}
	if page.CurrentPage+1 != page.TotalPages {
		page.NextPage = page.CurrentPage + 2
	}
	page.CurrentPage++
}
