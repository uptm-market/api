
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




-- Tabela Hours
CREATE TABLE hours (
    ID INT PRIMARY KEY,
    Hour TEXT,
    Professional TEXT,
    Local TEXT,
    UserId TEXT,
    UserName TEXT,
    ProfessionalId TEXT,
    CreatedTime TEXT
);

-- Tabela HoursService
CREATE TABLE hoursservice (
    ID TEXT PRIMARY KEY,
    Hour TIMESTAMP, -- Aqui você precisa ajustar de acordo com sua necessidade
    FOREIGN KEY (ID) REFERENCES Hours(ID)
);

-- Tabela Service
CREATE TABLE service (
    ID INT PRIMARY KEY,
    Name VARCHAR,
    Description VARCHAR,
    Professional VARCHAR,
    CreatedTime TIMESTAMP,
    FOREIGN KEY (ID) REFERENCES HoursService(ID)
);

-- Tabela UserMaster
CREATE TABLE usermaster (
    ID INT PRIMARY KEY,
    User TEXT,
    Passwor TEXT, -- Corrigi o nome da coluna
    Master BOOLEAN
);
