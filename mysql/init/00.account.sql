DROP DATABASE IF EXISTS account;
CREATE DATABASE account DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

use account;

DROP TABLE IF EXISTS plan;
CREATE TABLE plan (
  id INT unsigned NOT NULL auto_increment,
  name VARCHAR(255) NOT NULL,
  price INT unsigned NOT NULL DEFAULT 0,
  capacity INT unsigned  NOT NULL DEFAULT 3,
  PRIMARY KEY (id)
) COMMENT 'プラン情報';

DROP TABLE IF EXISTS company;
CREATE TABLE company (
  id INT unsigned NOT NULL auto_increment,
  name VARCHAR(255) NOT NULL,
  plan_id INT unsigned NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (plan_id) REFERENCES plan(id)
) COMMENT '会社情報';

DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id INT unsigned NOT NULL auto_increment,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(1023) NOT NULL,
  name VARCHAR(255) NOT NULL,
  company_id INT unsigned NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY (email),
  FOREIGN KEY (company_id) REFERENCES company(id)
) COMMENT 'ユーザー情報';

DROP TABLE IF EXISTS session;
CREATE TABLE session (
  id INT unsigned NOT NULL auto_increment,
  token VARCHAR(255) NOT NULL,
  user_id INT unsigned NOT NULL,
  company_id INT unsigned NOT NULL,
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  updated_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES user(id),
  FOREIGN KEY (company_id) REFERENCES company(id)
) COMMENT 'セッション情報';
