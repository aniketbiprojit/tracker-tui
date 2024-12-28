schema "main" {
}

table "projects" {
  schema = schema.main
  column "id" {
    type = int
  }
  column "name" {
    type = varchar(255)
  }
  column "client_id" {
    type = int
  }
  primary_key {
    columns = [
      column.id
    ]
  }
  foreign_key "client_fk" {
    columns = [column.client_id]
    ref_columns = [column.id]
  }
}
