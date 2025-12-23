package domain

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (t Todo) Validate() error {
	if t.Title == "" {
		return ErrNoTitle
	}
	return nil
}
