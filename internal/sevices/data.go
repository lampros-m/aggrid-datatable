package services

import (
	"aggriddatatable/internal/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/exp/rand"
)

var files []string = []string{
	"internal/data/jsons/example1.json",
	"internal/data/jsons/example2.json",
	"internal/data/jsons/example3.json",
}

type DataInterface interface {
	MockData(ctx context.Context) (interface{}, error)
	GetAll(ctx context.Context) (interface{}, error)
}

type Data struct {
}

func NewData() *Data {
	return &Data{}
}

func (o *Data) MockData(ctx context.Context) (interface{}, error) {
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

func (o *Data) GetAll(ctx context.Context) (interface{}, error) {
	c := config.GetConfig()

	if c.DATA_ENDPOINT == "" {
		return nil, fmt.Errorf("DATA_ENDPOINT is empty")
	}

	resp, err := http.Get(c.DATA_ENDPOINT)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from the endpoint: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var jsonData interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON data: %w", err)
	}

	return jsonData, nil
}
