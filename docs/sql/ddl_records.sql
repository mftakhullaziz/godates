CREATE TABLE accounts
(
    account_id    INTEGER AUTO_INCREMENT PRIMARY KEY,
    username      VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255)        NOT NULL,
    email         VARCHAR(255) UNIQUE NOT NULL,
    verified      BOOLEAN   DEFAULT FALSE,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users
(
    user_id       INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id    INTEGER NOT NULL,
    full_name     VARCHAR(255),
    date_of_birth DATE,
    age           INTEGER,
    gender        VARCHAR(5),
    address       VARCHAR(255),
    bio           TEXT,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);

CREATE TABLE storages
(
    storage_id  INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id  INTEGER NOT NULL,
    user_id     INTEGER NOT NULL,
    videos_path VARCHAR(255),
    photos_path VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);

CREATE TABLE swipes
(
    swipe_id   INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id INTEGER NOT NULL,
    user_id    INTEGER NOT NULL,
    action     VARCHAR(10) CHECK (action IN ('PASSED', 'LIKED')
) ,
    account_id_swipe INTEGER NOT NULL,
    swipe_date DATE DEFAULT (CURRENT_DATE),
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (account_id) REFERENCES accounts(account_id),
        FOREIGN KEY (account_id_swipe) REFERENCES accounts(account_id)
);

CREATE TABLE daily_quotas
(
    quota_id    INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id  INTEGER NOT NULL,
    date        DATE    DEFAULT (CURRENT_DATE),
    total_quota INTEGER DEFAULT 0,
    swipe_count INTEGER DEFAULT 0,
    FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);

CREATE TABLE packages
(
    package_id                  INTEGER AUTO_INCREMENT PRIMARY KEY,
    package_name                VARCHAR(255),
    description                 TEXT,
    package_duration_in_monthly INTEGER,
    price                       NUMERIC,
    unlimited_swipes            BOOLEAN   DEFAULT FALSE,
    status                      BOOLEAN   DEFAULT FALSE,
    created_at                  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at                  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE account_premiums
(
    purchase_id             INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id              INTEGER NOT NULL,
    package_id              INTEGER NOT NULL,
    purchase_date           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expiry_date             TIMESTAMP,
    unlimited_swipes_active BOOLEAN   DEFAULT FALSE,
    status                  BOOLEAN   DEFAULT FALSE,
    FOREIGN KEY (account_id) REFERENCES accounts (account_id),
    FOREIGN KEY (package_id) REFERENCES packages (package_id)
);

CREATE TABLE view_accounts
(
    view_id    INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id INTEGER NOT NULL,
    user_id    INTEGER NOT NULL,
    date       DATE DEFAULT (CURRENT_DATE),
    FOREIGN KEY (account_id) REFERENCES accounts (account_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE login_histories
(
    login_histories_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    user_id            INTEGER NOT NULL,
    account_id         INTEGER NOT NULL,
    login_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    logout_at          TIMESTAMP DEFAULT NULL,
    duration_in_seconds DOUBLE DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);

CREATE TABLE selection_histories
(
    selection_history_id  INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id_identifier INTEGER,
    account_id            INTEGER,
    selection_date        DATE DEFAULT (CURRENT_DATE),
    FOREIGN KEY (account_id) REFERENCES accounts (account_id),
    FOREIGN KEY (account_id_identifier) REFERENCES accounts (account_id)
);

CREATE TABLE task_histories
(
    task_id               INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id_identifier INTEGER,
    task_name             VARCHAR(255) NOT NULL,
    last_run_timestamp    BIGINT       NOT NULL,
    FOREIGN KEY (account_id_identifier) REFERENCES accounts (account_id)
);
