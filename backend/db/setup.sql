CREATE DATABASE IF NOT EXISTS smash_sp;
CREATE USER IF NOT EXISTS smash_sp IDENTIFIED BY 'password';
GRANT all ON smash_sp.* TO 'smash_sp'@'localhost';
FLUSH PRIVILEGES;
