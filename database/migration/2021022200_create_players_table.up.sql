CREATE TABLE players
(
    id            varchar(36) PRIMARY KEY,
    user_id       varchar(36) NOT NULL,
    name          varchar(255) NOT NULL,
    level         varchar(36) NOT NULL,
    job           varchar(255) NOT NULL,
    created_at    timestamp   DEFAULT CURRENT_TIMESTAMP,
    updated_at    timestamp   DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = innodb
  DEFAULT CHARSET = utf8mb4;