package preguntas

type Preguntas struct {
  Date_created string   `json:"date_created"`
	Item_id string        `json:"item_id"`
	Text string           `json:"text"`
	Status string         `json:"status"`
 	Answer string         `json:"answer"`
}
type QuestionsMeli struct {
	Questions []QuestionMeli  `json:"questions"`
}
