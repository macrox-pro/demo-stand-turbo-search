## Usage

```
Usage:
  server [flags]
  server [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  index:init  create search index
  index:sync  sync search index

Flags:
  -h, --help                help for server
  -i, --index:path string   index file (default is ./index.bleve) (default "./index.bleve")

Use "server [command] --help" for more information about a command.
```

### Create search index

```shell
server index:init
```

### Sync search index

```shell
server index:sync
```

### Run server

```shell
server
```