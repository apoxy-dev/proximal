CREATE TABLE middlewares (
  slug                VARCHAR(255) PRIMARY KEY,
  source_type         TEXT NOT NULL CHECK (source_type IN ('GITHUB', 'DIRECT')),
  ingest_params_json  BLOB NOT NULL,
  runtime_params_json BLOB NOT NULL,
  live_build_sha      VARCHAR(64),
  status              TEXT NOT NULL CHECK (status IN ('UNKNOWN', 'PENDING', 'READY', 'PENDING_READY', 'ERRORED')),
  status_detail       TEXT,
  created_at          DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  updated_at          DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW'))
);

CREATE TABLE builds (
  sha                 VARCHAR(64) NOT NULL,
  middleware_slug     VARCHAR(255) NOT NULL,
  status              TEXT NOT NULL CHECK (status IN ('UNKNOWN', "PREPARING", 'RUNNING', 'READY', 'ERRORED')),
  status_detail       TEXT NOT NULL,
  output_path         VARCHAR(4096),
  started_at          DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  updated_at          DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  deleted_at          DATETIME,
  PRIMARY KEY (middleware_slug, sha),
  FOREIGN KEY (middleware_slug) REFERENCES middlewares(slug)
);
