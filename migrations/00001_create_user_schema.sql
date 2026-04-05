-- +goose Up
create table users (
    id uuid primary key default uuidv7(),
    name text not null,
    email text not null,
    password_hash text not null,
    created_at timestamptz default now()
    last_login timestamptz
);

create table refresh_tokens (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null references users(id) on delete cascade,
    token varchar(255) unique not null,
    expires_at timestamp not null,
    created_at timestamp not null default current_timestamp,
    revoked boolean not null default false
);

-- +goose Down
drop table users;
drop table refresh_tokens;
