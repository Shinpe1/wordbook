package request

type DeleteBookComp struct {
	UserId     int     `json:"userId" binding:"required"`
	BookId     int64   `json:"bookId" binding:"required"`
	ContentsId []int64 `json:"contentsId"`
}
