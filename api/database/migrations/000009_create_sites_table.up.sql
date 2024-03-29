CREATE TABLE IF NOT EXISTS sites (
  id SERIAL PRIMARY KEY,
  url VARCHAR(255) NOT NULL,
  status VARCHAR(15) NOT NULL DEFAULT 'PENDING',
  alert_id INTEGER NULL REFERENCES alerts(id) ON DELETE SET NULL,
  source_id INTEGER NULL REFERENCES sources(id) ON DELETE SET NULL,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL
);
