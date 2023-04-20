package conv

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/semichkin-gopkg/conf"
)

func p[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

func Struct[T any](
	source any,
	updaters ...conf.Updater[mapstructure.DecoderConfig],
) (T, error) {
	var target T

	config := conf.New[mapstructure.DecoderConfig]().
		Fix(func(c *mapstructure.DecoderConfig) {
			c.Result = &target
		}).
		Append(func(c *mapstructure.DecoderConfig) {
			c.Metadata = nil
			c.TagName = "json"
		}).
		Append(updaters...).
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

func MustStruct[T any](
	source any,
	updaters ...conf.Updater[mapstructure.DecoderConfig],
) T {
	return p(Struct[T](source, updaters...))
}

func JSON(source any) ([]byte, error) {
	return json.Marshal(source)
}

func MustJSON[T any](source any) []byte {
	return p(JSON(source))
}

func Dbg(source any) string {
	return string(p(json.MarshalIndent(source, "", "  ")))
}

func FromJSON[T any](bytes []byte) (T, error) {
	var data T
	return data, json.Unmarshal(bytes, &data)
}

func MustFromJSON[T any](bytes []byte) T {
	return p(FromJSON[T](bytes))
}

func Pointer[T any](t T) *T {
	return &t
}
