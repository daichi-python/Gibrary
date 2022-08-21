USE gibrary;

CREATE TABLE IF NOT EXISTS post_home_items (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user INT,
    groupy INT,
    homeitem INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX i_user(user),
    INDEX i_groupy(groupy),
    INDEX i_home_item(homeitem),
    FOREIGN KEY k_user(user) REFERENCES users(id),
    FOREIGN KEY k_groupy(groupy) REFERENCES groupies(id),
    FOREIGN KEY k_home_item(homeitem) REFERENCES home_items(id)
)