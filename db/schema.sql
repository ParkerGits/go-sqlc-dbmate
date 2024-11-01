CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE book (
	id INTEGER PRIMARY KEY NOT NULL,
	title VARCHAR(255) NOT NULL,
	author VARCHAR(255) NOT NULL
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20241030034429');
