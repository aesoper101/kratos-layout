// Code generated by ent, DO NOT EDIT.

package schemagen

import (
	"context"
	"fmt"
)

// Database is the client that holds all ent builders.
type Database struct {
	client *Client
}

// NewDatabase creates a new database configured with the given options.
func NewDatabase(opts ...Option) *Database {
	return &Database{client: NewClient(opts...)}
}

// InTx runs the given function f within a transaction.
func (db *Database) InTx(ctx context.Context, f func(context.Context) error) error {
	tx := TxFromContext(ctx)
	if tx != nil {
		return f(ctx)
	}

	tx, err := db.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	if err = f(NewTxContext(ctx, tx)); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return fmt.Errorf("rolling back transaction: %v (original error: %w)", err2, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func (db *Database) loadClient(ctx context.Context) *Client {
	tx := TxFromContext(ctx)
	if tx != nil {
		return tx.Client()
	}
	return db.client
}

// User is the client for interacting with the User builders.
func (db *Database) User(ctx context.Context) *UserClient {
	return db.loadClient(ctx).User
}