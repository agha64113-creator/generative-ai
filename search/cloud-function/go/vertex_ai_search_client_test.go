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

package cloudfunction

import (
	"strings"
	"testing"
)

// --- stripContent ---

func TestStripContent_PlainText(t *testing.T) {
	got := stripContent("hello world")
	if got != "hello world" {
		t.Errorf("stripContent() = %q, want %q", got, "hello world")
	}
}

func TestStripContent_HTMLTags(t *testing.T) {
	input := "<p>Test <strong>content</strong></p>"
	want := "Test content"
	got := stripContent(input)
	if got != want {
		t.Errorf("stripContent(%q) = %q, want %q", input, got, want)
	}
}

func TestStripContent_HTMLEntities(t *testing.T) {
	input := `<p>Test &quot;quotes&quot; &amp; &lt;tags&gt;</p>`
	want := `Test "quotes" & <tags>`
	got := stripContent(input)
	if got != want {
		t.Errorf("stripContent(%q) = %q, want %q", input, got, want)
	}
}

func TestStripContent_Whitespace(t *testing.T) {
	input := "  hello  "
	want := "hello"
	got := stripContent(input)
	if got != want {
		t.Errorf("stripContent(%q) = %q, want %q", input, got, want)
	}
}

// --- parseChunkResult ---

func TestParseChunkResult_Basic(t *testing.T) {
	chunk := map[string]any{
		"id":              "chunk1",
		"relevance_score": 0.95,
		"content":         "Test content",
		"document_metadata": map[string]any{
			"title": "Test Title",
			"uri":   "https://example.com",
		},
		"page_span": map[string]any{
			"page_start": int32(1),
			"page_end":   int32(2),
		},
	}
	result := parseChunkResult(chunk)

	if result.PageContent != "Test content" {
		t.Errorf("PageContent = %q, want %q", result.PageContent, "Test content")
	}
	if result.Metadata["chunk_id"] != "chunk1" {
		t.Errorf("chunk_id = %v, want %v", result.Metadata["chunk_id"], "chunk1")
	}
	if result.Metadata["score"] != 0.95 {
		t.Errorf("score = %v, want %v", result.Metadata["score"], 0.95)
	}
	if result.Metadata["title"] != "Test Title" {
		t.Errorf("title = %v, want %v", result.Metadata["title"], "Test Title")
	}
	if result.Metadata["page"] != int32(1) {
		t.Errorf("page = %v, want %v", result.Metadata["page"], int32(1))
	}
}

func TestParseChunkResult_NoPageSpan(t *testing.T) {
	chunk := map[string]any{
		"id":      "chunk2",
		"content": "Some content",
	}
	result := parseChunkResult(chunk)
	if _, ok := result.Metadata["page"]; ok {
		t.Error("expected no 'page' key when page_span is absent")
	}
}

// --- parseSegments ---

func TestParseSegments_Basic(t *testing.T) {
	segments := []any{
		map[string]any{
			"content":     "Segment content",
			"page_number": "3",
			"score":       "0.9",
		},
	}
	got := parseSegments(segments)
	if got == "" {
		t.Error("parseSegments() returned empty string")
	}
	for _, want := range []string{"On page 3", "0.9", "Segment content"} {
		if !contains(got, want) {
			t.Errorf("parseSegments() = %q, want it to contain %q", got, want)
		}
	}
}

func TestParseSegments_Empty(t *testing.T) {
	got := parseSegments([]any{})
	if got != "" {
		t.Errorf("parseSegments([]) = %q, want empty string", got)
	}
}

func TestParseSegments_InvalidType(t *testing.T) {
	got := parseSegments("not a slice")
	if got != "" {
		t.Errorf("parseSegments(invalid) = %q, want empty string", got)
	}
}

// --- parseSnippets ---

func TestParseSnippets_Success(t *testing.T) {
	snippets := []any{
		map[string]any{"snippetStatus": "SUCCESS", "snippet": "First snippet"},
		map[string]any{"snippetStatus": "FAILED", "snippet": "Should be skipped"},
		map[string]any{"snippetStatus": "SUCCESS", "snippet": "Second snippet"},
	}
	got := parseSnippets(snippets)
	if !contains(got, "First snippet") {
		t.Errorf("parseSnippets() = %q, want it to contain 'First snippet'", got)
	}
	if !contains(got, "Second snippet") {
		t.Errorf("parseSnippets() = %q, want it to contain 'Second snippet'", got)
	}
	if contains(got, "Should be skipped") {
		t.Errorf("parseSnippets() = %q, should not contain 'Should be skipped'", got)
	}
}

func TestParseSnippets_AltStatusKey(t *testing.T) {
	snippets := []any{
		map[string]any{"snippet_status": "SUCCESS", "snippet": "Alt key snippet"},
	}
	got := parseSnippets(snippets)
	if !contains(got, "Alt key snippet") {
		t.Errorf("parseSnippets() = %q, want it to contain 'Alt key snippet'", got)
	}
}

// --- parseDocumentResult ---

func TestParseDocumentResult_ExtractiveAnswers(t *testing.T) {
	c := &Client{config: Config{EngineDataType: EngineDataTypeUnstructured}}
	doc := map[string]any{
		"derived_struct_data": map[string]any{
			"title": "Employee Benefits",
			"link":  "gs://company-docs/file.pdf",
			"extractive_answers": []any{
				map[string]any{"content": "Benefits text", "page_number": "1", "score": "0.8"},
			},
		},
	}
	result := c.parseDocumentResult(doc)
	if result.Metadata["title"] != "Employee Benefits" {
		t.Errorf("title = %v, want 'Employee Benefits'", result.Metadata["title"])
	}
	if !contains(result.PageContent, "Benefits text") {
		t.Errorf("PageContent = %q, want it to contain 'Benefits text'", result.PageContent)
	}
	if !contains(result.PageContent, "On page 1") {
		t.Errorf("PageContent = %q, want it to contain 'On page 1'", result.PageContent)
	}
}

func TestParseDocumentResult_Snippets(t *testing.T) {
	c := &Client{config: Config{EngineDataType: EngineDataTypeUnstructured}}
	doc := map[string]any{
		"derived_struct_data": map[string]any{
			"snippets": []any{
				map[string]any{"snippetStatus": "SUCCESS", "snippet": "A snippet"},
			},
		},
	}
	result := c.parseDocumentResult(doc)
	if !contains(result.PageContent, "A snippet") {
		t.Errorf("PageContent = %q, want it to contain 'A snippet'", result.PageContent)
	}
}

func TestParseDocumentResult_FallbackContent(t *testing.T) {
	c := &Client{config: Config{EngineDataType: EngineDataTypeUnstructured}}
	doc := map[string]any{
		"derived_struct_data": map[string]any{
			"content": "Fallback content",
		},
	}
	result := c.parseDocumentResult(doc)
	if result.PageContent != "Fallback content" {
		t.Errorf("PageContent = %q, want 'Fallback content'", result.PageContent)
	}
}

// --- simplifyResults ---

func TestSimplifyResults_Mixed(t *testing.T) {
	c := &Client{config: Config{EngineDataType: EngineDataTypeUnstructured}}
	resp := &SearchResponse{
		Results: []map[string]any{
			{"document": map[string]any{
				"id": "doc1",
				"derived_struct_data": map[string]any{"title": "Test"},
			}},
			{"chunk": map[string]any{
				"id":      "chunk1",
				"content": "Chunk content",
			}},
		},
	}
	c.simplifyResults(resp)

	if len(resp.SimplifiedResults) != 2 {
		t.Errorf("len(SimplifiedResults) = %d, want 2", len(resp.SimplifiedResults))
	}
}

// --- buildContentSearchSpec ---

func TestBuildContentSearchSpec_Snippets(t *testing.T) {
	c := &Client{config: Config{EngineChunkType: EngineChunkTypeDocumentWithSnippets}}
	spec := c.buildContentSearchSpec()
	if spec.SnippetSpec == nil || !spec.SnippetSpec.ReturnSnippet {
		t.Error("expected SnippetSpec.ReturnSnippet=true for DOCUMENT_WITH_SNIPPETS")
	}
	if spec.ExtractiveContentSpec != nil {
		t.Error("expected no ExtractiveContentSpec for DOCUMENT_WITH_SNIPPETS")
	}
}

func TestBuildContentSearchSpec_ExtractiveSegments(t *testing.T) {
	c := &Client{config: Config{
		EngineChunkType: EngineChunkTypeDocumentWithExtractiveSegments,
		SummaryType:     SummaryTypeVertexAISearch,
	}}
	spec := c.buildContentSearchSpec()
	if spec.SnippetSpec == nil || !spec.SnippetSpec.ReturnSnippet {
		t.Error("expected SnippetSpec.ReturnSnippet=true for DOCUMENT_WITH_EXTRACTIVE_SEGMENTS")
	}
	if spec.ExtractiveContentSpec == nil {
		t.Fatal("expected ExtractiveContentSpec for DOCUMENT_WITH_EXTRACTIVE_SEGMENTS")
	}
	if spec.ExtractiveContentSpec.MaxExtractiveAnswerCount != 1 {
		t.Errorf("MaxExtractiveAnswerCount = %d, want 1", spec.ExtractiveContentSpec.MaxExtractiveAnswerCount)
	}
	if spec.SummarySpec == nil {
		t.Fatal("expected SummarySpec for SummaryType=VERTEX_AI_SEARCH")
	}
	if spec.SummarySpec.SummaryResultCount != 5 {
		t.Errorf("SummaryResultCount = %d, want 5", spec.SummarySpec.SummaryResultCount)
	}
}

func TestBuildContentSearchSpec_NoSummary(t *testing.T) {
	c := &Client{config: Config{
		EngineChunkType: EngineChunkTypeChunk,
		SummaryType:     SummaryTypeNone,
	}}
	spec := c.buildContentSearchSpec()
	if spec.SummarySpec != nil {
		t.Error("expected no SummarySpec when SummaryType=NONE")
	}
}

// --- getEnv ---

func TestGetEnv_WithFallback(t *testing.T) {
	t.Setenv("TEST_KEY_ABSENT", "")
	got := getEnv("TEST_KEY_ABSENT", "fallback")
	if got != "fallback" {
		t.Errorf("getEnv() = %q, want 'fallback'", got)
	}
}

func TestGetEnv_WithValue(t *testing.T) {
	t.Setenv("TEST_KEY_PRESENT", "myvalue")
	got := getEnv("TEST_KEY_PRESENT", "fallback")
	if got != "myvalue" {
		t.Errorf("getEnv() = %q, want 'myvalue'", got)
	}
}

// contains is a test helper.
func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(sub) == 0 || strings.Contains(s, sub))
}
