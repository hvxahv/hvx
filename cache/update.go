package cache

import "github.com/hvxahv/hvx/errors"

// UPDATEAcct ...
func UPDATEAcct(username string, data []byte) error {
	if err := DELKey(username); err != nil {
		return errors.New("FAILED TO DELETE USER CACHE.")
	}
	if err := SETAcct(username, data, 0); err != nil {
		return errors.New("SYNC TO CACHE FAILED.")
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
