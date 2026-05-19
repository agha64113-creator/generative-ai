import test from 'node:test';
import assert from 'node:assert/strict';
import {
  createInitialSessionState,
  EventType,
  restoreSnapshot,
  SessionStatus,
  transitionWithPersistence,
  reduceSessionState,
} from '../src/session/sessionStateMachine.js';

class MemoryStorage {
  constructor() { this.map = new Map(); }
  setItem(key, value) { this.map.set(key, value); }
  getItem(key) { return this.map.get(key) ?? null; }
}

test('deterministic replay across pause/resume timestamps and block boundaries', () => {
  const storage = new MemoryStorage();
  const key = 'session';
  let state = createInitialSessionState();

  const sequence = [
    { type: EventType.START },
    { type: EventType.PAUSE, timerMs: 12000 },
    { type: EventType.RESUME, timerMs: 18000 },
    { type: EventType.PAUSE, timerMs: 30321 },
    { type: EventType.RESUME, timerMs: 40000 },
    { type: EventType.BLOCK_ADVANCE, reason: 'missing-required-answer' },
  ];

  for (const event of sequence) {
    state = transitionWithPersistence({ state, event, storage, key });
  }

  assert.equal(state.status, SessionStatus.BLOCKED);
  assert.equal(state.timerMs, 40000);
  assert.equal(state.blockMetadata.blockedAtQuestionIndex, 0);
  assert.equal(state.blockMetadata.blockedAtTimerMs, 40000);

  const restored = restoreSnapshot(storage, key);
  assert.equal(restored.mode, 'partial');

  const recovery = reduceSessionState(state, {
    type: EventType.RESTORE,
    snapshot: restored.state,
    partial: restored.mode === 'partial',
  });
  assert.equal(recovery.status, SessionStatus.RECOVERY);
  assert.equal(recovery.recoveryMode, true);
});

test('rejects corrupt checksum and version mismatch', () => {
  const storage = new MemoryStorage();
  const key = 'session';

  storage.setItem(key, JSON.stringify({ version: 1, timerMs: 1, checksum: 'bad' }));
  assert.equal(restoreSnapshot(storage, key).mode, 'corrupt');

  storage.setItem(key, JSON.stringify({ version: 999, timerMs: 1, checksum: 'x' }));
  assert.equal(restoreSnapshot(storage, key).mode, 'version_mismatch');
});
