DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uint256') THEN
        CREATE DOMAIN UINT256 AS NUMERIC
            CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
ELSE
    ALTER DOMAIN UINT256 DROP CONSTRAINT uint256_check;
        ALTER DOMAIN UINT256 ADD
            CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
END IF;
END $$;

CREATE TABLE IF NOT EXISTS state_root (
    guid                  VARCHAR PRIMARY KEY,
    l2_block_num          UINT256 NOT NULL,
    current_l1            UINT256 NOT NULL,
    current_l1_hash       VARCHAR NOT NULL,
    finalized_l1          UINT256 NOT NULL,
    safe_l1               UINT256 NOT NULL,
    finalized_l2          UINT256 NOT NULL,
    safe_l2               UINT256 NOT NULL,
    state_root            VARCHAR NOT NULL,
    signature             BYTEA,
    is_finalized          SMALLINT NOT NULL DEFAULT 0,
    timestamp             INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS state_root_timestamp ON state_root(timestamp);
CREATE INDEX IF NOT EXISTS state_root_current_l1 ON state_root(current_l1);
CREATE INDEX IF NOT EXISTS state_root_l2_block_num ON state_root(l2_block_num);
CREATE INDEX IF NOT EXISTS state_root_state_root ON state_root(state_root);
CREATE INDEX IF NOT EXISTS state_root_signature ON state_root(signature);

CREATE TABLE IF NOT EXISTS node (
    guid                  VARCHAR PRIMARY KEY,
    state_root            VARCHAR NOT NULL,
    signature             BYTEA NOT NULL,
    vote                  SMALLINT NOT NULL DEFAULT 0,
    timestamp             INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS node_timestamp ON node(timestamp);
CREATE INDEX IF NOT EXISTS node_state_root ON node(state_root);
CREATE INDEX IF NOT EXISTS node_signature ON node(signature);

CREATE TABLE IF NOT EXISTS vote (
    guid                  VARCHAR PRIMARY KEY,
    l2_block_num          UINT256 NOT NULL,
    node                  VARCHAR NOT NULL,
    signature             BYTEA,
    result                SMALLINT NOT NULL DEFAULT 0,
    timestamp             INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS vote_timestamp ON vote(timestamp);
CREATE INDEX IF NOT EXISTS vote_l2_block_num ON vote(l2_block_num);
CREATE INDEX IF NOT EXISTS vote_signature ON vote(signature);
CREATE INDEX IF NOT EXISTS vote_result ON vote(result);