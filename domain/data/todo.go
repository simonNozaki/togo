package data

import "time"

type State int

const (
	unprocessed State = iota
	inprogress        = iota
	done              = iota
	gone              = iota
)

func (s State) String() string {
	r := ""
	switch s {
	case unprocessed:
		r = "unprocessed"
	case inprogress:
		r = "in progress"
	case done:
		r = "done"
	case gone:
		r = "gone"
	}
	return r
}

func GetState(s string) State {
	r := unprocessed
	switch s {
	case "unprocessed":
		r = unprocessed
	case "in progress":
		r = inprogress
	case "done":
		r = done
	case "gone":
		r = gone
	}
	return r
}

type Todo struct {
	Id          string `json:"id"`
	UserId      string `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       `json:"state"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updateAt"`
}
