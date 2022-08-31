USE gibrary;

CREATE TABLE IF NOT EXISTS categories (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    category VARCHAR(10) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO categories(category) VALUES("game");
INSERT INTO categories(category) VALUES("sport");
INSERT INTO categories(category) VALUES("work");
INSERT INTO categories(category) VALUES("other");