# gorker
gorker node for the crawling

## TOC

### cli func
- config - config.d, gorker.conf, tls
- help
- info/graph - get topology and infrastructure info (redis, schema)
- env, GORKER_ENV_VARS = token, set per func, debug
- ui - http server, access token, default on pause,
- exec - execute crawling - main.go
- mode - broker, single, worker
- blades - get plugin, download, list, upgrade
- repo - set repo for downloads sign and plugins
- fetch - signatures
- scan type - full, inc, validation, vacuum,   
- data , set, get, mark - access to data, client mode
- put - data load
- set - manage gorkers
- test
- import - db/data transfer

features:
- add: webhooks
- add: grpc
- add: cron, data shape
