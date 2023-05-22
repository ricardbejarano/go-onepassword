package onepassword

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"strconv"
)

type Client struct {
	args []string
	path string
}

type ClientOptions struct {
	Account *string
	Cache   *bool
	Config  *string
	Session *string
}

func NewClient(opts *ClientOptions) (*Client, error) {
	args := []string{
		"op",
		"--format", "json",
		"--iso-timestamps",
		"--no-color",
	}
	if opts.Account != nil {
		args = append(args, "--account", *opts.Account)
	}
	if opts.Cache != nil {
		args = append(args, "--cache", strconv.FormatBool(*opts.Cache))
	}
	if opts.Config != nil {
		args = append(args, "--config", *opts.Config)
	}
	if opts.Session != nil {
		args = append(args, "--session", *opts.Session)
	}

	path, err := exec.LookPath(args[0])
	if err != nil {
		return nil, err
	}

	return &Client{
		args: args,
		path: path,
	}, nil
}

func (c *Client) runJson(data any, args ...string) error {
	output, err := c.runPlain(args...)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(output), data)
}

func (c *Client) runPlain(args ...string) (string, error) {
	buffer := new(bytes.Buffer)

	command := exec.Cmd{
		Path:   c.path,
		Args:   append(c.args, args...),
		Stdout: buffer,
	}

	if err := command.Run(); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
