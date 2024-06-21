package segments

import (
	"github.com/jandedobbeleer/oh-my-posh/src/platform"
	"github.com/jandedobbeleer/oh-my-posh/src/properties"
)

type Nix struct {
	props properties.Properties
	env   platform.Environment

	Purity string
}

func (n *Nix) Enabled() bool {
	inNixShell := n.env.Getenv("IN_NIX_SHELL")
	if inNixShell != "" {
		n.Purity = inNixShell
		return true
	}
	return false
}

func (n *Nix) Template() string {
	return "{{ .Purity }} "
}

func (n *Nix) Init(props properties.Properties, env platform.Environment) {
	n.props = props
	n.env = env
}
