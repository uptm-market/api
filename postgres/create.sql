CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    email_verification_key VARCHAR(255),
    email_verification_time TIMESTAMP,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    cell_phone VARCHAR(20) ,
    city VARCHAR(255),
    state VARCHAR(255),
    zip_code VARCHAR(20),
    district VARCHAR(255),
    street VARCHAR(255),
    street_number VARCHAR(20),
    complement VARCHAR(255),
    birthdate DATE,
    gender varchar,
    cpf VARCHAR(14),
    type_plan integer, 
    source SMALLINT,
    record_date TIMESTAMP ,
    level INTEGER
);



CREATE TABLE user_sub_accounts (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    key_api VARCHAR(255) UNIQUE
);

CREATE TABLE related_sub_accounts_in_master_user (
    sub_account_id INT,
    user_id INT,
    PRIMARY KEY (sub_account_id, user_id),
    FOREIGN KEY (sub_account_id) REFERENCES user_sub_accounts(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE facebook_campaign_ad_account (
    id INT PRIMARY KEY ,
    app_secret VARCHAR(255) ,
    business_id ,
    user_id INT NOT NULL, 
    FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TABLE facebook_campaign_ad_account_token(
    id INT PRIMARY KEY ,
    token_id VARCHAR(255) NOT NULL
);