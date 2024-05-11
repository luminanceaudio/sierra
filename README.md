# Sierra

## Prerequisites

1. Install `protoc`:

```
brew install protobuf
```

2. Run go generate to generate the DB schema and compile the protobuf files:

```
go generate ./...
```

## Running the app

```
cd src/electron
npm start
```

## Useful commands

### Adding a new migration

```
cd src/electron
(cd services/sierra/internal/ && atlas migrate diff --dir-format golang-migrate initial  --dir "file://modules/sierradb/migrations" --to "ent://modules/sierradb/schema" --dev-url "sqlite://file?mode=memory&_fk=1")
```

### Rehashing the migrations dir

```
cd src/sierra
atlas migrate hash --dir "file://services/sierra/internal/modules/sierradb/migrations
```
