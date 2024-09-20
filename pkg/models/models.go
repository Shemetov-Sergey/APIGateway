package models

type CreateCommentRequestBody struct {
	NewsId   uint64 `json:"news_id"`
	ParentId uint64 `json:"parent_id"`
	Text     string `json:"text"`
	UserId   uint64 `json:"user_id"`
	Censored bool
}
