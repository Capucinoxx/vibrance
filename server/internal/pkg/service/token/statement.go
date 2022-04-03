package token

const (
	StmtCreateToken = `
		INSERT INTO tokens (hash, client_key, client_secret, scopes)
		VALUES (?, ?, ?, ?)
			USING TTL ?
	`

	StmtDeleteToken = `
		DELETE FROM tokens
			WHERE hash = ?
	`

	StmtFindToken = `
		SELECT hash, client_key, client_secret, scopes
		FROM tokens
			WHERE hash = ?
	`
)
