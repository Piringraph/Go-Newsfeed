package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/*

CREATE TABLE IF NOT EXISTS "table-newsfeed" (
	"id"	INTEGER NOT NULL,
	"content"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
*/

func main() {
	db, err := sql.Open("sqlite3", "./newsfeed.db")

	db.Prepare ("CREATE TABLE IF NOT EXISTS "table-newsfeed" (
		"id"	INTEGER NOT NULL,
		"content"	TEXT,
		PRIMARY KEY("id" AUTOINCREMENT)
	);")
}
