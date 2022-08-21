USE gibrary;

CREATE TABLE IF NOT EXISTS board_items (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    type INT NOT NULL,
    category INT NOt NULL,
    title VARCHAR(30) NOT NULL,
    detail TEXT,
    applicants INT,
    is_opened BOOLEAN NOT NULL DEFAULT TRUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX i_type(type),
    INDEX i_category(category),
    FOREIGN KEY k_type(type) REFERENCES types(id),
    FOREIGN KEY k_category(category) REFERENCES categories(id)
)