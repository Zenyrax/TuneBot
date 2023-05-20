package appdata

type Config struct {
	Webhook string `json:"webhook"`
	License string `json:"license"`

	Version string `json:"-"`
}
