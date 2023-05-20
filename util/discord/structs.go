package discord

type Webhook struct {
	Embeds   []Embed   `json:"embeds"`
	Content  string     `json:"content"`
}
type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}
type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type Image struct {
	URL string `json:"url"`
}
type Thumbnail struct {
	URL string `json:"url"`
}
type Embed struct {
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Color     int       `json:"color"`
	Fields    []Field  `json:"fields"`
	Author    Author    `json:"author,omitempty"`
	Footer    Footer    `json:"footer,omitempty"`
	// Timestamp time.Time `json:"timestamp,omitempty"`
	Image     Image     `json:"image,omitempty"`
	Thumbnail Thumbnail `json:"thumbnail,omitempty"`
}
