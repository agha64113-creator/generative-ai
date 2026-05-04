# Repository Security Audit Analysis

**Repository:** `agha64113-creator/generative-ai`  
**Base:** Fork of `GoogleCloudPlatform/generative-ai`  
**Date:** 2026-05-04  
**Scope:** Full repository scan for security vulnerabilities, malicious code, backdoors, data exfiltration, and unsafe coding patterns.

---

## Executive Summary

A repository-wide security audit was conducted covering all Python, JavaScript, and TypeScript source files. Two actionable security vulnerabilities were identified and remediated. No malicious code, spyware, backdoors, persistence mechanisms, data-exfiltration payloads, or obfuscated code was found. The repository is a standard fork of the Google Cloud Platform generative-AI sample collection.

---

## Findings

### CRITICAL — Code Injection via `eval()` on Untrusted LLM Output

| Field | Detail |
|---|---|
| **Severity** | Critical |
| **CWE** | CWE-95 (Improper Neutralization of Directives in Dynamically Evaluated Code) |
| **File** | `gemini/agents/research-multi-agents/ev_agent/agent_handler/agent_02_PlanningAgent.py` |
| **Lines** | 273–274 (before fix) |

**Description:**  
`_determine_visualization_requirement()` returns `response.text` — the raw string produced by the Gemini `generate_content()` API. This untrusted string was passed directly to Python's built-in `eval()`:

```python
# VULNERABLE (before fix)
print("Type of Visual toggle:", type(eval(needs_visualization)))
if eval(needs_visualization):
```

Because `eval()` executes arbitrary Python expressions, a compromised or adversarially prompted model could return a payload such as `__import__('os').system('...')`, achieving remote code execution in the process running the agent.

**Remediation applied:**  
Replaced both `eval()` calls with an explicit boolean/string comparison:

```python
# SAFE (after fix)
if (isinstance(needs_visualization, bool) and needs_visualization) or (
    isinstance(needs_visualization, str)
    and needs_visualization.strip().lower() == "true"
):
```

This handles both the `bool` return path (exception fallback) and the `str` return path (successful LLM call) without executing any code from the model's response.

---

### MEDIUM — SQL Injection via Untrusted Pub/Sub Input

| Field | Detail |
|---|---|
| **Severity** | Medium |
| **CWE** | CWE-89 (Improper Neutralization of Special Elements used in an SQL Command) |
| **File** | `gemini/sample-apps/genwealth/function-scripts/analyze-prospectus/main.py` |
| **Line** | 54 (before fix) |

**Description:**  
The Cloud Function receives a `ticker` symbol from a Pub/Sub message and inserts it directly into an SQL query using Python string formatting:

```python
# VULNERABLE (before fix)
sql = f"SELECT content FROM {table_name} WHERE ticker = '{ticker}' ORDER BY page, page_chunk"
```

A message publisher with access to the Pub/Sub topic could inject SQL by providing a payload such as `' OR '1'='1`, potentially exfiltrating or corrupting the AlloyDB database.

**Remediation applied:**  
Converted to a SQLAlchemy parameterized query:

```python
# SAFE (after fix)
sql = sqlalchemy.text(
    f"SELECT content FROM {table_name} WHERE ticker = :ticker ORDER BY page, page_chunk"
)
# ...
result = db_conn.execute(sql, {"ticker": ticker}).fetchall()
```

The `ticker` value is now bound as a query parameter, preventing injection regardless of its content. The `table_name` variable is a trusted local constant, not derived from external input.

---

## Areas Reviewed — No Issues Found

| Area | Verdict |
|---|---|
| Malicious / autonomous agents | None found |
| Backdoors or persistence mechanisms | None found |
| Data exfiltration payloads | None found |
| Obfuscated code | None found |
| Hardcoded secrets or credentials | None found (`DEMO_KEY` is a documented public API placeholder) |
| References to unauthorized external applications | None found |
| `subprocess` usage | Safe — all calls use list arguments, no `shell=True` |
| Unsafe deserialization (`pickle`, `yaml.load`) | None found |
| Additional `eval()` / `exec()` on untrusted input | None found beyond the fixed instance |
| Additional SQL injection vectors | None found beyond the fixed instance |

---

## Recommendations

1. **Sanitize all LLM outputs before use in control flow.** Never pass model response text to `eval()`, `exec()`, or similar dynamic execution functions. Use structured output (e.g., Pydantic models, JSON schema enforcement via `response_schema`) to obtain typed boolean/enum values from the model.

2. **Always use parameterized queries.** Any SQL query that incorporates external input — including values from Pub/Sub, HTTP requests, or environment variables — must use bound parameters rather than string interpolation.

3. **Consider adding a SAST step to CI.** Tools such as Bandit (Python) can flag `eval()` and string-formatted SQL queries automatically in pull requests.
