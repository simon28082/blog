package document

type MarkdownParse interface {
	Paginate(paging uint)
	Detail(filename string)
}

type markdown struct {
	path string
}

func NewMarkdown(path string) *markdown  {
	return &markdown{
		path:path,
	}
}