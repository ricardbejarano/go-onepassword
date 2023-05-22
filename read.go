package onepassword

func (c *Client) Read(reference string) (string, error) {
	return c.runPlain("read", "--force", "--no-newline", reference)
}
