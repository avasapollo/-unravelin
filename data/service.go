package data

type parser struct {
}

func NewParser() Parser {
	return &parser{}
}

func (p parser) ParseMapToData(input map[string]interface{}) (*Data, error) {
	if input == nil {
		return nil, InputNotValidError
	}

	output := new(Data)
	// WebsiteUrl
	if s, ok := input["websiteUrl"].(string); ok {
		output.WebsiteUrl = s
	}
	// SessionId
	if s, ok := input["sessionId"].(string); ok {
		output.SessionId = s
	}

	// CopyAndPaste
	if m, ok := input["copiedAndPaste"].(map[string]bool); ok {
		output.CopyAndPaste = m
	}
	// FormCompletionTime
	if m, ok := input["time"].(int); ok {
		output.FormCompletionTime = m
	}
	// ResizeFrom
	output.ResizeFrom = getDimension(input["resizeFrom"])
	// ResizeTo
	output.ResizeTo = getDimension(input["resizeTo"])

	// TODO: I would like to add e validation of data here and return an error if doesn't contain some fields
	return output, nil
}

func getDimension(input interface{}) Dimension {
	if input == nil {
		return Dimension{}
	}

	m, ok := input.(map[string]string)
	if !ok {
		return Dimension{}
	}

	return Dimension{
		Width:  m["width"],
		Height: m["height"],
	}
}
