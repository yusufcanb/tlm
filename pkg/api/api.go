package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/yusufcanb/tlama/pkg/config"
)

type OllamaAPI struct {
	Config *config.TlamaConfig
}

func (o *OllamaAPI) PullModel(model string) {
	payload := pullModelRequestPayload{
		Name:   model,
		Stream: false,
	}

	jsonBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/pull", o.Config.Llm.Host), bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatal("Error creating HTTP request: ", err.Error())
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making HTTP request: ", err.Error())
	}

	// Read the response header
	fmt.Println("Response: Content-length:", resp.Header.Get("Content-length"))

	bytesRead := 0
	buf := make([]byte, 128)

	// Read the response body
	for {
		n, err := resp.Body.Read(buf)
		bytesRead += n

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Error reading HTTP response: ", err.Error())
		}
	}

	fmt.Println("Response: Read", bytesRead, "bytes")

}

func (o *OllamaAPI) Generate(prompt string) (string, error) {
	builder := strings.Builder{}
	builder.WriteString(prompt)
	builder.WriteString(fmt.Sprintf(". I'm using %s terminal", o.Config.Shell))
	builder.WriteString(fmt.Sprintf("on operating system: %s", runtime.GOOS))

	payload := generateRequestPayload{
		Model:  o.Config.Llm.Model,
		System: `You are software program specifically for Command Line Interface usage. User will ask you some thing that can be convertible to a UNIX or Windows command. You won't provide information or explanations and your output will be just an executable shell command inside three backticks.`,
		Prompt: builder.String(),
		Stream: false,
		Options: options{
			Temperature: o.Config.Llm.Parameters.Temperature,
			TopP:        o.Config.Llm.Parameters.TopP,
		},
	}

	jsonBytes, _ := json.Marshal(payload)

	resp, err := http.Post(fmt.Sprintf("%s/api/generate", o.Config.Llm.Host), "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	response := generateResponsePayload{}
	json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	retval := strings.Replace(response.Response, "```bash", "", -1)
	retval = strings.Replace(retval, "```", "", -1)
	retval = strings.Replace(retval, "\n", "", -1)

	return retval, nil
}

func New(cfg *config.TlamaConfig) *OllamaAPI {
	return &OllamaAPI{
		Config: cfg,
	}
}
