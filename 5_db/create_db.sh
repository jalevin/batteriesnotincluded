#! /bin/bash

sqlite3 Posts.db

CREATE TABLE posts(
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   title 					TEXT    NOT NULL,
   body 					TEXT
);

INSERT INTO posts (title, body) VALUES ("My First Post", "It's short form writing today!");

