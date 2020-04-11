package mysql

func CloseTransaction(tx TxManager, err error) error {
	if recover() != nil {
		err = tx.Rollback()
	} else if err != nil {
		err = tx.Rollback()
	} else {
		err = tx.Commit()
	}

	return err
}
