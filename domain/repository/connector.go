package repository

// Connector コネクションコントローラ
type Connector interface {
	// Begin トランザクションを開始する
	Begin() error

	// Commit トランザクションをコミットし、接続を閉じる
	Commit() error

	// Rollback トランザクションをロールバックし、接続を閉じる
	Rollback()
}
