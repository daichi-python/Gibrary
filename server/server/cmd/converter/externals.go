package converter

type UserInGroupyExternal struct {
	ID     string
	User   User
	Groupy Groupy
}

func (s UserInGroupyExternal) GetID() string {
	return s.ID
}

type Type struct {
	ID   string
	Type string
}

func (s Type) GetID() string {
	return s.ID
}

type Category struct {
	ID       string
	Category string
}

func (s Category) GetID() string {
	return s.ID
}

type GroupyBoardItemExternal struct {
	ID        string
	Groupy    Groupy
	BoardItem BoardItem
}

func (s GroupyBoardItemExternal) GetID() string {
	return s.ID
}

type UserLikeBoardItemExternal struct {
	ID        string
	User      User
	BoardItem BoardItem
}

func (s UserLikeBoardItemExternal) GetID() string {
	return s.ID
}

type PostHomeItemExternal struct {
	ID       string
	User     User
	Groupy   Groupy
	HomeItem HomeItem
}

func (s PostHomeItemExternal) GetID() string {
	return s.ID
}
