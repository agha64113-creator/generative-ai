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

// Package cloudfunction provides a Vertex AI Search client for use in
// Google Cloud Functions.
package cloudfunction

import (
	"context"
	"fmt"
	"html"
	"regexp"
	"strings"

	discoveryengine "cloud.google.com/go/discoveryengine/apiv1"
	"cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// EngineDataType defines the supported data types for the search engine.
type EngineDataType string

const (
	EngineDataTypeUnstructured EngineDataType = "UNSTRUCTURED"
	EngineDataTypeStructured   EngineDataType = "STRUCTURED"
	EngineDataTypeWebsite      EngineDataType = "WEBSITE"
	EngineDataTypeBlended      EngineDataType = "BLENDED"
)

// EngineChunkType defines the supported chunk types for search results.
type EngineChunkType string

const (
	EngineChunkTypeDocumentWithSnippets           EngineChunkType = "DOCUMENT_WITH_SNIPPETS"
	EngineChunkTypeDocumentWithExtractiveSegments EngineChunkType = "DOCUMENT_WITH_EXTRACTIVE_SEGMENTS"
	EngineChunkTypeChunk                          EngineChunkType = "CHUNK"
	EngineChunkTypeNone                           EngineChunkType = "NONE"
)

// SummaryType defines the supported summary generation modes.
type SummaryType string

const (
	SummaryTypeNone            SummaryType = "NONE"
	SummaryTypeVertexAISearch  SummaryType = "VERTEX_AI_SEARCH"
)

// Config holds the configuration for the Vertex AI Search client.
type Config struct {
	ProjectID      string
	Location       string
	DataStoreID    string
	EngineDataType EngineDataType
	EngineChunkType EngineChunkType
	SummaryType    SummaryType
}

// SearchResult represents a simplified search result item.
type SearchResult struct {
	Metadata    map[string]any `json:"metadata"`
	PageContent string         `json:"page_content"`
}

// SearchResponse represents the full response from a search query.
type SearchResponse struct {
	Results          []map[string]any `json:"results"`
	TotalSize        int32            `json:"total_size"`
	AttributionToken string           `json:"attribution_token"`
	NextPageToken    string           `json:"next_page_token"`
	CorrectedQuery   string           `json:"corrected_query"`
	Summary          map[string]any   `json:"summary,omitempty"`
	SimplifiedResults []SearchResult  `json:"simplified_results"`
}

// Client is a wrapper around the Vertex AI Search service client.
type Client struct {
	config        Config
	searchClient  *discoveryengine.SearchClient
	servingConfig string
}

// NewClient creates and returns a new Client using the provided Config.
// The caller is responsible for calling Close() on the returned client.
func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	if cfg.EngineDataType == "" {
		cfg.EngineDataType = EngineDataTypeUnstructured
	}
	if cfg.EngineChunkType == "" {
		cfg.EngineChunkType = EngineChunkTypeChunk
	}
	if cfg.SummaryType == "" {
		cfg.SummaryType = SummaryTypeVertexAISearch
	}

	var opts []option.ClientOption
	if cfg.Location != "global" && cfg.Location != "" {
		endpoint := fmt.Sprintf("%s-discoveryengine.googleapis.com:443", cfg.Location)
		opts = append(opts, option.WithEndpoint(endpoint))
	}

	sc, err := discoveryengine.NewSearchClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("creating search client: %w", err)
	}

	servingConfig := fmt.Sprintf(
		"projects/%s/locations/%s/dataStores/%s/servingConfigs/default_config",
		cfg.ProjectID, cfg.Location, cfg.DataStoreID,
	)

	return &Client{
		config:        cfg,
		searchClient:  sc,
		servingConfig: servingConfig,
	}, nil
}

// Close releases the resources held by the client.
func (c *Client) Close() error {
	return c.searchClient.Close()
}

// Search performs a search query and returns simplified results.
func (c *Client) Search(ctx context.Context, query string) (*SearchResponse, error) {
	return c.SearchWithPageSize(ctx, query, 10)
}

// SearchWithPageSize performs a search query with a specific page size.
func (c *Client) SearchWithPageSize(ctx context.Context, query string, pageSize int32) (*SearchResponse, error) {
	req := c.buildSearchRequest(query, pageSize)
	it := c.searchClient.Search(ctx, req)

	resp, err := c.collectResults(it)
	if err != nil {
		return nil, fmt.Errorf("collecting search results: %w", err)
	}

	c.simplifyResults(resp)
	return resp, nil
}

// buildSearchRequest constructs a SearchRequest based on the client config.
func (c *Client) buildSearchRequest(query string, pageSize int32) *discoveryenginepb.SearchRequest {
	req := &discoveryenginepb.SearchRequest{
		ServingConfig: c.servingConfig,
		Query:         query,
		PageSize:      pageSize,
		QueryExpansionSpec: &discoveryenginepb.SearchRequest_QueryExpansionSpec{
			Condition: discoveryenginepb.SearchRequest_QueryExpansionSpec_AUTO,
		},
		SpellCorrectionSpec: &discoveryenginepb.SearchRequest_SpellCorrectionSpec{
			Mode: discoveryenginepb.SearchRequest_SpellCorrectionSpec_AUTO,
		},
		ContentSearchSpec: c.buildContentSearchSpec(),
	}
	return req
}

// buildContentSearchSpec builds the ContentSearchSpec based on config.
func (c *Client) buildContentSearchSpec() *discoveryenginepb.SearchRequest_ContentSearchSpec {
	spec := &discoveryenginepb.SearchRequest_ContentSearchSpec{}

	switch c.config.EngineChunkType {
	case EngineChunkTypeDocumentWithSnippets:
		spec.SnippetSpec = &discoveryenginepb.SearchRequest_ContentSearchSpec_SnippetSpec{
			ReturnSnippet: true,
		}
	case EngineChunkTypeDocumentWithExtractiveSegments:
		spec.SnippetSpec = &discoveryenginepb.SearchRequest_ContentSearchSpec_SnippetSpec{
			ReturnSnippet: true,
		}
		spec.ExtractiveContentSpec = &discoveryenginepb.SearchRequest_ContentSearchSpec_ExtractiveContentSpec{
			MaxExtractiveAnswerCount:       1,
			ReturnExtractiveSegmentScore:   true,
		}
	case EngineChunkTypeChunk:
		spec.ChunkSpec = &discoveryenginepb.SearchRequest_ContentSearchSpec_ChunkSpec{
			NumPreviousChunks: 0,
			NumNextChunks:     0,
		}
	}

	if c.config.SummaryType == SummaryTypeVertexAISearch {
		spec.SummarySpec = &discoveryenginepb.SearchRequest_ContentSearchSpec_SummarySpec{
			SummaryResultCount:              5,
			IncludeCitations:                true,
			IgnoreAdversarialQuery:          true,
			IgnoreNonSummarySeekingQuery:    true,
		}
	}

	return spec
}

// collectResults iterates over the search iterator and collects all results.
func (c *Client) collectResults(it *discoveryengine.SearchResponse_SearchResultIterator) (*SearchResponse, error) {
	resp := &SearchResponse{}

	for {
		result, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		resp.Results = append(resp.Results, searchResultToMap(result))
	}

	// Populate metadata from the underlying response page.
	if page, ok := it.Response.(*discoveryenginepb.SearchResponse); ok && page != nil {
		resp.TotalSize = page.TotalSize
		resp.AttributionToken = page.AttributionToken
		resp.NextPageToken = page.NextPageToken
		resp.CorrectedQuery = page.CorrectedQuery

		if page.Summary != nil {
			resp.Summary = map[string]any{
				"summary_text": page.Summary.SummaryText,
			}
		}
	}

	return resp, nil
}

// searchResultToMap converts a SearchResponse_SearchResult proto to a map.
func searchResultToMap(r *discoveryenginepb.SearchResponse_SearchResult) map[string]any {
	m := map[string]any{"id": r.Id}
	if r.Document != nil {
		doc := map[string]any{
			"id":   r.Document.Id,
			"name": r.Document.Name,
		}
		if structData := r.Document.GetStructData(); structData != nil {
			fields := make(map[string]any)
			for k, v := range structData.Fields {
				fields[k] = v.AsInterface()
			}
			doc["struct_data"] = fields
		}
		if r.Document.DerivedStructData != nil {
			derived := make(map[string]any)
			for k, v := range r.Document.DerivedStructData.Fields {
				derived[k] = v.AsInterface()
			}
			doc["derived_struct_data"] = derived
		}
		if jsonData := r.Document.GetJsonData(); jsonData != "" {
			doc["json_data"] = jsonData
		}
		m["document"] = doc
	}
	if r.Chunk != nil {
		chunk := map[string]any{
			"id":              r.Chunk.Id,
			"content":         r.Chunk.Content,
			"relevance_score": r.Chunk.RelevanceScore,
		}
		if r.Chunk.PageSpan != nil {
			chunk["page_span"] = map[string]any{
				"page_start": r.Chunk.PageSpan.PageStart,
				"page_end":   r.Chunk.PageSpan.PageEnd,
			}
		}
		if r.Chunk.DocumentMetadata != nil {
			meta := map[string]any{
				"uri":   r.Chunk.DocumentMetadata.Uri,
				"title": r.Chunk.DocumentMetadata.Title,
			}
			if r.Chunk.DocumentMetadata.StructData != nil {
				for k, v := range r.Chunk.DocumentMetadata.StructData.Fields {
					meta[k] = v.AsInterface()
				}
			}
			chunk["document_metadata"] = meta
		}
		if r.Chunk.DerivedStructData != nil {
			derived := make(map[string]any)
			for k, v := range r.Chunk.DerivedStructData.Fields {
				derived[k] = v.AsInterface()
			}
			chunk["derived_struct_data"] = derived
		}
		m["chunk"] = chunk
	}
	return m
}

// simplifyResults populates the SimplifiedResults field of a SearchResponse.
func (c *Client) simplifyResults(resp *SearchResponse) {
	for _, result := range resp.Results {
		if doc, ok := result["document"].(map[string]any); ok {
			resp.SimplifiedResults = append(resp.SimplifiedResults, c.parseDocumentResult(doc))
		} else if chunk, ok := result["chunk"].(map[string]any); ok {
			resp.SimplifiedResults = append(resp.SimplifiedResults, parseChunkResult(chunk))
		}
	}
}

// parseDocumentResult converts a raw document map to a SearchResult.
func (c *Client) parseDocumentResult(document map[string]any) SearchResult {
	metadata := make(map[string]any)

	if derived, ok := document["derived_struct_data"].(map[string]any); ok {
		for k, v := range derived {
			metadata[k] = v
		}
	}
	if structData, ok := document["struct_data"].(map[string]any); ok {
		for k, v := range structData {
			metadata[k] = v
		}
	}

	sr := SearchResult{Metadata: metadata}

	if c.config.EngineDataType == EngineDataTypeStructured {
		if structData, ok := document["struct_data"].(map[string]any); ok {
			sr.PageContent = fmt.Sprintf("%v", structData)
		}
		return sr
	}

	if answers, ok := metadata["extractive_answers"]; ok {
		sr.PageContent = parseSegments(answers)
	} else if snippets, ok := metadata["snippets"]; ok {
		sr.PageContent = parseSnippets(snippets)
	} else {
		if content, ok := metadata["content"].(string); ok {
			sr.PageContent = content
		}
	}

	return sr
}

// parseChunkResult converts a raw chunk map to a SearchResult.
func parseChunkResult(chunk map[string]any) SearchResult {
	metadata := map[string]any{
		"chunk_id": chunk["id"],
		"score":    chunk["relevance_score"],
	}

	if pageSpan, ok := chunk["page_span"].(map[string]any); ok {
		metadata["page"] = pageSpan["page_start"]
		metadata["page_span_end"] = pageSpan["page_end"]
	}

	if docMeta, ok := chunk["document_metadata"].(map[string]any); ok {
		for k, v := range docMeta {
			metadata[k] = v
		}
	}

	if derived, ok := chunk["derived_struct_data"].(map[string]any); ok {
		for k, v := range derived {
			metadata[k] = v
		}
	}

	content := ""
	if c, ok := chunk["content"].(string); ok {
		content = stripContent(c)
	}

	return SearchResult{
		Metadata:    metadata,
		PageContent: content,
	}
}

// parseSegments formats extractive answer segments into a readable string.
func parseSegments(raw any) string {
	segments, ok := raw.([]any)
	if !ok {
		return ""
	}
	var parts []string
	for _, s := range segments {
		seg, ok := s.(map[string]any)
		if !ok {
			continue
		}
		content := stripContent(stringVal(seg, "content"))
		pageNum := stringVal(seg, "page_number")
		if pageNum == "" {
			pageNum = fmt.Sprintf("%v", seg["pageNumber"])
		}
		score := fmt.Sprintf("%v", seg["score"])
		parts = append(parts, fmt.Sprintf(
			"On page %s with a relevance score of %s:\n```\n%s\n```",
			pageNum, score, content,
		))
	}
	return strings.Join(parts, "\n\n")
}

// parseSnippets formats snippet results into a readable string.
func parseSnippets(raw any) string {
	snippets, ok := raw.([]any)
	if !ok {
		return ""
	}
	var parts []string
	for _, s := range snippets {
		snip, ok := s.(map[string]any)
		if !ok {
			continue
		}
		status := stringVal(snip, "snippetStatus")
		if status == "" {
			status = stringVal(snip, "snippet_status")
		}
		if status == "SUCCESS" {
			parts = append(parts, stripContent(stringVal(snip, "snippet")))
		}
	}
	return strings.Join(parts, "\n\n")
}

// stripContent removes HTML tags and unescapes HTML entities from text.
func stripContent(text string) string {
	re := regexp.MustCompile("<.*?>")
	text = re.ReplaceAllString(text, "")
	return strings.TrimSpace(html.UnescapeString(text))
}

// stringVal safely retrieves a string value from a map.
func stringVal(m map[string]any, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}
