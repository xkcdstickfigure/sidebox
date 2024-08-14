create table account (
    id             uuid          primary key,
    name           text          not null,
    email          text          not null unique,
    google_id      text          not null unique,
    created_at     timestamptz   not null default now(),
    last_used_at   timestamptz   not null default now(),
    ref            text          not null
);

create table session (
    id           uuid          primary key,
    account_id   uuid          not null references account on delete cascade,
    token        text          not null unique,
    address      text          not null,
    user_agent   text          not null,
    created_at   timestamptz   not null default now()
);

create table inbox (
    id           uuid          primary key,
    account_id   uuid          not null references account on delete cascade,
    code         text          not null unique,
    name         text          not null,
    muted        boolean       not null default false,
    unread       boolean       not null default false,
    created_at   timestamptz   not null default now()
);

create table message (
    id             uuid          primary key,
    inbox_id       uuid          not null references inbox on delete cascade,
    message_id     text          not null,
    from_name      text          not null,
    from_address   text          not null,
    subject        text          not null,
    body           text          not null,
    html           boolean       not null,
    date           timestamptz   not null default now()
);