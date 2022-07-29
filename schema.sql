-- table for currencies
create table currencies
(
    id              bigserial primary key,
    currency        text,
    value           text,
    created_at      timestamp with time zone,
    last_updated_at timestamp with time zone
);

alter table currencies
    owner to postgres;

create index idx_currency_last_updated_at
    on currencies (currency, last_updated_at);

--table for history in requests
create table history_apis
(
    id               bigserial
        primary key,
    created_at       timestamp with time zone,
    time_response_ms bigint,
    error            text
);

alter table history_apis
    owner to postgres;