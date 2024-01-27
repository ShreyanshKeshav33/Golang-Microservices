package repositories

type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:description`
}
type CreateResponse struct {
	Id    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}
