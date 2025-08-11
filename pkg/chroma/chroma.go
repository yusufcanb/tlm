package chroma

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChromaClient struct {
	baseURL string
	client  *http.Client
}

func NewChromaClient(baseURL string) *ChromaClient {
	return &ChromaClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

type AddRequest struct {
	Documents []string `json:"documents"`
	IDs       []string `json:"ids"`
}

type QueryRequest struct {
	QueryTexts []string `json:"query_texts"`
	NResults   int      `json:"n_results"`
}

type QueryResponse struct {
	Documents [][]string `json:"documents"`
}

func (c *ChromaClient) Add(collectionName string, req *AddRequest) error {
	url := fmt.Sprintf("%s/api/v1/collections/%s/add", c.baseURL, collectionName)
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add documents: %s", resp.Status)
	}

	return nil
}

func (c *ChromaClient) Query(collectionName string, req *QueryRequest) (*QueryResponse, error) {
	url := fmt.Sprintf("%s/api/v1/collections/%s/query", c.baseURL, collectionName)
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to query documents: %s", resp.Status)
	}

	var queryResp QueryResponse
	if err := json.NewDecoder(resp.Body).Decode(&queryResp); err != nil {
		return nil, err
	}

	return &queryResp, nil
}
