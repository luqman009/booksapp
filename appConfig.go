package main

import "path"

type AppConfig struct {
	dataPath string
	db       []*Book
}

func NewConfig() *AppConfig {
	const DATA_PATH = "data"
	const DATA_FILE_NAME = "book_conlletion.json"

	var bookCollection = make([]*Book, 0)
	return &AppConfig{
		dataPath: path.Join(DATA_PATH, DATA_FILE_NAME),
		db:       bookCollection,
	}
}

func (c *AppConfig) getDb() []*Book {
	return c.db
}

func (c *AppConfig) getDbPath() string {
	return c.dataPath
}
