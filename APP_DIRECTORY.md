# 📱 App Directory — Generative AI Repository

> A comprehensive index of every application, sample app, demo, and deployable project in this repository (fork of [GoogleCloudPlatform/generative-ai](https://github.com/GoogleCloudPlatform/generative-ai)).

---

## Table of Contents

1. [Audio Apps](#1-audio-apps)
2. [Gemini — Agent Apps](#2-gemini--agent-apps)
3. [Gemini — AutoCal](#3-gemini--autocal)
4. [Gemini — Computer Use](#4-gemini--computer-use)
5. [Gemini — Function Calling Apps](#5-gemini--function-calling-apps)
6. [Gemini — MCP Apps](#6-gemini--mcp-apps)
7. [Gemini — Multimodal Live API Apps](#7-gemini--multimodal-live-api-apps)
8. [Gemini — Sample Apps](#8-gemini--sample-apps)
9. [Gemini — Quickbot Templates](#9-gemini--quickbot-templates)
10. [Gemini — Use Cases & Cloud Functions](#10-gemini--use-cases--cloud-functions)
11. [Gemini — Evaluation & Tuning](#11-gemini--evaluation--tuning)
12. [Genkit Apps](#12-genkit-apps)
13. [Partner Model Apps](#13-partner-model-apps)
14. [Search Apps](#14-search-apps)
15. [Tools](#15-tools)
16. [Vision Apps](#16-vision-apps)
17. [Summary Statistics](#17-summary-statistics)
18. [Quick-Start Index](#18-quick-start-index)

---

## 1. Audio Apps

### Live Translator
| Field | Details |
|---|---|
| **Path** | `audio/speech/sample-apps/live-translator/` |
| **Tech Stack** | Python, Streamlit |
| **Deployment** | ✅ Docker / Cloud Run |
| **Key Dependencies** | Vertex AI Speech-to-Text, Cloud Translation API, Streamlit |

**Description:** A Cloud Run application that uses Streamlit to provide real-time speech translation. Users speak in one language and receive live text and audio output in another language.

---

## 2. Gemini — Agent Apps

### Always-On Memory Agent
| Field | Details |
|---|---|
| **Path** | `gemini/agents/always-on-memory-agent/` |
| **Tech Stack** | Python |
| **Deployment** | Local / Cloud Run |
| **Key Dependencies** | Google ADK, Gemini Flash-Lite, Python |

**Description:** An always-on AI agent with persistent memory built with Google ADK. Unlike typical agents that forget between sessions, this lightweight background process continuously reads, consolidates, and connects information — no vector database required.

---

### Gen AI Experience Concierge — LangGraph Demo
| Field | Details |
|---|---|
| **Path** | `gemini/agents/genai-experience-concierge/langgraph-demo/` |
| **Tech Stack** | Python (backend), Python/Streamlit (frontend) |
| **Deployment** | ✅ Docker (separate backend & frontend containers) |
| **Key Dependencies** | LangGraph, FastAPI, Streamlit, Vertex AI, Gemini |

**Description:** A collection of agent design patterns built with LangGraph, packaged as a deployable full-stack application. Demonstrates guardrail classifiers, semantic routing, function-calling streaming, and task-planner multi-agent patterns. The FastAPI server is compatible with the LangGraph Cloud API spec.

---

## 3. Gemini — AutoCal

### AutoCal
| Field | Details |
|---|---|
| **Path** | `gemini/autocal/` |
| **Tech Stack** | Node.js / Next.js (frontend), Python Cloud Function (backend) |
| **Deployment** | ✅ Cloud Functions + Firebase |
| **Key Dependencies** | Gemini 2.0 Flash, Firestore, Cloud Functions, Next.js, MUI |

**Description:** A screenshot-to-calendar web app powered by Gemini 2.0 Flash. Users upload screenshots and the app automatically extracts event information and adds events to their calendar. Uses Firestore triggers for event-driven architecture.

---

## 4. Gemini — Computer Use

### Web Agent
| Field | Details |
|---|---|
| **Path** | `gemini/computer-use/web-agent/` |
| **Tech Stack** | Python |
| **Deployment** | Local |
| **Key Dependencies** | Gemini Computer Use Tool, Playwright, Vertex AI |

**Description:** A Python script demonstrating a web automation agent powered by Gemini's Computer Use tool. Uses Playwright to control a real browser and perform multi-step web tasks based on natural language prompts.

---

## 5. Gemini — Function Calling Apps

### Function Calling Service
| Field | Details |
|---|---|
| **Path** | `gemini/function-calling/function_calling_service/` |
| **Tech Stack** | Python, Flask |
| **Deployment** | Cloud Run |
| **Key Dependencies** | Vertex AI, Gemini Function Calling, Flask |

**Description:** A Flask-based Cloud Run service that demonstrates how to migrate notebook-style function calling code into a production-ready REST API. Wraps an address lookup tool as a callable service.

---

### SQL Talk — Natural Language to BigQuery
| Field | Details |
|---|---|
| **Path** | `gemini/function-calling/sql-talk-app/` |
| **Tech Stack** | Python, Streamlit |
| **Deployment** | Cloud Run |
| **Key Dependencies** | Vertex AI, Gemini Function Calling, BigQuery, Streamlit |

**Description:** A Streamlit app that lets users query BigQuery databases using natural language instead of SQL. Uses Gemini's function calling to translate natural language questions into BigQuery API calls and return structured results.

---

## 6. Gemini — MCP Apps

### ADK MCP App
| Field | Details |
|---|---|
| **Path** | `gemini/mcp/adk_mcp_app/` |
| **Tech Stack** | Python |
| **Deployment** | Local |
| **Key Dependencies** | Google ADK, MCP (Model Context Protocol), Gemini |

**Description:** A web application demonstrating integration between Google ADK agents and the Model Context Protocol (MCP). Fetches cocktail information from TheCocktailDB API via a local MCP server, showing how ADK agents can use MCP as a data transport layer.

---

### ADK Multi-Agent MCP App
| Field | Details |
|---|---|
| **Path** | `gemini/mcp/adk_multiagent_mcp_app/` |
| **Tech Stack** | Python |
| **Deployment** | ✅ Docker / Cloud Run |
| **Key Dependencies** | Google ADK, MCP, Gemini, Multi-agent orchestration |

**Description:** A multi-agent system where a root agent coordinates specialized sub-agents (Cocktail and Booking agents) that interact with their respective MCP servers. Demonstrates hierarchical agent delegation and MCP-based tool discovery.

---

### MCP Orchestration App
| Field | Details |
|---|---|
| **Path** | `gemini/mcp/mcp_orchestration_app/` |
| **Tech Stack** | Python |
| **Deployment** | Local |
| **Key Dependencies** | Gemini API, MCP (Model Context Protocol) |

**Description:** Demonstrates a cost-effective MCP server-client architecture where affordable Gemini models act as orchestrators for specialized AI capabilities. Exposes Gemini's features as structured MCP "tools" for better composability and control.

---

## 7. Gemini — Multimodal Live API Apps

### Gradio Voice Chat
| Field | Details |
|---|---|
| **Path** | `gemini/multimodal-live-api/gradio-voice/` |
| **Tech Stack** | Python, Gradio |
| **Deployment** | Cloud Run |
| **Key Dependencies** | Gemini 2.0 Live API, Gradio, WebRTC |

**Description:** A low-latency bidirectional streaming voice chat application built with Gradio and Gemini. Enables real-time audio conversations via WebRTC with multiple voice options. Pure Python — no JavaScript required.

---

### PCM Audio Debugger
| Field | Details |
|---|---|
| **Path** | `gemini/multimodal-live-api/pcm-audio-debugger/` |
| **Tech Stack** | HTML / Vanilla JavaScript |
| **Deployment** | Static (open in browser) |
| **Key Dependencies** | None (zero dependencies) |

**Description:** A zero-dependency, single-file HTML tool for debugging PCM audio data. Combines a microphone recorder (converts to Base64 PCM) and a packet player (decodes and plays Base64 PCM). Supports multiple sample rates, mono/stereo, and 8/16/32-bit formats.

---

### Project Livewire
| Field | Details |
|---|---|
| **Path** | `gemini/multimodal-live-api/project-livewire/` |
| **Tech Stack** | Node.js (client), Python (server), Python Cloud Functions |
| **Deployment** | ✅ Docker (client + server) + Cloud Functions |
| **Key Dependencies** | Gemini 2.0 Flash Live API, WebSockets, Cloud Functions |

**Description:** A production-oriented real-time multimodal chat app — the "Star Trek computer" experience. Supports voice, webcam, screen sharing, and instant streamed audio responses. Features integrated weather and calendar tools via Cloud Functions. Built for production-readiness with a containerized client and server.

---

### WebSocket Demo App
| Field | Details |
|---|---|
| **Path** | `gemini/multimodal-live-api/websocket-demo-app/` |
| **Tech Stack** | Python (backend), HTML/JavaScript (frontend) |
| **Deployment** | ✅ Docker |
| **Key Dependencies** | Gemini 2.0 Multimodal Live API, WebSockets, Python |

**Description:** A tutorial-style web application demonstrating bidirectional WebSocket streaming with Gemini's Multimodal Live API. The Python backend handles authentication and proxying; the HTML/JS frontend provides voice and camera controls.

---

### Customer Support Demo App
| Field | Details |
|---|---|
| **Path** | `gemini/multimodal-live-api/native-audio-websocket-demo-apps/customer-support-demo-app/` |
| **Tech Stack** | React (frontend), Python (backend) |
| **Deployment** | Local / Cloud Run |
| **Key Dependencies** | Gemini Live API, React, Node.js, Python, WebSockets |

**Description:** A next-generation customer support agent demo using Gemini Live API. Features multimodal input (video/audio), emotion detection for affective dialogue, and real-time tool execution (process refunds, transfer to human). Showcases futuristic customer service interactions.

---

### Gaming Assistant Demo App
| Field | Details |
|---|---|
| **Path** | `gemini/multimodal-live-api/native-audio-websocket-demo-apps/gaming-assistant-demo-app/` |
| **Tech Stack** | React (frontend), Python (backend) |
| **Deployment** | Local / Cloud Run |
| **Key Dependencies** | Gemini Live API, React, Node.js, Python, WebSockets |

**Description:** A real-time AI gaming assistant that watches gameplay and listens to the player's voice to provide contextual tips and guides. Demonstrates proactive audio responses and persona-based prompting with a "gaming companion" personality.

---

### Plain JS Demo App
| Field | Details |
|---|---|
| **Path** | `gemini/multimodal-live-api/native-audio-websocket-demo-apps/plain-js-demo-app/` |
| **Tech Stack** | Vanilla JavaScript (frontend), Python (backend) |
| **Deployment** | Local |
| **Key Dependencies** | Gemini Live API, Python WebSocket proxy, Vanilla JS |

**Description:** A no-framework WebSocket client for Gemini Live API with audio and video streaming support. Demonstrates that you can integrate Gemini Live API without any JavaScript framework — just pure browser-native APIs.

---

### Plain JS + Python SDK Demo App
| Field | Details |
|---|---|
| **Path** | `gemini/multimodal-live-api/native-audio-websocket-demo-apps/plain-js-python-sdk-demo-app/` |
| **Tech Stack** | Vanilla JavaScript (frontend), Python/FastAPI (backend) |
| **Deployment** | Local |
| **Key Dependencies** | Google Gen AI Python SDK, FastAPI, Vanilla JS |

**Description:** Demonstrates using the Google Gen AI Python SDK as the backend for a Gemini Live API application, with vanilla JS on the frontend. Showcases how to build a robust Python backend that handles the API connection securely.

---

### React Demo App
| Field | Details |
|---|---|
| **Path** | `gemini/multimodal-live-api/native-audio-websocket-demo-apps/react-demo-app/` |
| **Tech Stack** | React (frontend), Python (backend proxy) |
| **Deployment** | Local |
| **Key Dependencies** | Gemini Live API, React, Python WebSocket proxy |

**Description:** A React-based client for Gemini Live API featuring real-time audio and video streaming. Uses a Python WebSocket proxy for secure authentication and to hide API credentials from the browser.

---

### Real-time Advisor Demo App
| Field | Details |
|---|---|
| **Path** | `gemini/multimodal-live-api/native-audio-websocket-demo-apps/realtime-advisor-demo-app/` |
| **Tech Stack** | React (frontend), Python (backend) |
| **Deployment** | Local / Cloud Run |
| **Key Dependencies** | Gemini Live API, React, Node.js, Python, WebSockets |

**Description:** A real-time AI business advisor that listens to conversations and proactively provides relevant insights from a knowledge base. Supports two modes: silent visual mode (modal popups) and active audio mode (speaks insights aloud). Demonstrates function calling and precise interruption control.

---

## 8. Gemini — Sample Apps

### Accelerating Product Innovation
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/accelerating_product_innovation/` |
| **Tech Stack** | Python, Streamlit, Cloud Functions |
| **Deployment** | Cloud Run + Cloud Functions |
| **Key Dependencies** | Vertex AI (Gemini, Imagen), Cloud Functions, Streamlit, BigQuery |

**Description:** A Streamlit platform for retail product managers and R&D analysts to rapidly generate new product concepts using Generative AI. Accelerates new product development by generating bulk product ideas, addressing market trends, and assisting with regulatory compliance.

---

### Finance Advisor Spanner (Finvest)
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/finance-advisor-spanner/` |
| **Tech Stack** | Python |
| **Deployment** | ✅ Docker / Cloud Run |
| **Key Dependencies** | Cloud Spanner (relational + text + vector + graph search), Vertex AI, Gemini |

**Description:** A financial advisory demo app showcasing Cloud Spanner's multi-model capabilities. Combines relational search, full-text/fuzzy search, KNN vector search (for "socially responsible" funds), and graph traversal (fund-of-funds exposure analysis) in a single platform.

---

### Gemini HallCheck
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/gemini-hallcheck/` |
| **Tech Stack** | Python |
| **Deployment** | Local (CLI tool) |
| **Key Dependencies** | Gemini API / Vertex AI, Hugging Face (MMLU), google-genai |

**Description:** A confidence-targeted, abstention-aware hallucination evaluator for Gemini. Implements the "answer only if > t confident; otherwise say IDK" framework. Produces risk-coverage curves, supports MMLU benchmarking, uses an LLM semantic judge, and includes async evaluation with quota-aware retries.

---

### Gemini Live Telephony App (AI Care Assistant)
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/gemini-live-telephony-app/` |
| **Tech Stack** | Python, FastAPI |
| **Deployment** | ✅ Docker / Cloud Run |
| **Key Dependencies** | Twilio, Gemini Live API, FastAPI, python-samplerate, Cloud Run |

**Description:** A real-time telephony voice app integrating Twilio for phone calls with Gemini Live API for conversational AI. Features high-quality streaming audio resampling (`python-samplerate`), low cold-start Cloud Run deployment (`min-instances=1`), and session affinity for in-memory DSP state.

---

### Gemini Mesop Cloud Run
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/gemini-mesop-cloudrun/` |
| **Tech Stack** | Python, Mesop |
| **Deployment** | Cloud Run |
| **Key Dependencies** | Vertex AI, Gemini API, Mesop framework |

**Description:** A Cloud Run application using the Mesop framework to build a Gemini-powered UI. Demonstrates an alternative to Streamlit/Gradio for building Python-first web interfaces for generative AI applications.

---

### Gemini Quart Cloud Run
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/gemini-quart-cloudrun/` |
| **Tech Stack** | Python, Quart (async Flask) |
| **Deployment** | ✅ Docker / Cloud Run |
| **Key Dependencies** | Vertex AI, Gemini Live API, Quart, Cloud Run |

**Description:** A non-blocking chat app using Quart (async Flask) with the Gemini Live API on Cloud Run. Demonstrates asynchronous streaming with proper interruption handling in a Python web server context.

---

### Gemini Streamlit Cloud Run
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/gemini-streamlit-cloudrun/` |
| **Tech Stack** | Python, Streamlit |
| **Deployment** | ✅ Docker / Cloud Run |
| **Key Dependencies** | Vertex AI, Gemini API, Streamlit |

**Description:** A Cloud Run ready Streamlit application demonstrating Gemini API integration. A reference template for building and deploying Streamlit-based generative AI applications on Google Cloud with a single-click deploy button.

---

### GenWealth — Financial Advisory Platform
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/genwealth/` |
| **Tech Stack** | Node.js/Express (API), React (UI), Python (Cloud Functions) |
| **Deployment** | ✅ Docker (multi-component) + Cloud Functions |
| **Key Dependencies** | AlloyDB AI, Vertex AI (Gemini + embeddings), Cloud Run, Cloud Functions, React, Express |

**Description:** A comprehensive full-stack financial advisory platform for a fictional firm "GenWealth". Adds three Gen AI features to an investment advisory app: semantic search (AlloyDB AI vector similarity), customer segmentation (RAG), and an AI chatbot grounded in application data. Showcases AlloyDB AI + Vertex AI integration.

---

### Image Bash Jam
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/image-bash-jam/` |
| **Tech Stack** | Bash |
| **Deployment** | Local / Cloud Shell |
| **Key Dependencies** | Gemini API (via curl/gcloud), Bash |

**Description:** Demonstrates calling the Gemini API directly from Bash shell scripts to analyze and describe images. A minimal, no-framework example showing how to integrate Gemini into shell automation workflows.

---

### LlamaDeploy on Cloud Run
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/llamadeploy-on-cloud-run/` |
| **Tech Stack** | Python |
| **Deployment** | ✅ Docker / Cloud Run |
| **Key Dependencies** | LlamaIndex, llama-deploy, Gemini, Firestore |

**Description:** A LlamaIndex Workflow application deployed via the `llama-deploy` library on Cloud Run. Implements a complex RAG workflow using Gemini models and Firestore as the vector/document store.

---

### LlamaIndex Advanced Agentic RAG
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/llamaindex-rag/` |
| **Tech Stack** | Python (FastAPI backend), Python/Streamlit (UI) |
| **Deployment** | ✅ Docker (multi-component) |
| **Key Dependencies** | LlamaIndex, Vertex AI (Gemini, Vector Search), FastAPI, Streamlit |

**Description:** An advanced Retrieval-Augmented Generation system for rapid experimentation with different indexing strategies, retrieval algorithms, and LLMs. Features a FastAPI backend for query processing, Streamlit frontend, and Vertex AI Vector Search. Includes evaluation metrics and deployment on Google Cloud.

---

### Photo Discovery (Vertex AI Search + Flutter)
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/photo-discovery/` |
| **Tech Stack** | Python/Flask (web), Flutter/Dart (mobile) |
| **Deployment** | ✅ Docker (web component) |
| **Key Dependencies** | Vertex AI Search, Agent Engine (Reasoning Engine), Flutter, Python |

**Description:** A multi-platform demo (Flutter mobile + web) that lets users take or upload a photo of a landmark or Google merchandise item. An AI agent identifies the subject, provides descriptions, and answers follow-up questions grounded in Wikipedia and merchandise catalog data.

---

### SWOT Agent
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/swot-agent/` |
| **Tech Stack** | Python, FastAPI, HTMX, Tailwind CSS |
| **Deployment** | ✅ Docker / Cloud Run |
| **Key Dependencies** | Gemini 2.0 Flash, Pydantic AI, FastAPI, HTMX |

**Description:** A web application performing automated SWOT analysis (Strengths, Weaknesses, Opportunities, Threats) using Gemini 2.0 Flash and the Pydantic AI agent framework. Delivers structured strategic analysis reports through a clean HTMX + Tailwind UI.

---

### E2E Gen AI App Starter Pack *(Relocated)*
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/e2e-gen-ai-app-starter-pack/` |
| **Note** | ⚠️ Moved to [agent-starter-pack](https://github.com/GoogleCloudPlatform/agent-starter-pack) |

---

## 9. Gemini — Quickbot Templates

> **Quickbot** is a no-code platform for deploying AI agent applications on Google Cloud Platform. Each template follows the same architecture: **Angular frontend** + **Python/FastAPI backend**, fully containerized with Docker Compose.

### Conversational App — Multi-Playbook
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/quickbot/conversational-app-multi-playbook/` |
| **Tech Stack** | Angular (frontend), Python/FastAPI (backend), Cloud Functions |
| **Deployment** | ✅ Docker Compose |
| **Key Dependencies** | Vertex AI Conversation (Agent Builder), Gemini, Cloud Functions (RAG), Angular |

**Description:** A full-stack conversational chatbot that routes user intents to multiple specialized playbooks. Uses a custom Cloud Function for RAG and Vertex AI Agent Builder for multi-playbook conversation orchestration.

---

### Conversational App — Single Playbook
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/quickbot/conversational-app-single-playbook/` |
| **Tech Stack** | Angular (frontend), Python/FastAPI (backend), Cloud Functions |
| **Deployment** | ✅ Docker Compose |
| **Key Dependencies** | Vertex AI Conversation (Agent Builder), Gemini, Cloud Functions (RAG), Angular |

**Description:** A single-intent conversational agent backed by one configurable playbook. Uses a custom Cloud Function for RAG and Vertex AI Agent Builder for conversation management.

---

### Document Search — Agent Builder
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/quickbot/document-search-using-agent-builder/` |
| **Tech Stack** | Angular (frontend), Python/FastAPI (backend) |
| **Deployment** | ✅ Docker Compose |
| **Key Dependencies** | Vertex AI Search (Agent Builder), Gemini, Angular |

**Description:** A document search application powered by Vertex AI Search (Agent Builder). Users can search through indexed documents and receive AI-generated answers grounded in document content.

---

### Image Background Changer — Imagen3
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/quickbot/image-background-changer-using-imagen3/` |
| **Tech Stack** | Angular (frontend), Python/FastAPI (backend) |
| **Deployment** | ✅ Docker Compose |
| **Key Dependencies** | Vertex AI (Imagen3), Angular, FastAPI |

**Description:** An image editing application that uses Imagen3 on Vertex AI to replace or modify image backgrounds based on text prompts. Provides a user-friendly interface for AI-powered image manipulation.

---

### LinkedIn Profile Image Generation — Imagen3
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/quickbot/linkedin-profile-image-generation-using-imagen3/` |
| **Tech Stack** | Angular (frontend), Python/FastAPI (backend) |
| **Deployment** | ✅ Docker Compose |
| **Key Dependencies** | Vertex AI (Imagen3), Angular, FastAPI |

**Description:** Generates professional LinkedIn post images by combining text prompts with pre-designed templates, powered by Imagen3 on Vertex AI. Helps users create eye-catching professional profile images.

---

### Multi-Agent Travel Concierge — ADK + Agent Engine
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/quickbot/multi-agent-travel-concierge-with-adk/` |
| **Tech Stack** | Angular (frontend), Python/FastAPI (backend) |
| **Deployment** | ✅ Docker Compose |
| **Key Dependencies** | Google ADK, Agent Engine (Vertex AI), Gemini, Angular |

**Description:** A sophisticated multi-agent travel planning system using Google ADK and Vertex AI Agent Engine. Orchestrates specialized agents for flights, accommodations, local activities, and real-time alerts. Based on the official ADK Travel Concierge sample.

---

### Text-to-Image — Imagen3
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/quickbot/text-to-image-using-imagen3/` |
| **Tech Stack** | Angular (frontend), Python/FastAPI (backend) |
| **Deployment** | ✅ Docker Compose |
| **Key Dependencies** | Vertex AI (Imagen3), Angular, FastAPI |

**Description:** Generates high-quality images from text prompts using Imagen3 on Vertex AI. A clean frontend allows users to enter prompts and receive generated images in real time.

---

### Website Search — Custom RAG
| Field | Details |
|---|---|
| **Path** | `gemini/sample-apps/quickbot/website-search-using-agent-builder/` |
| **Tech Stack** | Angular (frontend), Python/FastAPI (backend), Cloud Functions |
| **Deployment** | ✅ Docker Compose |
| **Key Dependencies** | Vertex AI Search, Cloud Functions (embeddings/indexing), Angular |

**Description:** An intelligent website search application using custom RAG via Cloud Functions. Indexes website content, generates embeddings, and provides AI-powered search with grounded answers.

---

## 10. Gemini — Use Cases & Cloud Functions

### Entity Extraction Service
| Field | Details |
|---|---|
| **Path** | `gemini/use-cases/entity-extraction/` |
| **Tech Stack** | Python, Flask |
| **Deployment** | ✅ Docker / Cloud Run |
| **Key Dependencies** | Vertex AI, Gemini, Flask, Cloud Run |

**Description:** A Flask-based Cloud Run service for classifying and extracting structured information from varied documents using the Gemini family of models. Ideal for documents with variable structure that traditional parsers cannot handle reliably.

---

### Document AI + Gemini Entity Extraction
| Field | Details |
|---|---|
| **Path** | `gemini/use-cases/applying-llms-to-data/gemini-and-documentai-for-entity-extraction/` |
| **Tech Stack** | Python |
| **Deployment** | Local / Cloud |
| **Key Dependencies** | Google Cloud Document AI, Vertex AI, Gemini API |

**Description:** A Python script comparing Document AI and Gemini API for extracting entities from PDF documents. Demonstrates both online and batch processing requests with side-by-side output comparison to help users choose the right tool.

---

### Gemini with BigQuery Remote Functions
| Field | Details |
|---|---|
| **Path** | `gemini/use-cases/applying-llms-to-data/using-gemini-with-bigquery-remote-functions/` |
| **Tech Stack** | Python, Cloud Functions |
| **Deployment** | Cloud Functions |
| **Key Dependencies** | BigQuery Remote Functions, Vertex AI, Gemini API, Cloud Functions |

**Description:** Cloud Functions that expose Gemini image and text analysis as BigQuery Remote Functions. Enables analysts to call Gemini directly from SQL queries to analyze images and text stored in BigQuery.

---

### Multimodal Data Curation Pipeline
| Field | Details |
|---|---|
| **Path** | `gemini/use-cases/multimodal-data-curation/` |
| **Tech Stack** | Python |
| **Deployment** | Serverless (Google Cloud) |
| **Key Dependencies** | Vertex AI, Gemini, Cloud Storage, Serverless pipeline |

**Description:** A fully serverless data curation pipeline for building Text-to-Video (T2V) training datasets at scale. Handles large-scale multimodal data processing, annotation, and quality filtering on Google Cloud.

---

## 11. Gemini — Evaluation & Tuning

### Synthetic Data Evals — Agent Evaluation Framework
| Field | Details |
|---|---|
| **Path** | `gemini/evaluation/synthetic-data-evals/` |
| **Tech Stack** | Python |
| **Deployment** | Local |
| **Key Dependencies** | Vertex AI, Gemini, Python (pyproject.toml) |

**Description:** A framework for generating synthetic evaluation datasets and testing LLM-powered agents in customer support scenarios. Includes a customer support agent, a dataset generator, and an evaluation harness for measuring agent performance.

---

### Gen AI MLOps Tune and Eval
| Field | Details |
|---|---|
| **Path** | `gemini/tuning/genai-mlops-tune-and-eval/` |
| **Tech Stack** | Python |
| **Deployment** | Vertex AI Pipelines |
| **Key Dependencies** | Vertex AI Pipelines, Gemini tuning, Model evaluation |

**Description:** A tutorial using Vertex AI Pipelines to automate the full cycle of LLM tuning and evaluation. Demonstrates tuning a model for medical data summarization (glucose readings), then automatically evaluating it against a previously tuned baseline.

---

## 12. Genkit Apps

### Generate Synthetic Database
| Field | Details |
|---|---|
| **Path** | `genkit/generate-synthetic-database/` |
| **Tech Stack** | Node.js, Cloud Functions |
| **Deployment** | Cloud Functions |
| **Key Dependencies** | Firebase Genkit, Gemini 2.0, Firestore |

**Description:** Uses Firebase Genkit's structured output feature with Gemini 2.0 to generate realistic synthetic sales data for a fictional dog food company. Data is stored row-by-row in Firestore, demonstrating Genkit's structured generation capabilities for database seeding.

---

## 13. Partner Model Apps

### Anthropic Claude Computer Use Demo
| Field | Details |
|---|---|
| **Path** | `partner-models/claude/computer-use-demo/` |
| **Tech Stack** | Python (inferred) |
| **Deployment** | Cloud Shell / VM |
| **Key Dependencies** | Anthropic Claude API (via Vertex AI), Computer Use beta feature |

**Description:** A demo of Anthropic's Claude computer use capability running on Google Cloud. Allows Claude to control a virtual desktop and perform tasks autonomously. Includes best-practice security guidance (isolated VMs, restricted permissions, allowlisted domains).

---

## 14. Search Apps

### Auto RAG Eval
| Field | Details |
|---|---|
| **Path** | `search/auto-rag-eval/` |
| **Tech Stack** | Python |
| **Deployment** | Local |
| **Key Dependencies** | Vertex AI, Gemini, RAG evaluation frameworks |

**Description:** An automated benchmark generation tool for RAG (Retrieval-Augmented Generation) systems. Generates realistic evaluation datasets and measures RAG pipeline quality metrics such as faithfulness and relevance.

---

### Vertex AI Search — Cloud Function
| Field | Details |
|---|---|
| **Path** | `search/cloud-function/python/` |
| **Tech Stack** | Python, Cloud Functions |
| **Deployment** | Cloud Functions (HTTPS trigger) |
| **Key Dependencies** | Vertex AI Search Python client, Cloud Functions |

**Description:** A Python Cloud Function exposing Vertex AI Search as a REST API. Returns search results, snippets, metadata, and LLM-generated summaries grounded in search results. A ready-to-deploy serverless search backend.

---

### Vertex AI Search Web App Demo
| Field | Details |
|---|---|
| **Path** | `search/web-app/` |
| **Tech Stack** | Python |
| **Deployment** | ✅ Docker / Cloud Run |
| **Key Dependencies** | Vertex AI Search, Enterprise Knowledge Graph API, Python |

**Description:** A web app demonstrating full-text, semantic, and Knowledge Graph search using Vertex AI Search (formerly Enterprise Search). Includes document corpus search and public Cloud Knowledge Graph queries.

---

### Gemini Box Office — License Manager
| Field | Details |
|---|---|
| **Path** | `search/gemini-enterprise/group-licensing/` |
| **Tech Stack** | Python |
| **Deployment** | ✅ Docker / Cloud Run Jobs |
| **Key Dependencies** | Cloud Identity API, Discovery Engine API, Cloud Run Jobs, Cloud Scheduler |

**Description:** A Cloud Run Job that automates Gemini Enterprise license lifecycle management. Bridges Google Cloud Identity group membership and Discovery Engine license assignment. Runs two scheduled jobs: a daily "joiner" (provision licenses) and a 6-hourly "garbage collector" (revoke stale licenses).

---

### VAIS Building Blocks *(Notebook Collection)*
| Field | Details |
|---|---|
| **Path** | `search/vais-building-blocks/` |
| **Tech Stack** | Python (Jupyter Notebooks) |
| **Deployment** | Notebooks |
| **Key Dependencies** | Vertex AI Search, Python |

**Description:** A series of Jupyter notebooks demonstrating individual Vertex AI Search (VAIS) features and components — structured data search, unstructured document search, website search, and third-party data source integration.

---

## 15. Tools

### LLM EvalKit
| Field | Details |
|---|---|
| **Path** | `tools/llmevalkit/` |
| **Tech Stack** | Python, Streamlit |
| **Deployment** | Local |
| **Key Dependencies** | Vertex AI, Gemini, Python evaluation frameworks |

**Description:** A comprehensive toolkit for evaluating and improving LLM performance on specific tasks. Supports human and model-based evaluation, prompt management, dataset versioning, and automated prompt optimization. Streamlines the prompt engineering and evaluation iteration cycle.

---

## 16. Vision Apps

### V-Start — Veo Prompting & Evaluation Toolkit
| Field | Details |
|---|---|
| **Path** | `vision/sample-apps/V-Start/` |
| **Tech Stack** | Node.js |
| **Deployment** | ✅ Docker |
| **Key Dependencies** | Veo (video generation), Vertex AI, Node.js |

**Description:** An experimental toolkit that helps users create effective prompts for Veo (Google's video generation model) and evaluate how well generated videos align with the intended prompts. Split into two tools: a Prompting tool and an Evaluation tool.

---

### GenMedia Live
| Field | Details |
|---|---|
| **Path** | `vision/sample-apps/genmedia-live/` |
| **Tech Stack** | Python, Streamlit |
| **Deployment** | Local / Cloud Run |
| **Key Dependencies** | Gemini Live API, Imagen, Veo, Vertex AI |

**Description:** Extends the Gemini Live API with image and video generation tools. Enables real-time multimodal creation via voice or text: generate/edit images, create videos from text, animate images, create short movies, and answer questions about camera input — all through function calling.

---

### HEY_LLM / IMAGEN — Google Sheets Functions
| Field | Details |
|---|---|
| **Path** | `vision/use-cases/hey_llm/` |
| **Tech Stack** | Google Apps Script / JavaScript |
| **Deployment** | Google Sheets (Apps Script) |
| **Key Dependencies** | Gemini API, Imagen, Google Apps Script |

**Description:** Custom Google Sheets functions that bring LLMs and image generation into spreadsheets. `HEY_LLM` calls Gemini for natural language inferences over cell values; `IMAGEN` generates images from text prompts directly in spreadsheet cells.

---

## 17. Summary Statistics

### Total Apps Found: **65**

### Breakdown by Category
| Category | Count |
|---|---|
| Gemini — Quickbot Templates | 8 |
| Gemini — Multimodal Live API Apps | 10 |
| Gemini — Sample Apps | 12 |
| Gemini — Use Cases & Cloud Functions | 4 |
| Gemini — Agent Apps | 2 |
| Gemini — MCP Apps | 3 |
| Gemini — Evaluation & Tuning | 2 |
| Gemini — Function Calling Apps | 2 |
| Gemini — AutoCal | 1 |
| Gemini — Computer Use | 1 |
| Audio Apps | 1 |
| Genkit Apps | 1 |
| Partner Model Apps | 1 |
| Search Apps | 4 |
| Tools | 1 |
| Vision Apps | 3 |
| Notebook Collections *(not standalone apps)* | 9+ |

### Breakdown by Primary Tech Stack
| Tech Stack | Count |
|---|---|
| Python | 40+ |
| Node.js / JavaScript | 12+ |
| React | 5 |
| Angular | 8 (all Quickbot templates) |
| Python + Streamlit | 7 |
| Python + FastAPI | 10+ |
| Python + Flask | 3 |
| Bash | 1 |
| Google Apps Script | 1 |
| Flutter/Dart | 1 |

### Deployment Readiness
| Status | Count |
|---|---|
| ✅ Docker / Fully Containerized | 35+ |
| ✅ Cloud Run Compatible | 30+ |
| ✅ Cloud Functions | 8 |
| ✅ Cloud Run Jobs | 1 |
| 🔷 Local / Notebook Only | 15+ |

---

## 18. Quick-Start Index

### 🟢 Beginner-Friendly (Simple, Few Prerequisites)

| App | Path | Why It's Beginner-Friendly |
|---|---|---|
| **PCM Audio Debugger** | `gemini/multimodal-live-api/pcm-audio-debugger/` | Zero dependencies — open HTML file in browser |
| **Image Bash Jam** | `gemini/sample-apps/image-bash-jam/` | Just Bash + gcloud, no code required |
| **Gemini Streamlit Cloud Run** | `gemini/sample-apps/gemini-streamlit-cloudrun/` | One-click deploy button, minimal configuration |
| **SQL Talk App** | `gemini/function-calling/sql-talk-app/` | Clear tutorial app, single Python file |
| **Plain JS Demo App** | `gemini/multimodal-live-api/native-audio-websocket-demo-apps/plain-js-demo-app/` | No framework, just HTML + Python |
| **Gradio Voice Chat** | `gemini/multimodal-live-api/gradio-voice/` | Pure Python, minimal setup |
| **HEY_LLM** | `vision/use-cases/hey_llm/` | Runs inside Google Sheets |
| **Function Calling Service** | `gemini/function-calling/function_calling_service/` | Clear migration pattern from notebook |

---

### 🟡 Intermediate (Requires Cloud Setup, Some Configuration)

| App | Path | Complexity Notes |
|---|---|---|
| **WebSocket Demo App** | `gemini/multimodal-live-api/websocket-demo-app/` | Requires GCP auth, Docker |
| **SWOT Agent** | `gemini/sample-apps/swot-agent/` | Single container, clean FastAPI app |
| **Gemini Mesop Cloud Run** | `gemini/sample-apps/gemini-mesop-cloudrun/` | Single container, Cloud Run deploy |
| **Gemini Quart Cloud Run** | `gemini/sample-apps/gemini-quart-cloudrun/` | Single container, async patterns |
| **Live Translator** | `audio/speech/sample-apps/live-translator/` | Requires Speech/Translation API setup |
| **ADK MCP App** | `gemini/mcp/adk_mcp_app/` | Requires ADK + local MCP server |
| **VAIS Web App Demo** | `search/web-app/` | Requires Vertex AI Search index |
| **Auto RAG Eval** | `search/auto-rag-eval/` | Requires GCS bucket and search index |
| **Entity Extraction** | `gemini/use-cases/entity-extraction/` | Single container Flask app |
| **Always-On Memory Agent** | `gemini/agents/always-on-memory-agent/` | Requires ADK setup |

---

### 🔴 Advanced / Complex (Multi-Component, Multiple Cloud Services)

| App | Path | Complexity Notes |
|---|---|---|
| **GenWealth** | `gemini/sample-apps/genwealth/` | AlloyDB + Cloud Functions + React + Node.js API + Docker |
| **Project Livewire** | `gemini/multimodal-live-api/project-livewire/` | 2 containers + Cloud Functions + WebRTC |
| **LlamaIndex Advanced Agentic RAG** | `gemini/sample-apps/llamaindex-rag/` | FastAPI + Streamlit + Vertex AI Vector Search |
| **Gen AI Experience Concierge** | `gemini/agents/genai-experience-concierge/` | LangGraph + FastAPI + Streamlit + 2 containers |
| **Multi-Agent Travel Concierge** | `gemini/sample-apps/quickbot/multi-agent-travel-concierge-with-adk/` | ADK + Agent Engine + Angular + FastAPI |
| **AutoCal** | `gemini/autocal/` | Next.js + Cloud Functions + Firestore + Gemini |
| **Finance Advisor Spanner** | `gemini/sample-apps/finance-advisor-spanner/` | Cloud Spanner multi-model (relational + vector + graph + text) |
| **ADK Multi-Agent MCP App** | `gemini/mcp/adk_multiagent_mcp_app/` | Multiple agents + multiple MCP servers |
| **Gemini Live Telephony App** | `gemini/sample-apps/gemini-live-telephony-app/` | Twilio + Gemini Live API + audio DSP |
| **Photo Discovery** | `gemini/sample-apps/photo-discovery/` | Flutter (mobile) + Reasoning Engine + Vertex AI Search |
| **Quickbot Templates (×8)** | `gemini/sample-apps/quickbot/*/` | Angular + FastAPI + Docker Compose + Cloud Functions |
| **Gemini Box Office** | `search/gemini-enterprise/group-licensing/` | Cloud Run Jobs + Cloud Scheduler + Identity APIs |
| **Accelerating Product Innovation** | `gemini/sample-apps/accelerating_product_innovation/` | Streamlit + Cloud Functions (×3) + BigQuery |

---

*Last updated: 2026-05-04 | Repository: [agha64113-creator/generative-ai](https://github.com/agha64113-creator/generative-ai)*
