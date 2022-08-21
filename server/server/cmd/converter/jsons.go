package converter

type JsonStruct interface {
	GetID() string
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	Is_active string `json:"is_active"`
}

func (s User) GetID() string {
	return s.ID
}

type Groupy struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Introduce  string `json:"introduce"`
	Max_people string `json:"max_people"`
	Group_key  string `json:"group_key"`
	Is_opened  string `json:"is_opened"`
}

func (s Groupy) GetID() string {
	return s.ID
}

type UserInGroupy struct {
	ID     string `json:"id"`
	User   string `json:"user"`
	Groupy string `json:"groupy"`
}

func (s UserInGroupy) GetID() string {
	return s.ID
}

type BoardItem struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Category   string `json:"category"`
	Title      string `json:"title"`
	Detail     string `json:"detail"`
	Applicants string `json:"applicants"`
	Is_opened  string `json:"is_opened"`
}

func (s BoardItem) GetID() string {
	return s.ID
}

type GroupyBoardItem struct {
	ID        string `json:"id"`
	Groupy    string `json:"groupy"`
	BoardItem string `json:"board_item"`
}

func (s GroupyBoardItem) GetID() string {
	return s.ID
}

type UserLikeBoardItem struct {
	ID        string `json:"id"`
	User      string `json:"user"`
	BoardItem string `json:"board_item"`
}

func (s UserLikeBoardItem) GetID() string {
	return s.ID
}

type HomeItem struct {
	ID        string `json:"id"`
	Detail    string `json:"detail"`
	Is_opened string `json:"is_opened"`
}

func (s HomeItem) GetID() string {
	return s.ID
}

type PostHomeItem struct {
	ID        string `json:"id"`
	User      string `json:"user"`
	Groupy    string `json:"groupy"`
	BoardItem string `json:"board_item"`
}

func (s PostHomeItem) GetID() string {
	return s.ID
}
