CREATE TABLE proxies (
  key 	       VARCHAR(255) PRIMARY KEY,
  created_at   DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  updated_at   DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  deleted_at   DATETIME
);

CREATE TABLE proxy_middlewares (
  proxy_key       VARCHAR(255) NOT NULL,
  middleware_slug VARCHAR(255) NOT NULL,
  created_at      DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  updated_at      DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  deleted_at      DATETIME,
  PRIMARY KEY (proxy_key, middleware_slug),
  FOREIGN KEY (proxy_key) REFERENCES proxies(key),
  FOREIGN KEY (middleware_slug) REFERENCES middlewares(slug)
);

CREATE TABLE proxy_upstreams (
  proxy_key       VARCHAR(255) NOT NULL,
  upstream        BLOB NOT NULL,
  created_at      DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  updated_at      DATETIME DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),
  deleted_at      DATETIME,
  PRIMARY KEY (proxy_key, upstream),
  FOREIGN KEY (proxy_key) REFERENCES proxies(key)
);
