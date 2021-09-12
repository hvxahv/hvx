package channel

//
//func TestAddAdmin(t *testing.T) {
//	TestInitChannelConfig(t)
//	nmc, err := NewAddAdmins("hvturingga", "Ja5QZv-fgxhg182", "alice")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	admin, s, err := nmc.AddAdmin()
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Printf("code: %v; message: %s", admin, s)
//
//	nmc2, err := NewAddAdmins("alice", "f6574uSSqGQ7CJX", "hvturingga")
//	admin2, s2, err := nmc2.AddAdmin()
//	if err != nil {
//		return
//	}
//	fmt.Printf("code: %v; message: %s", admin2, s2)
//
//}

//func TestNewAdmins(t *testing.T) {
//	TestInitChannelConfig(t)
//	nca, err := NewAddAdmins("alice", "f6574uSSqGQ7CJX", "bob")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(nca)
//
//}
//
//func TestGetListByName(t *testing.T) {
//	TestInitChannelConfig(t)
//	name := NewAdminsByName("hvturingga")
//	_, cls, _ := name.GetListByName()
//	fmt.Println(cls)
//
//	for _, i := range cls {
//		fmt.Println(i.Name)
//	}
//
//}
//
//func TestGetAdmListByID(t *testing.T) {
//	TestInitChannelConfig(t)
//	adm := NewAdminsByID("hvturingga_chan")
//	code, accounts, err := adm.GetAdmLisByID()
//	if err != nil {
//		return
//	}
//
//	fmt.Println(code, accounts)
//
//}
//
//func TestAdmins_RemoveAdmin(t *testing.T) {
//	TestInitChannelConfig(t)
//	na := NewAdminsByName("bob")
//	code, res, err := na.RemoveAdmin()
//	if err != nil {
//		return
//	}
//
//	fmt.Println(code, res)
//}