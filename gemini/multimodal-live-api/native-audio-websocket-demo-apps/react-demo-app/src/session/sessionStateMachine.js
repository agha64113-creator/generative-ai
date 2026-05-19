import crypto from 'node:crypto';

const SNAPSHOT_VERSION = 1;

export const SessionStatus = Object.freeze({
  IDLE: 'idle',
  RUNNING: 'running',
  PAUSED: 'paused',
  SUBMITTED: 'submitted',
  BLOCKED: 'blocked',
  RECOVERY: 'recovery',
});

export const EventType = Object.freeze({
  START: 'START',
  PAUSE: 'PAUSE',
  RESUME: 'RESUME',
  SUBMIT: 'SUBMIT',
  BLOCK_ADVANCE: 'BLOCK_ADVANCE',
  RESTORE: 'RESTORE',
});

export function createInitialSessionState() {
  return {
    status: SessionStatus.IDLE,
    timerMs: 0,
    questionIndex: 0,
    blockMetadata: null,
    answerMap: {},
    recoveryMode: false,
  };
}

export function reduceSessionState(currentState, event) {
  const state = structuredClone(currentState);

  switch (event.type) {
    case EventType.START:
      return { ...state, status: SessionStatus.RUNNING, recoveryMode: false };
    case EventType.PAUSE:
      return { ...state, status: SessionStatus.PAUSED, timerMs: event.timerMs ?? state.timerMs };
    case EventType.RESUME:
      return { ...state, status: SessionStatus.RUNNING, timerMs: event.timerMs ?? state.timerMs };
    case EventType.SUBMIT:
      return { ...state, status: SessionStatus.SUBMITTED, answerMap: { ...state.answerMap, ...(event.answers ?? {}) } };
    case EventType.BLOCK_ADVANCE:
      return {
        ...state,
        status: SessionStatus.BLOCKED,
        blockMetadata: {
          reason: event.reason,
          blockedAtQuestionIndex: state.questionIndex,
          blockedAtTimerMs: state.timerMs,
        },
      };
    case EventType.RESTORE:
      return {
        ...event.snapshot,
        status: event.partial ? SessionStatus.RECOVERY : event.snapshot.status,
        recoveryMode: Boolean(event.partial),
      };
    default:
      return state;
  }
}

function stableStringify(value) {
  if (value === null || typeof value !== 'object') return JSON.stringify(value);
  if (Array.isArray(value)) return `[${value.map((x) => stableStringify(x)).join(',')}]`;
  const keys = Object.keys(value).sort();
  return `{${keys.map((k) => `${JSON.stringify(k)}:${stableStringify(value[k])}`).join(',')}}`;
}

function checksumPayload(payload) {
  return crypto.createHash('sha256').update(stableStringify(payload)).digest('hex');
}

export function buildSnapshot(state) {
  const payload = {
    version: SNAPSHOT_VERSION,
    timerMs: state.timerMs,
    questionIndex: state.questionIndex,
    blockMetadata: state.blockMetadata,
    answerMap: state.answerMap,
    status: state.status,
  };

  return {
    ...payload,
    checksum: checksumPayload(payload),
  };
}

export function persistSnapshot(storage, key, state) {
  const snapshot = buildSnapshot(state);
  storage.setItem(key, JSON.stringify(snapshot));
  return snapshot;
}

export function restoreSnapshot(storage, key) {
  const raw = storage.getItem(key);
  if (!raw) return { state: null, mode: 'missing' };

  let parsed;
  try {
    parsed = JSON.parse(raw);
  } catch {
    return { state: null, mode: 'corrupt' };
  }

  const { checksum, ...payload } = parsed;
  if (payload.version !== SNAPSHOT_VERSION) {
    return { state: null, mode: 'version_mismatch' };
  }

  if (checksumPayload(payload) !== checksum) {
    return { state: null, mode: 'corrupt' };
  }

  const restored = {
    status: payload.status,
    timerMs: payload.timerMs,
    questionIndex: payload.questionIndex,
    blockMetadata: payload.blockMetadata,
    answerMap: payload.answerMap,
    recoveryMode: false,
  };

  const hasPartial = restored.status === SessionStatus.PAUSED || restored.status === SessionStatus.BLOCKED;
  return { state: restored, mode: hasPartial ? 'partial' : 'complete' };
}

export function transitionWithPersistence({ state, event, storage, key }) {
  const next = reduceSessionState(state, event);
  persistSnapshot(storage, key, next);
  return next;
}
