CREATE KEYSPACE IF NOT EXISTS cdb
WITH replication = {
    'class': 'NetworkTopologyStrategy',
    'datacenter1': 1
};

CREATE TABLE IF NOT EXISTS cdb.clients (
    id          uuid,
    key         uuid,
    secret      text,
    created_at  timestamp,
    deleted_at  timestamp,
    PRIMARY KEY (key)
) WITH comment = 'The id field is the external identifier.';

CREATE TABLE IF NOT EXISTS cdb.tokens (
    hash          text,
    client_key    text,
    client_secret text,
    scopes        set<text>,
    PRIMARY KEY (hash, client_key)
) WITH comment = 'Holds both access and refresh tokens identified by type field.';