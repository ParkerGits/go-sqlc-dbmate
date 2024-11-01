-- migrate:up
CREATE TABLE book (
	id INTEGER PRIMARY KEY NOT NULL,
	title VARCHAR(255) NOT NULL,
	author VARCHAR(255) NOT NULL
);

-- migrate:down
DROP TABLE book;
