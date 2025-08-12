PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS secrets (
  id INTEGER PRIMARY KEY,
  key_name TEXT NOT NULL UNIQUE,
  value_ct BLOB NOT NULL,
  created_at INTEGER NOT NULL DEFAULT (strftime ('%s', 'now')),
  updated_at INTEGER NOT NULL DEFAULT (strftime ('%s', 'now'))
);

CREATE INDEX IF NOT EXISTS idx_secrets_key ON secrets (key_name);

PRAGMA user_version = 1;
