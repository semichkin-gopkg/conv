package converter

import "github.com/semichkin-gopkg/configurator"

type Params struct {
	Tag string
}

func WithTag(tag string) configurator.Updater[Params] {
	return func(p *Params) {
		p.Tag = tag
	}
}
