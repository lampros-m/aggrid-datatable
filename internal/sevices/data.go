package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/exp/rand"
)

var files []string = []string{
	"internal/data/jsons/example1.json",
	"internal/data/jsons/example2.json",
	"internal/data/jsons/example3.json",
}

type DataInterface interface {
	GetAll(ctx context.Context) (interface{}, error)
}

type Data struct {
}

func NewData() *Data {
	return &Data{}
}

func (o *Data) GetAll(ctx context.Context) (interface{}, error) {

	idx := rand.Intn(len(files))
	selectedFile := files[idx]

	file, err := os.Open(selectedFile)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		file.Close()
	}()

	var jsonData interface{}

	err = json.Unmarshal(content, &jsonData)
	if err != nil {
		log.Fatalf("Error parsing JSON data: %v", err)
	}

	_ = content

	fmt.Printf("Retrieved file: %s\n", selectedFile)

	return jsonData, nil
}
