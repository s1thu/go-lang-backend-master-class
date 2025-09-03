-- Accounts table
CREATE TABLE accounts (
    id bigserial PRIMARY KEY,
    owner varchar(255) NOT NULL,
    balance bigint NOT NULL,
    currency varchar(255) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);

-- Entries table
CREATE TABLE entries (
    id bigserial PRIMARY KEY,
    account_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT fk_entries_account FOREIGN KEY (account_id) REFERENCES accounts (id)
);

-- Transfers table
CREATE TABLE transfers (
    id bigserial PRIMARY KEY,
    from_account_id bigint NOT NULL,
    to_account_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT fk_transfers_from FOREIGN KEY (from_account_id) REFERENCES accounts (id),
    CONSTRAINT fk_transfers_to FOREIGN KEY (to_account_id) REFERENCES accounts (id)
);

-- Indexes
CREATE INDEX accounts_index_0 ON accounts (owner);
CREATE INDEX entries_index_1 ON entries (account_id);
CREATE INDEX transfers_index_2 ON transfers (from_account_id);
CREATE INDEX transfers_index_3 ON transfers (to_account_id);
CREATE INDEX transfers_index_4 ON transfers (from_account_id, to_account_id);

-- Comments
COMMENT ON COLUMN entries.amount IS 'can be negative or positive';
COMMENT ON COLUMN transfers.amount IS 'must be positive';