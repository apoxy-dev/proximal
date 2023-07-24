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
