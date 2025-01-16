
CREATE TRIGGER ensure_single_active_project_insert
AFTER INSERT ON projects
FOR EACH ROW
WHEN NEW.active = 1
BEGIN
  UPDATE projects
  SET active = 0
  WHERE id != NEW.id;
END;

CREATE TRIGGER ensure_single_active_project_update
AFTER UPDATE ON projects
FOR EACH ROW
WHEN NEW.active = 1
BEGIN
  -- Set active = 0 for all other rows
  UPDATE projects
  SET active = 0
  WHERE id != NEW.id;
END;
