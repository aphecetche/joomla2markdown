Custom tool to convert a Joomla site to a Hugo one.

Usage:

- start the containers reproducing the Joomla site env (at least the db)
- `go run main.go` will extract stuff from the MySQL database
- look at output in `content` directory, tweak main.go and repeat

