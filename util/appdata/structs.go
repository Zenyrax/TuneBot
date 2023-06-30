package appdata

type Config struct {
	Webhook string `json:"webhook"`
	License string `json:"license"` // This line is useless for this project!

	Version string `json:"-"`
}
