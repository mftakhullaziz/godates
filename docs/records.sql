CREATE TABLE users
(
    user_id       SERIAL PRIMARY KEY,
    username      VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255)        NOT NULL,
    email         VARCHAR(255) UNIQUE NOT NULL,
    verified      BOOLEAN   DEFAULT FALSE,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_profiles
(
    profile_id SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (user_id),
    age        INTEGER,
    bio        TEXT,
    photos     JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE swipes
(
    swipe_id   SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (user_id),
    profile_id INTEGER,
    action     VARCHAR(5) CHECK (action IN ('left', 'right')
) ,
    swipe_date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE daily_quotas
(
    quota_id    SERIAL PRIMARY KEY,
    user_id     INTEGER REFERENCES users (user_id),
    date        DATE,
    swipe_count INTEGER DEFAULT 0
);

CREATE TABLE premium_packages
(
    package_id       SERIAL PRIMARY KEY,
    description      TEXT,
    price            NUMERIC,
    unlimited_swipes BOOLEAN   DEFAULT FALSE,
    verified         BOOLEAN   DEFAULT FALSE,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_premiums
(
    purchase_id             SERIAL PRIMARY KEY,
    user_id                 INTEGER REFERENCES users (user_id),
    package_id              INTEGER REFERENCES premium_packages (package_id),
    purchase_date           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expiry_date             TIMESTAMP,
    unlimited_swipes_active BOOLEAN   DEFAULT FALSE,
    verified_active         BOOLEAN   DEFAULT FALSE
);

CREATE TABLE viewed_profiles
(
    view_id    SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (user_id),
    profile_id INTEGER,
    date       DATE DEFAULT CURRENT_DATE
);
