CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    trx_number VARCHAR(255),
    value DOUBLE PRECISION NOT NULL,
    status BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    actioned_at TIMESTAMPTZ,
    owner_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(owner_id) REFERENCES users(id)
);
