package sierradb

//go:generate go run -mod=mod entgo.io/ent/cmd/ent --feature sql/versioned-migration --feature sql/upsert generate ./schema --target ./internal/sierraent
