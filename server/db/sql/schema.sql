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

CREATE TABLE endpoints (
  cluster        VARCHAR(255) PRIMARY KEY,
  is_domain      BOOLEAN NOT NULL,
  use_tls        BOOLEAN DEFAULT FALSE,
  lookup_family  TEXT NOT NULL CHECK (lookup_family IN ('V4_ONLY', 'V4_FIRST', 'V6_ONLY', 'V6_FIRST')),
  created_at     DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  updated_at     DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW'))
);

CREATE TABLE endpoint_default (
  cluster      VARCHAR(255) PRIMARY KEY,
  updated_at   DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  FOREIGN KEY (cluster) REFERENCES endpoints(cluster)
);

CREATE TABLE endpoint_addresses (
  cluster      VARCHAR(255) NOT NULL,
  host         VARCHAR(255) NOT NULL,
  port         INTEGER NOT NULL,
  created_at   DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  updated_at   DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  PRIMARY KEY (cluster, host, port),
  FOREIGN KEY (cluster) REFERENCES endpoints(cluster)
);

CREATE TABLE access_logs (
  request_id     VARCHAR(255) PRIMARY KEY,
  start_time     INTEGER NOT NULL,
  end_time       INTEGER NOT NULL,
  entry 	 BLOB NOT NULL
);

CREATE TABLE traces (
  request_id     VARCHAR(255) PRIMARY KEY,
  trace   	 BLOB NOT NULL
);
