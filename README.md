# plextools


## rename
### rename all the files on disk to a standard format
usage:
```go run cmd/rename/main.go <plex addr:port> [rename]```

#### notes
* arg 2 (rename) is used to specify if you wish to actually rename the files instead of just show what the rename would be
* you should wait for plex to re-index the files before you run again or you might get errors renaming