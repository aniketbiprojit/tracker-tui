## tracker-tui


Update TURSO_TOKEN in `.env` and run:

```sh
export (cat .env | xargs -n1)
atlas schema apply --env turso --to file://schema.sql --dev-url "sqlite://dev?mode=memory"
go run main.go
```
