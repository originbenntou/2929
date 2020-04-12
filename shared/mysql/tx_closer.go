package mysql

func CloseTransaction(tx TxManager, err error) error {
	var txErr error
	if recover() != nil {
		txErr = tx.Rollback()
	} else if err != nil {
		txErr = tx.Rollback()
	} else {
		txErr = tx.Commit()
	}

	if txErr != nil {
		err = txErr
	}

	return err
}
