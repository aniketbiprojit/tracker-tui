CREATE TABLE projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name INTEGER NOT NULL,
    tracked_time_in_seconds INTEGER NOT NULL
);

CREATE TABLE time_tracking (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    project_id INTEGER NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

CREATE TABLE client {}
