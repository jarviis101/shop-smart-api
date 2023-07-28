CREATE TABLE IF NOT EXISTS otp (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) NOT NULL,
    is_used BOOLEAN NOT NULL DEFAULT false,
    owner_id INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expired_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(owner_id) REFERENCES users(id) ON DELETE CASCADE
);