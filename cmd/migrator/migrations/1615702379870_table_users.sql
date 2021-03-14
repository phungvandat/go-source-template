-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    company_id UUID NOT NULL,
    fullname TEXT NOT NULL,
    username VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    role_type INT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_company_id FOREIGN KEY (company_id) REFERENCES companies(id),
    CONSTRAINT uni_username_company_id UNIQUE(username, company_id)
);

CREATE TRIGGER trig_set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE func_set_updated_at();

-- +migrate Down
DROP TRIGGER IF EXISTS trig_set_updated_at ON users;

DROP TABLE IF EXISTS users;