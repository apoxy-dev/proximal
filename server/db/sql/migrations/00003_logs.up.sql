CREATE TABLE access_logs (
  request_id     VARCHAR(255) PRIMARY KEY,
  start_time     INTEGER NOT NULL,
  end_time       INTEGER NOT NULL,
  entry 	 BLOB NOT NULL
);

CREATE INDEX access_logs_start_time_end_time ON access_logs (start_time, end_time);

CREATE TABLE traces (
  request_id     VARCHAR(255) PRIMARY KEY,
  trace   	 BLOB NOT NULL
);
