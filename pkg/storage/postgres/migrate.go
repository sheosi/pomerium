package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"

	"github.com/pomerium/pomerium/pkg/cryptutil"
)

var migrations = []func(context.Context, pgx.Tx) error{
	1: func(ctx context.Context, tx pgx.Tx) error {
		_, err := tx.Exec(ctx, `
			CREATE TABLE `+schemaName+`.`+recordsTableName+` (
				type TEXT NOT NULL,
				id TEXT NOT NULL,
				version BIGINT NOT NULL,
				data BYTEA NOT NULL,
				modified_at TIMESTAMPTZ NOT NULL DEFAULT(NOW()),

				index_cidr INET NULL,

				PRIMARY KEY (type, id)
			)
		`)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, `
			CREATE INDEX ON `+schemaName+`.`+recordsTableName+`
			USING gist (index_cidr inet_ops);
		`)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, `
			CREATE TABLE `+schemaName+`.`+recordChangesTableName+` (
				type TEXT NOT NULL,
				id TEXT NOT NULL,
				version BIGSERIAL,
				data BYTEA NOT NULL,
				modified_at TIMESTAMPTZ NOT NULL,
				deleted_at TIMESTAMPTZ NULL,

				PRIMARY KEY (version)
			)
		`)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, `
			CREATE TABLE `+schemaName+`.`+recordOptionsTableName+` (
				type TEXT NOT NULL,
				capacity BIGINT NULL,

				PRIMARY KEY (type)
			)
		`)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, `
			CREATE TABLE `+schemaName+`.`+leasesTableName+` (
				name TEXT NOT NULL,
				id TEXT NOT NULL,
				expires_at TIMESTAMPTZ NOT NULL,

				PRIMARY KEY (name)
			)
		`)
		if err != nil {
			return err
		}

		return nil
	},
}

func migrate(ctx context.Context, tx pgx.Tx) (serverVersion uint64, err error) {
	_, err = tx.Exec(ctx, `CREATE SCHEMA IF NOT EXISTS `+schemaName)
	if err != nil {
		return serverVersion, err
	}

	_, err = tx.Exec(ctx, `
			CREATE TABLE IF NOT EXISTS `+schemaName+`.`+migrationInfoTableName+` (
				server_version BIGINT NOT NULL,
				migration_version SMALLINT NOT NULL
			)
		`)
	if err != nil {
		return serverVersion, err
	}

	var migrationVersion uint64
	err = tx.QueryRow(ctx, `
			SELECT server_version, migration_version
			  FROM `+schemaName+`.migration_info
		`).Scan(&serverVersion, &migrationVersion)
	if errors.Is(err, pgx.ErrNoRows) {
		serverVersion = uint64(cryptutil.NewRandomUInt32()) // we can't actually store a uint64, just an int64, so just generate a uint32
		_, err = tx.Exec(ctx, `
				INSERT INTO `+schemaName+`.`+migrationInfoTableName+` (server_version, migration_version)
				VALUES ($1, $2)
			`, serverVersion, 0)
	}
	if err != nil {
		return serverVersion, err
	}

	for version := migrationVersion + 1; version < uint64(len(migrations)); version++ {
		err = migrations[version](ctx, tx)
		if err != nil {
			return serverVersion, err
		}
		_, err = tx.Exec(ctx, `
				UPDATE `+schemaName+`.`+migrationInfoTableName+`
				SET migration_version = $1
			`, version)
		if err != nil {
			return serverVersion, err
		}
	}

	return serverVersion, nil
}