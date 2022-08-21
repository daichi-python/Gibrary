USE gibrary;

CREATE TABLE IF NOT EXISTS user_in_groupy (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user INT NOT NULL,
    groupy INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX i_user(user),
    INDEX i_groupy(groupy),
    FOREIGN KEY k_user(user) REFERENCES users(id),
    FOREIGN KEY k_groupy(groupy) REFERENCES groupies(id)
)