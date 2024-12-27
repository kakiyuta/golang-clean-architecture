package db

type LocalConnector struct{}

// NewLocalConnector ローカルコネクタを生成する
func NewLocalConnector() *LocalConnector {
	return &LocalConnector{}
}

// Begin トランザクションを開始する
func (c *LocalConnector) Begin() error {
	return nil
}

// Commit トランザクションをコミットし、接続を閉じる
func (c *LocalConnector) Commit() error {
	return nil
}

// Rollback トランザクションをロールバックし、接続を閉じる
func (c *LocalConnector) Rollback() {}
