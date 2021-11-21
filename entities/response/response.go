package response

type Response struct {
	Meta       Meta        `json:"meta"`
	Pagination Pagination  `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
