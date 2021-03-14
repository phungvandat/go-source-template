-- +migrate Up
CREATE TABLE IF NOT EXISTS companies (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    domain VARCHAR(255) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT uni_domain UNIQUE(domain)
);

CREATE TRIGGER trig_set_updated_at
BEFORE UPDATE ON companies
FOR EACH ROW
EXECUTE PROCEDURE func_set_updated_at();

-- +migrate Down
DROP TRIGGER IF EXISTS trig_set_updated_at ON companies;

DROP TABLE IF EXISTS companies;