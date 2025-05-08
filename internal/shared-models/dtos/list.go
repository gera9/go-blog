package shareddtos

type ListResponse[T any] struct {
	Total int `json:"total"`
	Items []T `json:"items"`
}
