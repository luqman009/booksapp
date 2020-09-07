package main

import (
	"encoding/json"
	"io/ioutil"
)

type BookRepositoryInfrastructure struct {
	dataPath string
}

func NewBookRepoInfra(dataPath string) *BookRepositoryInfrastructure {
	return &BookRepositoryInfrastructure{dataPath}
}

func (bri *BookRepositoryInfrastructure) saveToFile(bookCollection *[]*Book) {
	file, _ := json.MarshalIndent(bookCollection, "", " ")
	_ = ioutil.WriteFile(bri.dataPath, file, 0644)
}

func (bri *BookRepositoryInfrastructure) readFile(bookCollection *[]*Book) {
	file, _ := ioutil.ReadFile(bri.dataPath)
	_ = json.Unmarshal(file, bookCollection)
}
