package dto

type ExportCSVRequest struct {
	Filter Filter
	Url    Url
}
type Filter struct {
	Status    string `form:"status"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}
type Url struct {
	ClientUrl string
	PathUrl   string
}

type ErrorResponse struct {
	Error string `json:"error"`
}
type MessageResponse struct {
	Message string `json:"message"`
}
