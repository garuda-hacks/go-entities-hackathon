package response

type Pagination struct {
	TotalRecords int64 `json:"total_records"`
	TotalPage    int   `json:"total_page"`
	LimitPage    int   `json:"limit_page"`
	CurrentPage  int   `json:"current_page"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.LimitPage <= 0 {
		p.LimitPage = 10
	}
	return p.LimitPage
}

func (p *Pagination) GetPage() int {
	if p.CurrentPage <= 0 {
		p.CurrentPage = 1
	}
	return p.CurrentPage
}
