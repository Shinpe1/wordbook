package request

type NewBook struct {
	UserId   int       `json:"userId" binding:"required"`
	Title    string    `json:"title"`
	Contents []content `json:"contents"`
}

type content struct {
	Head string `json:"head"`
	Tail string `json:"tail"`
}
