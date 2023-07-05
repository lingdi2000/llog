package llog

type Config struct {
	Level  Level  `json:"level"`
	Mode   Mode   `json:",default=console,options=[console,file,console|file]"`
	Path   string `json:",default=logs"`
	Skip   int    `json:"./default=2"`
	Suffix string `json:"suffix"`
}
