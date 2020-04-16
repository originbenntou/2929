DROP DATABASE IF EXISTS trend;
CREATE DATABASE trend DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

use trend;

DROP TABLE IF EXISTS suggest;
CREATE TABLE suggest (
  id INT unsigned NOT NULL auto_increment,
  keyword VARCHAR(255) NOT NULL,
  result json DEFAULT NULL,
  status tinyint(1) NOT NULL DEFAULT 0,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY (keyword)
) COMMENT 'サジェストキーワード';

DROP TABLE IF EXISTS history;
CREATE TABLE history (
  id INT unsigned NOT NULL auto_increment,
  user_id INT unsigned NOT NULL,
  suggest_id INT unsigned NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY (user_id, suggest_id),
  FOREIGN KEY (user_id) REFERENCES account.user(id),
  FOREIGN KEY (suggest_id) REFERENCES suggest(id)
) COMMENT 'トレンド検索履歴';
