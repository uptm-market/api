
create Table users(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    email_verification_key VARCHAR(255),
    email_verification_time TIMESTAMP,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    cell_phone VARCHAR(20) NOT NULL,
    city VARCHAR(255),
    state VARCHAR(255),
    zip_code VARCHAR(20),
    district VARCHAR(255),
    street VARCHAR(255),
    street_number VARCHAR(20),
    complement VARCHAR(255),
    birthdate VARCHAR(20),
    gender VARCHAR(10),
    cpf VARCHAR(14),
    source TINYINT,
    record_date TIMESTAMP,
    level INTEGER
);


