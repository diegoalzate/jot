-- +goose Up
-- +goose StatementBegin
CREATE TABLE discord_servers (
    id UUID PRIMARY KEY,
    discord_id TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE github_installations (
    id UUID PRIMARY KEY,
    installation_id TEXT UNIQUE NOT NULL,    
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE server_installations (
    id UUID PRIMARY KEY,
    discord_server_id UUID NOT NULL REFERENCES discord_servers(id),
    github_installation_id UUID NOT NULL REFERENCES github_installations(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE discord_servers;
DROP TABLE github_installations;
DROP TABLE server_installations;
-- +goose StatementEnd

