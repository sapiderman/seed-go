DROP TABLE IF EXISTS users, device;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    username VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR UNIQUE NOT NULL, 
    email VARCHAR(100) UNIQUE NOT NULL ,
    password VARCHAR(255) NOT NULL,
    pin INT,
    device INT REFERENCES device(id)
    );


CREATE TABLE IF NOT EXISTS device (
    id INT PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    phone_brand VARCHAR(255) NOT NULL,
    phone_model VARCHAR(100) NOT NULL, 
    year VARCHAR(100) NOT NULL ,
    push_notif_id VARCHAR,
    device_id VARCHAR
);
