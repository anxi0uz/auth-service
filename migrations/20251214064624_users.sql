-- +goose Up
-- +goose StatementBegin
create table auth_users(
    id UUID PRIMARY KEY NOT NULL,
    login text not null,
    email text not null,
    password text not null,
    confirmed_at timestamptz,
    is_active boolean not null default true,
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    failed_login_attempts integer not null default 0,
    last_login_at timestamptz,
    password_changed_at timestamptz
);

create index idx_auth_users_email on auth_users(email);
create index idx_auth_users_login on auth_users(login);
create index idx_auth_users_is_active on auth_users(is_active);
create index idx_auth_users_created_at on auth_users(created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_auth_users_email;
drop index if exists idx_auth_users_created_at;
drop index if exists idx_auth_users_is_active;
drop index if exists idx_auth_users_login;

drop table auth_users
-- +goose StatementEnd
