package internal

import "github.com/hvxahv/hvx/cfg"

func init() {
	cfg.Default()
}

// func TestSaved_Create(t *testing.T) {
// 	d := v1alpha1.CreateSavedRequest{
// 		AccountId:   "733124680636596225",
// 		Name:        "YUI",
// 		Description: "description",
// 		Cid:         "QmVgBz2p2P3PnfiicJUHpyPVaiXDCNAKBnhimF9rP8c2zD",
// 		Types:       "jpeg/png",
// 		IsPrivate:   false,
// 	}

// 	s := saved{}
// 	create, err := s.CreateSaved(context.Background(), &d)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	fmt.Println(create)
// }

// func TestSaved_GetSaves(t *testing.T) {
// 	s := saved{}
// 	saves, err := s.GetSaves(context.Background(), &v1alpha1.GetSavesRequest{
// 		AccountId: "733124680636596225",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	fmt.Println(saves)
// }

// func TestSaved_GetSaved(t *testing.T) {
// 	s := saved{}
// 	save, err := s.GetSaved(context.Background(), &v1alpha1.GetSavedRequest{
// 		AccountId: "733124680636596225",
// 		Id:        "738165894748110849",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	fmt.Println(save)
// }

// func TestSaved_EditSaved(t *testing.T) {
// 	d := v1alpha1.EditSavedRequest{
// 		Id:          "738165894748110849",
// 		Name:        "",
// 		Description: "",
// 	}

// 	s := saved{}
// 	edit, err := s.EditSaved(context.Background(), &d)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	fmt.Println(edit)
// }
