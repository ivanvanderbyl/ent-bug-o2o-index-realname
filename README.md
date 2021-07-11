# Ent Bug: Drop index when node name is too long

This repo reproduces the issue mentioned in ####

## Usage:

```shell
createdb ent_index_test
go run cmd/migrate/main.go // works
go run cmd/migrate/main.go // fails on second try
```
