DROP DATABASE IF EXISTS sample_docker_compose;
CREATE DATABASE IF NOT EXISTS sample_docker_compose DEFAULT CHARACTER SET utf8;
USE sample_docker_compose;
CREATE TABLE IF NOT EXISTS user ( usr_id INT NOT NULL AUTO_INCREMENT, usr_name VARCHAR(200) NOT NULL, PRIMARY KEY (usr_id) );
INSERT INTO user (usr_name) VALUES( 'test' );
