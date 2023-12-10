package main

import (
	"database/sql"
	"log"
	"os"
)

var statements = map[string]*sql.Stmt{}

func dbInit() {

	_, err := os.Stat(fileDB)

	if os.IsNotExist(err) {
		*reset = true
	}

	db, err = sql.Open("sqlite3", fileDB)
	if err != nil {
		log.Fatal(err)
	}
	if *reset {
		_, err := db.Exec(`
		DROP TABLE IF EXISTS users;
		DROP TABLE IF EXISTS session;
		DROP TABLE IF EXISTS post;
		DROP TABLE IF EXISTS category;
		DROP TABLE IF EXISTS post_category;
		DROP TABLE IF EXISTS comment;
		DROP TABLE IF EXISTS message;
		
		CREATE TABLE comment (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			user_id INTEGER NOT NULL REFERENCES users (id),
			post_id INTEGER NOT NULL REFERENCES post (id),
			content VARCHAR NOT NULL,
			picture BLOB,
			created_at DATETIME NOT NULL
			);
		CREATE TABLE message (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			from_id INTEGER NOT NULL REFERENCES users (id),
			to_id INTEGER NOT NULL REFERENCES users (id),
			content VARCHAR NOT NULL,
			created_at DATETIME NOT NULL
			);
		CREATE TABLE groups (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT,
			creator INTEGER,
			created_at TIMESTAMP,
			privacy TEXT,
			FOREIGN KEY (creator) REFERENCES users (id)
			);
		CREATE TABLE group_members (
			group_id INTEGER,
			member_id INTEGER,
			PRIMARY KEY (group_id, member_id),
			FOREIGN KEY (group_id) REFERENCES groups(id),
			FOREIGN KEY (member_id) REFERENCES users(id)
			);
		CREATE TABLE group_pending_members (
			group_id INTEGER,
			member_id INTEGER,
			PRIMARY KEY (group_id, member_id),
			FOREIGN KEY (group_id) REFERENCES groups(id),
			FOREIGN KEY (member_id) REFERENCES users(id)
			);
		CREATE TABLE category (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
			);
		CREATE TABLE followers (
			user_id INTEGER,
			follower_id INTEGER,
			PRIMARY KEY (user_id, follower_id),
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (follower_id) REFERENCES users(id)
			);
		CREATE TABLE followers_pending (
			user_id INTEGER,
			follower_id INTEGER,
			PRIMARY KEY (user_id, follower_id),
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (follower_id) REFERENCES users(id)
			);
		CREATE TABLE post_category (
			post_id INTEGER NOT NULL,
			category_id INTEGER,
			PRIMARY KEY (post_id, category_id),
			FOREIGN KEY (post_id) REFERENCES post(id),
			FOREIGN KEY (category_id) REFERENCES category(id)
			);
		CREATE TABLE almost_private (
			user_id INTEGER,
			post_id INTEGER,
			PRIMARY KEY (user_id, post_id),
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (post_id) REFERENCES post(id)
			);`)

		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("DB reset")
	}
}

// it was last added by @sagarishere, just moved it here, perhaps it better and will be used
// "getGroupPosts": `SELECT post.id, nickname, first_name, last_name, title, categories, content, created_at
// 							FROM post INNER JOIN users ON user_id=users.id WHERE post.id IN
// 							(SELECT post_id FROM post_category WHERE category_id = ?)
// 							ORDER BY created_at DESC;`,
