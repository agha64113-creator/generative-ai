# Security Incident Log: Credential Rotation

- Incident date: 2026-05-19 (UTC)
- Trigger: Requested emergency credential hygiene hardening.

## Affected services
- Gemini / Google GenAI API-key consumers
- Vertex/ADC service-account consumers

## Rotation execution record
- Rotation started: 2026-05-19T00:00:00Z (to be updated by operator)
- Rotation completed: PENDING OPERATOR ACTION
- Old key revocation completed: PENDING OPERATOR ACTION

## Post-rotation validation
- Secret scanning controls added in repo CI + pre-commit.
- Provider-side live key rotation/revocation still required and must be executed in cloud/provider consoles.
- Validation status after operator action: PENDING
