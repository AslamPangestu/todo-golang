package dtos

type pagination struct {
	Total       int  `json:"total"`
	PerPage     int  `json:"per_page"`
	CurrentPage int  `json:"current_page"`
	LastPage    int  `json:"last_page"`
	HasNext     bool `json:"has_next"`
	HasPrev     bool `json:"has_prev"`
}

type PaginationResponse struct {
	Pagination pagination  `json:"pagination"`
	Data       interface{} `json:"data"`
}

func PaginationResponseAdapter(page, pageSize, total int, data interface{}) PaginationResponse {
	if page == 0 {
		page = 1
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	var lastPage, remaining int
	remaining = total % pageSize
	if remaining != 0 {
		lastPage = total/pageSize + 1
	} else {
		lastPage = total / pageSize
	}

	return PaginationResponse{
		Data: data,
		Pagination: pagination{
			Total:       total,
			PerPage:     pageSize,
			CurrentPage: page,
			LastPage:    lastPage,
			HasPrev:     page > 1,
			HasNext:     page < lastPage && page >= 1,
		},
	}
}
