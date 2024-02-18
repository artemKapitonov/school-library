package entity

type Book struct {
	Students []Student `json:"students,omitempty"`
	ID       uint64    `json:"id,omitempty"`
	Author   string    `json:"author,omitempty"`
	Title    string    `json:"title,omitempty"`
	Year     uint32    `json:"year,omitempty"`
	Count    uint32    `json:"count,omitempty"`
}

type Student struct {
	Class string `json:"class,omitempty"`
	Books []Book `json:"books,omitempty"`
	ID    uint64 `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}
