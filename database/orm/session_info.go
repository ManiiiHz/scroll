package orm

import (
	"context"
	"github.com/jmoiron/sqlx"
	"scroll-tech/common/types"
)

type sessionInfoOrm struct {
	db *sqlx.DB
}

var _ SessionInfoOrm = (*sessionInfoOrm)(nil)

// NewSessionInfoOrm create an sessionInfoOrm instance
func NewSessionInfoOrm(db *sqlx.DB) SessionInfoOrm {
	return &sessionInfoOrm{db: db}
}

func (o *sessionInfoOrm) GetSessionInfosByHashes(hashes []string) ([]*types.SessionInfo, error) {
	if len(hashes) == 0 {
		return nil, nil
	}
	query, args, err := sqlx.In("SELECT * FROM session_info WHERE hash IN (?);", hashes)
	if err != nil {
		return nil, err
	}
	rows, err := o.db.Queryx(o.db.Rebind(query), args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var sessionInfos []*types.SessionInfo
	for rows.Next() {
		var sessionInfo types.SessionInfo
		if err = rows.StructScan(&sessionInfo); err != nil {
			return nil, err
		}
		sessionInfos = append(sessionInfos, &sessionInfo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sessionInfos, nil
}

func (o *sessionInfoOrm) SetSessionInfo(rollersInfo *types.SessionInfo) error {
	sqlStr := "INSERT INTO session_info (task_id, roller_public_key, prove_type, roller_name, proving_status) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (hash, roller_public_key) DO UPDATE SET proving_status = EXCLUDED.proving_status;"
	_, err := o.db.Exec(sqlStr, rollersInfo.TaskID, rollersInfo.RollerPublicKey, rollersInfo.ProveType, rollersInfo.RollerName, rollersInfo.ProvingStatus)
	return err
}

// UpdateSessionInfoProvingStatus update the session info proving status
func (o *sessionInfoOrm) UpdateSessionInfoProvingStatus(ctx context.Context, dbTx *sqlx.Tx, hash string, pk string, status types.ProvingStatus) error {
	if _, err := dbTx.ExecContext(ctx, o.db.Rebind("update session_info set proving_status = ? where hash = ? and roller_public_key = ?;"), int(status), hash, pk); err != nil {
		return err
	}
	return nil
}
