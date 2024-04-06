package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GraphqlClient struct {
	url string
}

func (c *GraphqlClient) callApi(query string, variables map[string]interface{}) (graphqlResponse, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"query":     query,
		"variables": variables,
	})
	if err != nil {
		return graphqlResponse{}, fmt.Errorf("failed creating request body: %v", err)
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(requestBody))
	if err != nil {
		return graphqlResponse{}, fmt.Errorf("failed creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return graphqlResponse{}, fmt.Errorf("failed sending request: %v", err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return graphqlResponse{}, fmt.Errorf("failed reading response body: %v", err)
	}

	var response graphqlResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return graphqlResponse{}, fmt.Errorf("failed parsing JSON: %v", err)
	}

	return response, nil
}

func NewGraphqlClient(url string) *GraphqlClient {
	return &GraphqlClient{url: url}
}
