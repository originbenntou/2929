DROP DATABASE IF EXISTS trend;
CREATE DATABASE trend DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

use trend;

DROP TABLE IF EXISTS suggest;
CREATE TABLE suggest (
  id INT unsigned NOT NULL auto_increment COMMENT 'サジェストID',
  keyword VARCHAR(255) NOT NULL COMMENT 'トレンド検索キーワード',
  date DATE NOT NULL COMMENT '検索開始日',
  result json DEFAULT NULL COMMENT 'トレンド検索結果',
  status tinyint(1) NOT NULL DEFAULT 0 COMMENT 'トレンド検索進捗',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (id),
  UNIQUE KEY (keyword, date)
) COMMENT 'サジェスト情報';

DROP TABLE IF EXISTS history;
CREATE TABLE history (
  id INT unsigned NOT NULL auto_increment COMMENT 'トレンド検索履歴ID',
  user_id INT unsigned NOT NULL COMMENT 'ユーザーID',
  suggest_id INT unsigned NOT NULL COMMENT 'サジェストID',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  PRIMARY KEY (id),
  UNIQUE KEY (user_id, suggest_id),
  FOREIGN KEY (user_id) REFERENCES account.user(id),
  FOREIGN KEY (suggest_id) REFERENCES suggest(id)
) COMMENT 'トレンド検索履歴';
