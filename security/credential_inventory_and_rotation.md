# Credential Inventory and Rotation Plan

Date: 2026-05-19 (UTC)

## 1) Inventory of credentials referenced by app modules

| Service | Credential/env var | Example modules |
|---|---|---|
| Gemini / Google GenAI | `GOOGLE_API_KEY` | `gemini/mcp/mcp_orchestration_app/src/gemini_client.py`, `gemini/multimodal-live-api/project-livewire/server/config/config.py`, `gemini/sample-apps/gemini-streamlit-cloudrun/app.py` |
| Gemini (legacy naming) | `GEMINI_API_KEY` | `gemini/sample-apps/gemini-quart-cloudrun/app/app.py`, `gemini/sample-apps/gemini-live-telephony-app/main.py` |
| Google Cloud auth | `GOOGLE_APPLICATION_CREDENTIALS` | `gemini/sample-apps/swot-agent/agent.py`, `gemini/mcp/mcp_orchestration_app/example.env` |
| Vertex toggle | `GOOGLE_GENAI_USE_VERTEXAI` | `gemini/multimodal-live-api/project-livewire/server/config/config.py`, `gemini/mcp/adk_mcp_app/README.md` |
| Voice provider options | model voice config values (`VOICE_*`) | `gemini/multimodal-live-api/project-livewire/server/config/config.py` |

No active OpenAI-specific, Google Drive-specific, or third-party voice API key variables (e.g., `OPENAI_API_KEY`, `DEEPGRAM_API_KEY`, `ELEVENLABS_API_KEY`, `ASSEMBLYAI_API_KEY`) were found in the scoped app modules during this pass.

## 2) Rotation + revocation runbook (mandatory)

Because this repository cannot revoke provider credentials directly, perform these steps immediately in provider consoles:

1. Create replacement keys/service-account credentials for each live secret above.
2. Update secret manager entries / CI env vars / runtime env vars.
3. Redeploy every affected service.
4. Validate smoke tests and auth paths.
5. Revoke or disable previous keys immediately.

## 3) Key restrictions and quota ceilings

Apply for each provider where supported:
- Restrict by application origin/domain for browser-exposed keys.
- Restrict by source IP/CIDR for server-side keys.
- Restrict by API scope/service.
- Configure per-key quota ceilings and alerting.

## 4) Validation checklist

- [ ] New keys deployed to all environments.
- [ ] Old keys revoked.
- [ ] API calls succeed with new keys.
- [ ] 401/403 checks confirm old keys invalid.
- [ ] Budget/quota alerts configured.
