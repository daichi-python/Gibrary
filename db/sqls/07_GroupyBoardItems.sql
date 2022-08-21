USE gibrary;

CREATE TABLE IF NOT EXISTS groupy_board_items (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    groupy INT NOT NULL,
    boardItem INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX i_groupy(groupy),
    INDEX i_board_item(boarditem),
    FOREIGN KEY k_groupy(groupy) REFERENCES groupies(id),
    FOREIGN KEY k_board_item(boarditem) REFERENCES board_items(id)
)