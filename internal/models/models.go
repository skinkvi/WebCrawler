package model

type Page struct {
	Text string
}

type Link struct {
	URL    string
	Text   string
	PageID int
}

func NewPage(text string) *Page {
	return &Page{Text: text}
}

type Result struct {
	Title string
	URL   string
	Text  *string
}
