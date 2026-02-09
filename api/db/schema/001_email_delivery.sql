-- +goose up
CREATE TYPE delivery_status AS ENUM ('pending', 'sent', 'delivered', 'failed', 'bounced', 'complained');

CREATE TABLE email_delivery (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "to" TEXT NOT NULL,
    "from" TEXT NOT NULL,
    cc TEXT,
    bcc TEXT,
    subject TEXT NOT NULL,
    body TEXT NOT NULL,
    delivery_status delivery_status NOT NULL DEFAULT 'pending',
    is_opened BOOLEAN NOT NULL DEFAULT false,
    opened_at TIMESTAMP,
    sent_at TIMESTAMP,
    failed_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    error_message TEXT,
    retry_count INTEGER DEFAULT 0,
    max_retries INTEGER DEFAULT 3,
    metadata JSONB,
    external_id TEXT UNIQUE
);

-- +goose down
DROP TABLE IF EXISTS email_delivery;
DROP TYPE IF EXISTS delivery_status; 