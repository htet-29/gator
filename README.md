# RSS AGGREGATOR WITH GO!

## Requirement software to run this project

- Postgresql (Database)
- Go (program)

## How to install

```go
go install github.com/htet-29/gator
```

## Set up

1. Create `.gatorconfig.json` in your $home directory.
2. You need to register user first with
   - go run . register <username>
3. And then you can add feed with
   - go run . addfeed <title> <url>
4. Now you can aggregate feed with single command
   - go run . agg <time-between-aggregate>

## Avaliable Commands

- `go run . login <username>` for switch user
- `go run . users` for listing user
- `go run . feeds` for listing feeds
- `go run . follow` for following feed with current user
- `go run . unfollow` for unfollowing feed
- `go run . following` for listing all following feeds for current user
- `go run . browse` for browsing all of the posts that current user aggregate
- `go run . reset` for resetting database
