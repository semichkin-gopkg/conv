package converter

import (
	"github.com/mitchellh/mapstructure"
	"github.com/semichkin-gopkg/configurator"
)

func Convert[T any](
	source any,
	updaters ...configurator.Updater[Params],
) (T, error) {
	var target T

	config := configurator.New[Params]().
		Append(func(c *Params) {
			c.Tag = "json"
		}).
		Append(updaters...).
		Apply()

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		TagName:  config.Tag,
		Result:   &target,
	})
	if err != nil {
		return target, err
	}

	if err = decoder.Decode(source); err != nil {
		return target, err
	}

	return target, nil
}
