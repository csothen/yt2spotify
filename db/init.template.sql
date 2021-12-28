CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    session_id VARCHAR(255) NOT NULL UNIQUE,
    access_token VARCHAR(255),
    refresh_token VARCHAR(255),
    token_type VARCHAR(50),
    expires_in DATETIME
);
GRANT ALL PRIVILEGES ON (db_user).* TO '(db_name)' @'%'