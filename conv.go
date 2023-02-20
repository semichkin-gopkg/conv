package conv

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/semichkin-gopkg/conf"
)

func Struct[T any](
	source any,
	configurators ...conf.Updater[mapstructure.DecoderConfig],
) (T, error) {
	var target T

	config := conf.NewBuilder[mapstructure.DecoderConfig]().
		Fix(func(c *mapstructure.DecoderConfig) {
			c.Result = &target
		}).
		Append(func(c *mapstructure.DecoderConfig) {
			c.Metadata = nil
			c.TagName = "json"
		}).
		Append(configurators...).
		Build()

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return target, err
	}

	if err = decoder.Decode(source); err != nil {
		return target, err
	}

	return target, nil
}

func JSON(data any) ([]byte, error) {
	return json.Marshal(data)
}

func FromJSON[T any](bytes []byte) (T, error) {
	var data T
	return data, json.Unmarshal(bytes, &data)
}

func Pointer[T any](t T) *T {
	return &t
}
