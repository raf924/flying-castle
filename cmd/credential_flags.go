package cmd

type CredentialFlags struct {
	UserName string `flag:"name" required:"true"`
	Password string `flag:"password" required:"true"`
	ApiKey   string `flag:"key"`
}

func (c *CredentialFlags) Validate() {
	if c.Password == "" && c.ApiKey == "" {
		panic("no valid credentials were given")
	}
}
