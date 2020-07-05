Custom tool to convert a Joomla site to a markdown one.

Usage:

-   start the containers reproducing the Joomla site env (at least the db)
-   npm install (to get the html2md tool)
-   `go run main.go` will extract stuff from the MySQL database
-   look at output in `content` directory, tweak main.go and repeat
