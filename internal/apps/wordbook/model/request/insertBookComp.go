package request

type InsertBookComp struct {
	UserId   int              `json:"userId" binding:"required"`
	BookId   int64            `json:"bookId" binding:"required"`
	Contents []content4insert `json:"contents"`
}

type content4insert struct {
	Head string `json:"head"`
	Tail string `json:"tail"`
}
