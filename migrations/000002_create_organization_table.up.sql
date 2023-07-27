CREATE TABLE IF NOT EXISTS organizations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    kpp VARCHAR(255) NOT NULL,
    orgn VARCHAR(255) NOT NULL,
    inn VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    owner_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(owner_id) references users(id)
);

ALTER TABLE users ADD COLUMN organization_id INT,
ADD CONSTRAINT fk_organization FOREIGN KEY (organization_id) REFERENCES organizations(id);

ALTER TABLE users ADD COLUMN roles TEXT[];