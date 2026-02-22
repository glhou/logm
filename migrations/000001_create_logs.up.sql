CREATE TABLE logs (
  id SERIAL PRIMARY KEY,
  service TEXT NOT NULL,
  level TEXT NOT NULL,
  message TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_logs_created_at ON logs(created_at DESC);
CREATE INDEX idx_logs_service ON logs(service);
