package serialize

type Config struct {
	Port     string     `json:"port"`
	Requests []Requests `json:"requests"`
}

type Requests struct {
	Verb         string   `json:"verb"`
	Path         string   `json:"path"`
	ResponseType string   `json:"responseType"`
	Response     Response `json:"response"`
	Rules        []string `json:"rules,omitempty"`
	BodyType     string   `json:"bodyType,omitempty"`
	Body         any      `json:"body,omitempty"`
}

type Response struct {
	Body  any      `json:"body"`
	Rules []string `json:"rules"`
}
