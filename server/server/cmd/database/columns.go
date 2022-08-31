package database

type ColumnParams []string

var ColumnValues = map[string]ColumnParams{
	"user": {
		"id", "name", "email", "birthday", "gender", "is_active",
	},

	"groupy": {
		"id", "name", "introduce", "max_people", "group_key", "is_opened",
	},

	"user_in_groupy": {
		"id", "users.id", "users.name", "users.email", "users.birthday", "users.gender", "users.is_active",
		"groupies.id", "groupies.name", "groupies.introduce", "groupies.max_people", "groupies.group_key", "groupies.is_opened",
	},

	"board_item": {
		"id", "types.id", "types.type", "categories.id", "categories.category",
		"title", "detail", "applicants", "is_opened",
	},

	"groupy_board_item": {
		"id", "groupies.id", "groupies.name", "groupies.introduce", "groupies.max_people", "groupies.group_key", "groupies.is_opened",
		"board_item.id", "board_item.types.id", "board_item.types.type", "board_item.categories.id", "board_item.categories.category",
		"title", "detail", "applicants", "is_opened",
	},
}
