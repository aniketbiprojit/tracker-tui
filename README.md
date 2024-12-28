## tracker-tui


Update TURSO_TOKEN in `.env` and run:

```sh
export (cat .env | xargs -n1)
atlas schema apply --env turso --to file://schema.sql --dev-url "sqlite://dev?mode=memory"                                                 14:02:02
go run main.go
```
