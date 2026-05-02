// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cloudfunction provides an HTTP Cloud Function that performs
// searches using the Vertex AI Search API.
//
// For deployment instructions, environment variable setup, and usage examples,
// please refer to the README.md file.
package cloudfunction

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("VertexAISearch", vertexAISearch)
}

// newClientFromEnv creates a new Client from environment variables.
func newClientFromEnv(ctx context.Context) (*Client, error) {
	cfg := Config{
		ProjectID:       getEnv("PROJECT_ID", "your-project"),
		Location:        getEnv("LOCATION", "global"),
		DataStoreID:     getEnv("DATA_STORE_ID", "your-data-store"),
		EngineDataType:  EngineDataType(getEnv("ENGINE_DATA_TYPE", string(EngineDataTypeUnstructured))),
		EngineChunkType: EngineChunkType(getEnv("ENGINE_CHUNK_TYPE", string(EngineChunkTypeChunk))),
		SummaryType:     SummaryType(getEnv("SUMMARY_TYPE", string(SummaryTypeVertexAISearch))),
	}
	return NewClient(ctx, cfg)
}

// vertexAISearch handles HTTP requests for Vertex AI Search.
//
// It reads the search_term from the JSON request body or query parameters,
// performs the search, and returns the results as JSON. It also sets CORS
// headers to allow cross-origin requests.
func vertexAISearch(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for preflight requests.
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	searchTerm := ""

	// Try JSON body first.
	if r.Header.Get("Content-Type") == "application/json" {
		var body struct {
			SearchTerm string `json:"search_term"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err == nil && body.SearchTerm != "" {
			searchTerm = body.SearchTerm
		}
	}

	// Fall back to query parameter.
	if searchTerm == "" {
		searchTerm = r.URL.Query().Get("search_term")
	}

	if searchTerm == "" {
		http.Error(w, `{"error":"No search term provided"}`, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	client, err := newClientFromEnv(ctx)
	if err != nil {
		http.Error(w,
			fmt.Sprintf(`{"error":"Failed to create search client: %s"}`, err),
			http.StatusInternalServerError,
		)
		return
	}
	defer client.Close()

	results, err := client.Search(ctx, searchTerm)
	if err != nil {
		http.Error(w,
			fmt.Sprintf(`{"error":"Error calling Vertex AI Search API: %s"}`, err),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w,
			fmt.Sprintf(`{"error":"Failed to encode response: %s"}`, err),
			http.StatusInternalServerError,
		)
	}
}

// getEnv returns the value of the environment variable named by key,
// or fallback if the variable is not set or is empty.
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
