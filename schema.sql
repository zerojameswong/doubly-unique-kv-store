PRAGMA foreign_keys = ON;

CREATE TABLE stores(
store_id INTEGER PRIMARY KEY AUTOINCREMENT,
store_name TEXT NOT NULL
);

CREATE TABLE entries(
entry_id INTEGER PRIMARY KEY AUTOINCREMENT,
key TEXT NOT NULL,
value TEXT NOT NULL,
category TEXT NOT NULL,
store_id INTEGER,
FOREIGN KEY(store_id) REFERENCES stores(store_id)
);

INSERT INTO stores VALUES (NULL, "My Store");
INSERT INTO entries VALUES (NULL, "key 1", "value 1", "category 1", 1);
INSERT INTO entries VALUES (NULL, "key 2", "value 2", "category 1", 1);
INSERT INTO entries VALUES (NULL, "key 3", "value 3", "category 2", 1);