# Vertex AI Search Cloud Function — Go

This directory contains a Go implementation of a Google Cloud Function that
wraps the [Vertex AI Search](https://cloud.google.com/generative-ai-app-builder/docs/enterprise-search-introduction)
API as an HTTP endpoint.

It mirrors the functionality of the [Python version](../python/README.md) and
supports the same configuration options and data types.

## Prerequisites

1. A Google Cloud project with the Discovery Engine API enabled.
2. A configured Vertex AI Search data store. See
   [Get started with generic search](https://cloud.google.com/generative-ai-app-builder/docs/try-enterprise-search).
3. [Go 1.21+](https://go.dev/dl/) installed locally.

## Configuration

Copy `.env.example` to `.env` and fill in your values:

```
PROJECT_ID=your-project-id
LOCATION=global
DATA_STORE_ID=your-data-store-id
ENGINE_DATA_TYPE=UNSTRUCTURED
ENGINE_CHUNK_TYPE=DOCUMENT_WITH_EXTRACTIVE_SEGMENTS
SUMMARY_TYPE=VERTEX_AI_SEARCH
```

### Supported values

| Variable            | Options                                                                                          |
|---------------------|--------------------------------------------------------------------------------------------------|
| `ENGINE_DATA_TYPE`  | `UNSTRUCTURED` \| `STRUCTURED` \| `WEBSITE` \| `BLENDED`                                        |
| `ENGINE_CHUNK_TYPE` | `DOCUMENT_WITH_SNIPPETS` \| `DOCUMENT_WITH_EXTRACTIVE_SEGMENTS` \| `CHUNK` \| `NONE`            |
| `SUMMARY_TYPE`      | `NONE` \| `VERTEX_AI_SEARCH`                                                                     |

## Running locally

```bash
# Install the functions-framework CLI
go install github.com/GoogleCloudPlatform/functions-framework-go/funcframework/cmd/funcframework@latest

# Export environment variables
export $(cat .env | xargs)

# Start the local server
funcframework --target VertexAISearch --port 8080
```

Then send a request:

```bash
curl -X POST http://localhost:8080 \
  -H "Content-Type: application/json" \
  -d '{"search_term": "your search query"}'
```

Or via query parameter:

```bash
curl "http://localhost:8080?search_term=your+search+query"
```

## Running tests

### Unit tests

```bash
go test -v -run 'Test[^I]' ./...
```

### Integration tests

Integration tests call the real Vertex AI Search API and require the
environment variables from `.env` to be set:

```bash
export $(cat .env | xargs)
go test -v -run TestIntegration ./...
```

## Deploying to Cloud Functions

```bash
gcloud functions deploy vertex-ai-search \
  --gen2 \
  --runtime go121 \
  --region us-central1 \
  --source . \
  --entry-point VertexAISearch \
  --trigger-http \
  --allow-unauthenticated \
  --set-env-vars PROJECT_ID=$PROJECT_ID,DATA_STORE_ID=$DATA_STORE_ID,LOCATION=global
```
