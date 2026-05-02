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

// Package cloudfunction contains integration tests for the Vertex AI Search
// client. These tests interact with the actual Vertex AI Search API and
// require the following environment variables to be set:
//
//	PROJECT_ID      - Your Google Cloud project ID
//	LOCATION        - The location of your data store (default: "global")
//	DATA_STORE_ID   - The ID of your Vertex AI Search data store
//	ENGINE_DATA_TYPE  - UNSTRUCTURED | STRUCTURED | WEBSITE | BLENDED
//	ENGINE_CHUNK_TYPE - DOCUMENT_WITH_SNIPPETS | DOCUMENT_WITH_EXTRACTIVE_SEGMENTS | CHUNK
//	SUMMARY_TYPE      - NONE | VERTEX_AI_SEARCH
//
// Run with: go test -v -run Integration ./...
package cloudfunction

import (
	"context"
	"os"
	"testing"
)

// integrationTestConfig reads configuration from environment variables.
// The test is skipped when PROJECT_ID or DATA_STORE_ID are not set.
func integrationTestConfig(t *testing.T) Config {
	t.Helper()

	projectID := os.Getenv("PROJECT_ID")
	dataStoreID := os.Getenv("DATA_STORE_ID")

	if projectID == "" || dataStoreID == "" {
		t.Skip("Skipping integration test: PROJECT_ID and DATA_STORE_ID environment variables are required")
	}

	return Config{
		ProjectID:       projectID,
		Location:        getEnv("LOCATION", "global"),
		DataStoreID:     dataStoreID,
		EngineDataType:  EngineDataType(getEnv("ENGINE_DATA_TYPE", string(EngineDataTypeUnstructured))),
		EngineChunkType: EngineChunkType(getEnv("ENGINE_CHUNK_TYPE", string(EngineChunkTypeDocumentWithExtractiveSegments))),
		SummaryType:     SummaryType(getEnv("SUMMARY_TYPE", string(SummaryTypeVertexAISearch))),
	}
}

// TestIntegration_Search performs a search using the actual Vertex AI Search API
// and verifies the response structure.
func TestIntegration_Search(t *testing.T) {
	cfg := integrationTestConfig(t)
	ctx := context.Background()

	client, err := NewClient(ctx, cfg)
	if err != nil {
		t.Fatalf("NewClient() error: %v", err)
	}
	defer client.Close()

	resp, err := client.Search(ctx, "test query")
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if resp == nil {
		t.Fatal("Search() returned nil response")
	}

	// Verify top-level response fields.
	if resp.Results == nil {
		t.Error("expected non-nil Results slice")
	}
	if resp.TotalSize < 0 {
		t.Errorf("TotalSize = %d, want >= 0", resp.TotalSize)
	}
	if resp.SimplifiedResults == nil {
		t.Error("expected non-nil SimplifiedResults slice")
	}

	// Verify simplified result structure when results are present.
	for i, sr := range resp.SimplifiedResults {
		if sr.Metadata == nil {
			t.Errorf("SimplifiedResults[%d].Metadata is nil", i)
		}
	}
}

// TestIntegration_SearchWithSummary verifies that summary generation works
// when SummaryType is set to VERTEX_AI_SEARCH.
func TestIntegration_SearchWithSummary(t *testing.T) {
	cfg := integrationTestConfig(t)
	cfg.SummaryType = SummaryTypeVertexAISearch
	ctx := context.Background()

	client, err := NewClient(ctx, cfg)
	if err != nil {
		t.Fatalf("NewClient() error: %v", err)
	}
	defer client.Close()

	resp, err := client.Search(ctx, "What is the name of the company?")
	if err != nil {
		t.Fatalf("Search() error: %v", err)
	}

	if resp == nil {
		t.Fatal("Search() returned nil response")
	}

	// Response structure should always be present.
	if resp.Results == nil {
		t.Error("expected non-nil Results slice")
	}

	// When a summary is returned, verify its structure.
	if resp.Summary != nil {
		if _, ok := resp.Summary["summary_text"]; !ok {
			t.Error("expected 'summary_text' key in Summary")
		}
	}
}

// TestIntegration_SearchWithPageSize verifies that the page_size parameter
// limits the number of results returned.
func TestIntegration_SearchWithPageSize(t *testing.T) {
	cfg := integrationTestConfig(t)
	ctx := context.Background()

	client, err := NewClient(ctx, cfg)
	if err != nil {
		t.Fatalf("NewClient() error: %v", err)
	}
	defer client.Close()

	const pageSize = int32(3)
	resp, err := client.SearchWithPageSize(ctx, "test", pageSize)
	if err != nil {
		t.Fatalf("SearchWithPageSize() error: %v", err)
	}

	if int32(len(resp.Results)) > pageSize {
		t.Errorf("len(Results) = %d, want <= %d", len(resp.Results), pageSize)
	}
}
