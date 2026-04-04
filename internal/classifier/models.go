package classifier

type ContentType string

// TODO: Classify into more: books, reddit posts, substrack posts, audio?, and more
const (
	TypeArticle     ContentType = "article"
	TypeVideo       ContentType = "video"
	TypeBook        ContentType = "book"
	TypeTwitterPost ContentType = "twitter post"
	TypeUnknown     ContentType = "unknown"
)
