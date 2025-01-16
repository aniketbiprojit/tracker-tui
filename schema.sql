CREATE TABLE client (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_name TEXT NOT NULL
);

CREATE TABLE projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    tracked_time_in_seconds INTEGER NOT NULL,
    client_id INTEGER NOT NULL,
    active BOOLEAN NOT NULL DEFAULT 0,
    FOREIGN KEY (client_id) REFERENCES client(id)
);

CREATE TABLE time_trackings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    project_id INTEGER NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

