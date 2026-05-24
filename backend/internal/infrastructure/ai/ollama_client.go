package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OllamaClient struct {
	baseURL string
	model   string
}

func NewOllamaClient() *OllamaClient {
	return &OllamaClient{
		baseURL: os.Getenv("OLLAMA_URL"),
		model:   "qwen:3b", // ou deepseek-v2
	}
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func (c *OllamaClient) GenerateInsight(ctx context.Context, data string) (string, error) {
	prompt := fmt.Sprintf("Analise estes dados financeiros e dê um insight curto e acionável: %s", data)

	reqBody := OllamaRequest{
		Model:  c.model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, _ := json.Marshal(reqBody)
	resp, err := http.Post(c.baseURL+"/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ollamaResp OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return "", err
	}

	return ollamaResp.Response, nil
}
