CREATE DATABASE IF NOT EXISTS vgs CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
CREATE USER IF NOT EXISTS 'user'@'%' IDENTIFIED BY 'root';
GRANT ALL PRIVILEGES ON vgs.* TO 'user'@'%';
SET session wait_timeout=259200;
SET global max_allowed_packet=10485760;

FLUSH PRIVILEGES;

CREATE TABLE IF NOT EXISTS vgs.hoge1 (id int auto_increment NOT NULL PRIMARY KEY, name varchar(10), index(id));
