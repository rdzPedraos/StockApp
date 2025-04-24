package gemini

type Prompt struct {
	Text string `json:"text"`
}

type PartsPrompt struct {
	Parts []Prompt `json:"parts"`
}

type GeminiRequest struct {
	Contents []PartsPrompt `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
	} `json:"candidates"`
}

func (s *service) Ask(question []string) (string, error) {
	path := "/v1beta/models/gemini-2.0-flash:generateContent?key=" + s.apiKey

	body := GeminiRequest{
		Contents: s.ConvertToPrompts(question),
	}

	var response GeminiResponse
	err := s.client.Post(path, body, &response)

	if err != nil {
		return "", err
	}

	return response.Candidates[0].Content.Parts[0].Text, nil
}

func (s *service) ConvertToPrompts(question []string) []PartsPrompt {
	var prompts []Prompt

	for _, q := range question {
		prompts = append(prompts, Prompt{Text: q})
	}

	return []PartsPrompt{
		{
			Parts: prompts,
		},
	}
}
