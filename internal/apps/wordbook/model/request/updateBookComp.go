package request

type UpdateBookComp struct {
	UserId   int              `json:"userId" binding:"required"`
	BookId   int64            `json:"bookId" binding:"required"`
	Title    string           `json:"title"`
	Contents []content4update `json:"contents"`
}

type content4update struct {
	ContentsId int64  `json:"contentsId" binding:"required"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}
