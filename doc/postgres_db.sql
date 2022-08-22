CREATE TABLE IF NOT EXISTS users
(
    id         BIGSERIAL PRIMARY KEY,
    token      VARCHAR     NOT NULL UNIQUE,
    login      VARCHAR(40) NOT NULL,
    service    VARCHAR(20) NOT NULL,
    uchproc_id BIGINT      NOT NULL,
    reg_date   TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    expire_at  TIMESTAMP   NOT NULL,
    status     varchar(30) NOT NULL,
    UNIQUE (login, service)
);