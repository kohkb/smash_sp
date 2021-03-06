CREATE TABLE IF NOT EXISTS smash_sp.favorites(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    video_id VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
