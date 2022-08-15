package internal

func init() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	// Initialize the database.
	if err := cockroach.NewRoach().Dial(); err != nil {
		fmt.Println(err)
		return
	}
}

func TestIsExist() {
	a := NewActorsIsExist("example.com", "alice")
	if a.IsExist() {
		t.Log("Actor exists")
	} else {
		t.Log("Actor does not exist")
	}
}

func TestCreate() {
	a := NewActors("alice", "alice", "Person")
	if err := a.Create(); err != nil {
		t.Error(err)
	}
}

func TestGet() {
	a := NewActorsId(1)
	if err := a.Get(); err != nil {
		t.Error(err)
	}
}

func TestGetActorsByPreferredUsername() {
	a := NewActorsPreferredUsername("alice", "example.com")
	if err := a.GetActorsByPreferredUsername(); err != nil {
		t.Error(err)
	}
}

func AddActor() {
	a := NewAddActors("alice", "alice", "Person")
	if err := a.AddActor(); err != nil {
		t.Error(err)
	}
}

func TestGetActorByUsername() {
	a := NewActorsPreferredUsername("alice", "example.com")
	if err := a.GetActorByUsername(); err != nil {
		t.Error(err)
	}
}

func TestEdit() {
	a := NewActorsId(1)
	if err := a.Edit(); err != nil {
		t.Error(err)
	}
}

func TestDelete() {
	a := NewActorsId(1)
	if err := a.Delete(); err != nil {
		t.Error(err)
	}
}
