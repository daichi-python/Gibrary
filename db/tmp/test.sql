use gibrary;

SET foreign_key_checks = 0;

TRUNCATE groupy_board_items;
TRUNCATE user_like_board_items;
TRUNCATE user_in_groupy;
TRUNCATE post_home_items;

TRUNCATE users;
TRUNCATE groupies;
TRUNCATE board_items;
TRUNCATE home_items;

SET foreign_key_checks = 1;

INSERT INTO users(name, email) VALUES("Monkey", "monkey001.gmail.com");
INSERT INTO users(name, email) VALUES("Cat", "cat002.gmail.com");
INSERT INTO users(name, email) VALUES("Dog", "dog003.gmail.com");
INSERT INTO users(name, email) VALUES("Banana", "banana004.gmail.com");
INSERT INTO users(name, email) VALUES("Grape", "grape005.gmail.com");
INSERT INTO users(name, email) VALUES("Apple", "apple006.gmail.com");
INSERT INTO users(name, email) VALUES("Orange", "orange007.gmail.com");

INSERT INTO groupies(name, introduce) VALUES("Google", "This is google group.");
INSERT INTO groupies(name, introduce) VALUES("Yahoo", "This is yahoo group.");
INSERT INTO groupies(name, introduce) VALUES("Facebook", "This is facebook group.");
INSERT INTO groupies(name, introduce) VALUES("Informatica", "This is informatica group.");
INSERT INTO groupies(name, introduce) VALUES("Splunk", "This is splunk group.");
INSERT INTO groupies(name, introduce) VALUES("Twitter", "This is twitter group.");
INSERT INTO groupies(name, introduce) VALUES("Amazon", "This is amazon group.");

INSERT INTO user_in_groupy(user, groupy) VALUES(1, 1);
INSERT INTO user_in_groupy(user, groupy) VALUES(2, 2);
INSERT INTO user_in_groupy(user, groupy) VALUES(3, 3);
INSERT INTO user_in_groupy(user, groupy) VALUES(4, 4);
INSERT INTO user_in_groupy(user, groupy) VALUES(5, 5);
INSERT INTO user_in_groupy(user, groupy) VALUES(6, 6);
INSERT INTO user_in_groupy(user, groupy) VALUES(7, 7);

INSERT INTO board_items(type, category, title, detail) VALUES(1, 1, "Mario", "Mario board item.");
INSERT INTO board_items(type, category, title, detail) VALUES(2, 4, "Rock", "Rock board item.");
INSERT INTO board_items(type, category, title, detail) VALUES(2, 2, "Soccer", "Soccer board item.");
INSERT INTO board_items(type, category, title, detail) VALUES(2, 2, "Baseball", "Baseball board item.");
INSERT INTO board_items(type, category, title, detail) VALUES(1, 3, "Bank", "Bank board item.");
INSERT INTO board_items(type, category, title, detail) VALUES(1, 1, "SB", "SB board item.");
INSERT INTO board_items(type, category, title, detail) VALUES(2, 1, "Card Game", "Card game board item.");

INSERT INTO groupy_board_items(groupy, boardItem) VALUES(1, 1);
INSERT INTO groupy_board_items(groupy, boardItem) VALUES(2, 2);
INSERT INTO groupy_board_items(groupy, boardItem) VALUES(3, 3);
INSERT INTO groupy_board_items(groupy, boardItem) VALUES(4, 4);
INSERT INTO groupy_board_items(groupy, boardItem) VALUES(5, 5);
INSERT INTO groupy_board_items(groupy, boardItem) VALUES(6, 6);
INSERT INTO groupy_board_items(groupy, boardItem) VALUES(7, 7);

INSERT INTO user_like_board_items(user, boardItem) VALUES(1, 1);
INSERT INTO user_like_board_items(user, boardItem) VALUES(2, 2);
INSERT INTO user_like_board_items(user, boardItem) VALUES(3, 3);
INSERT INTO user_like_board_items(user, boardItem) VALUES(4, 4);
INSERT INTO user_like_board_items(user, boardItem) VALUES(5, 5);
INSERT INTO user_like_board_items(user, boardItem) VALUES(6, 6);
INSERT INTO user_like_board_items(user, boardItem) VALUES(7, 7);

INSERT INTO home_items(detail) VALUES("Power Center");
INSERT INTO home_items(detail) VALUES("Cloud Data Integration");
INSERT INTO home_items(detail) VALUES("Cloud Application Integration");
INSERT INTO home_items(detail) VALUES("Cloud Integration Hub");
INSERT INTO home_items(detail) VALUES("API Manager");
INSERT INTO home_items(detail) VALUES("Data Quality");
INSERT INTO home_items(detail) VALUES("Process Engine");

INSERT INTO post_home_items(user, homeitem) VALUES(1, 1);
INSERT INTO post_home_items(groupy, homeitem) VALUES(2, 2);
INSERT INTO post_home_items(groupy, homeitem) VALUES(3, 3);
INSERT INTO post_home_items(user, homeitem) VALUES(4, 4);
INSERT INTO post_home_items(groupy, homeitem) VALUES(5, 5);
INSERT INTO post_home_items(user, homeitem) VALUES(6, 6);
INSERT INTO post_home_items(groupy, homeitem) VALUES(7, 7);