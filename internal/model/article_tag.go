package model

type ArticelTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticelTag) TableName() string {
	return "blog_article_tag"
}
