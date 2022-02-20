package cache

import "github.com/pkg/errors"

// UPDATEAcct ...
func UPDATEAcct(username string, data []byte) error {
	if err := DELKey(username); err != nil {
		return errors.Errorf("FAILED TO DELETE USER CACHE.")
	}
	if err := SETAcct(username, data, 0); err != nil {
		return errors.Errorf("SYNC TO CACHE FAILED.")
	}
	return nil
}

func UPDATEMail(mail, newMail string) error {
	if err := DELAcctMail(mail); err != nil {
		return err
	}
	if err := SETAcctMail(newMail); err != nil {
		return err
	}
	return nil
}
