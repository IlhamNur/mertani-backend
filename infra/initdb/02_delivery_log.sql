CREATE TABLE IF NOT EXISTS delivery_log (
    id SERIAL PRIMARY KEY,
    device_id INTEGER,
    payload JSONB,
    client_url TEXT,
    status VARCHAR(30) DEFAULT 'pending',
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 5,
    next_attempt_at TIMESTAMP DEFAULT now(),
    last_attempt_at TIMESTAMP,
    last_error TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_delivery_next_attempt ON delivery_log (next_attempt_at);

CREATE INDEX IF NOT EXISTS idx_delivery_status ON delivery_log (status);