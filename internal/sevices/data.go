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

type DataEndpoints map[string]string

type DataInterface interface {
	GetByUrl(ctx context.Context, urlID string) (interface{}, error)
	MockData(ctx context.Context) (interface{}, error)
}

type Data struct {
}

func NewData() *Data {
	return &Data{}
}

func (o *Data) GetByUrl(ctx context.Context, urlID string) (interface{}, error) {
	var err error

	file, err := os.Open(config.GetConfig().DATA_ENDPOINTS)
	if err != nil {
		return nil, fmt.Errorf("failed to open dataendpoints json file: %w", err)
	}
	defer file.Close()

	var dataEndpoints DataEndpoints
	err = json.NewDecoder(file).Decode(&dataEndpoints)
	if err != nil {
		return nil, fmt.Errorf("failed to decode dataendpoints json file: %w", err)
	}

	url, found := dataEndpoints[urlID]
	if !found {
		return nil, fmt.Errorf("urlID not found: %s", urlID)
	}

	resp, err := http.Get(url)
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

	// Uncomment for debugging
	// _ = content
	// fmt.Printf("Retrieved file: %s\n", selectedFile)

	return jsonData, nil
}
