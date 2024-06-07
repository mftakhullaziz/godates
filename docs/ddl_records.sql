CREATE TABLE accounts
(
    account_id    INTEGER AUTO_INCREMENT PRIMARY KEY,
    username      VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255)        NOT NULL,
    email         VARCHAR(255) UNIQUE NOT NULL,
    phone_number  VARCHAR(255)        NOT NULL,
    verified      BOOLEAN   DEFAULT FALSE,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users
(
    user_id       INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id    INTEGER NOT NULL,
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
    action     VARCHAR(5) CHECK (action IN ('left', 'right')
) ,
    swipe_date DATE DEFAULT (CURRENT_DATE),
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (account_id) REFERENCES accounts(account_id)
);

CREATE TABLE daily_quotas
(
    quota_id    INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id  INTEGER NOT NULL,
    date        DATE,
    swipe_count INTEGER DEFAULT 0,
    FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);

CREATE TABLE premium_packages
(
    package_id       INTEGER AUTO_INCREMENT PRIMARY KEY,
    description      TEXT,
    price            NUMERIC,
    unlimited_swipes BOOLEAN   DEFAULT FALSE,
    verified         BOOLEAN   DEFAULT FALSE,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE account_premiums
(
    purchase_id             INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id              INTEGER NOT NULL,
    package_id              INTEGER NOT NULL,
    purchase_date           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expiry_date             TIMESTAMP,
    unlimited_swipes_active BOOLEAN   DEFAULT FALSE,
    verified_active         BOOLEAN   DEFAULT FALSE,
    FOREIGN KEY (account_id) REFERENCES accounts (account_id),
    FOREIGN KEY (package_id) REFERENCES premium_packages (package_id)
);

CREATE TABLE viewed_user_accounts
(
    view_id    INTEGER AUTO_INCREMENT PRIMARY KEY,
    account_id INTEGER NOT NULL,
    user_id    INTEGER,
    date       DATE DEFAULT (CURRENT_DATE),
    FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);

CREATE TABLE login_histories
(
    login_histories_id INTEGER AUTO_INCREMENT PRIMARY KEY,
    user_id            INTEGER NOT NULL,
    account_id         INTEGER NOT NULL,
    login_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    logout_at          TIMESTAMP DEFAULT NULL,
    active_duration    TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);