# smash_sp
## Commands
```bash
# start server
go run main.go
```


```bash
# search videos
go run cobra/command.go search video_title
```

```bash
# migration
migrate -source file://db/migrations -database 'mysql://smash_sp:password@tcp(127.0.0.1:3306)/smash_sp' up
```
