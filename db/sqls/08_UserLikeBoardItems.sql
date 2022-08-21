USE gibrary;

CREATE TABLE IF NOT EXISTS user_like_board_items (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user INT NOT NULL,
    boarditem INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX i_user(user),
    INDEX i_board_item(boarditem),
    FOREIGN KEY k_user(user) REFERENCES users(id),
    FOREIGN KEY k_board_item(boarditem) REFERENCES board_items(id)
)