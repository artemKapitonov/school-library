package entity

type Book struct {
	ID       uint64    `json:"id,omitempty"`
	Author   string    `json:"author,omitempty"`
	Title    string    `json:"title,omitempty"`
	Year     uint32    `json:"year,omitempty"`
	Count    uint32    `json:"count,omitempty"`
	Students []Student `json:"students,omitempty"`
}

type Student struct {
	ID    uint64 `json:"id,omitempty"`
	Class string `json:"class,omitempty"`
	Name  string `json:"name,omitempty"`
	Books []Book `json:"books,omitempty"`
}
