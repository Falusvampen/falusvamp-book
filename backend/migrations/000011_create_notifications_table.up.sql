--  create notifications TABLE, migrate up

CREATE TABLE notifications (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    to_id INTEGER NOT NULL,
    notitype VARCHAR NOT NULL,
    content VARCHAR NOT NULL,
    is_read VARCHAR NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY (to_id) REFERENCES users (id)
    );
