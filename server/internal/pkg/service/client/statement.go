package client

const (
	StmtCreateClient = `
		INSERT INTO clients (id, key, secret, created_at, deleted_at)
		VALUES (?, ?, ?, ?, ?) IF NOT EXISTS
	`

	StmtFindClient = `
		SELECT id, key, secret, created_at, deleted_at
		FROM clients
			WHERE key = ?
	`

	StmtUpdateSecretClient = `
		UPDATE clients
		SET secret = ?
			WHERE key = ? IF EXISTS
	`

	StmtSoftDeleteClient = `
		UPDATE clients
		SET deleted_at = ?
			WHERE key = ? IF EXISTS
	`

	StmtDeleteClient = `
		DELETE FROM clients
			WHERE key = ? IF EXISTS
	`
)
