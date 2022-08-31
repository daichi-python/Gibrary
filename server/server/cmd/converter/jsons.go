package converter

type JsonStruct interface {
	GetID() string
}

type User struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Birthday  string `json:"birthday" db:"birthday"`
	Gender    string `json:"gender" db:"gender"`
	Is_active string `json:"is_active" db:"is_active"`
}

func (s User) GetID() string {
	return s.ID
}

type Groupy struct {
	ID         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Introduce  string `json:"introduce" db:"introduce"`
	Max_people string `json:"max_people" db:"max_people"`
	Group_key  string `json:"group_key" db:"group_key"`
	Is_opened  string `json:"is_opened" db:"is_opened"`
}

func (s Groupy) GetID() string {
	return s.ID
}

type UserInGroupy struct {
	ID     string `json:"id" db:"id"`
	User   string `json:"user" db:"user"`
	Groupy string `json:"groupy" db:"groupy"`
}

func (s UserInGroupy) GetID() string {
	return s.ID
}

type BoardItem struct {
	ID         string `json:"id" db:"id"`
	Type       string `json:"type" db:"type"`
	Category   string `json:"category" db:"category"`
	Title      string `json:"title" db:"title"`
	Detail     string `json:"detail" db:"detail"`
	Applicants string `json:"applicants" db:"applicants"`
	Is_opened  string `json:"is_opened" db:"is_opened"`
}

func (s BoardItem) GetID() string {
	return s.ID
}

type GroupyBoardItem struct {
	ID        string `json:"id" db:"id"`
	Groupy    string `json:"groupy" db:"groupy"`
	BoardItem string `json:"board_item" db:"board_item"`
}

func (s GroupyBoardItem) GetID() string {
	return s.ID
}

type UserLikeBoardItem struct {
	ID        string `json:"id" db:"id"`
	User      string `json:"user" db:"user"`
	BoardItem string `json:"board_item" db:"boarditem"`
}

func (s UserLikeBoardItem) GetID() string {
	return s.ID
}

type HomeItem struct {
	ID        string `json:"id" db:"id"`
	Detail    string `json:"detail" db:"detail"`
	Is_opened string `json:"is_opened" db:"is_opened"`
}

func (s HomeItem) GetID() string {
	return s.ID
}

type PostHomeItem struct {
	ID        string `json:"id" db:"id"`
	User      string `json:"user" db:"user"`
	Groupy    string `json:"groupy" db:"groupy"`
	BoardItem string `json:"board_item" db:"board_item"`
}

func (s PostHomeItem) GetID() string {
	return s.ID
}
