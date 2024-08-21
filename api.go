package oficonnectbot

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct{}

type EventsResponse struct {
	Status string   `json:"status"`
	Events []*Event `json:"eventos"`
}

type RegistrationRequest struct {
	ID       string `json:"id"`
	EventID  string `json:"id_evento"`
	UserID   string `json:"id_usuario"`
	Confimed string `json:"confirmado"`
}

type RegistrationResponse struct {
	Status         string `json:"status"`
	Limit          int    `json:"limite"`
	TotalConfirmed int    `json:"total_confirmados"`
}

func BuildClient() *Client {
	return &Client{}
}

func (c *Client) Get(url string) ([]byte, error) {
	return request(url, "GET", nil)
}

func (c *Client) Post(url string, payload io.Reader) ([]byte, error) {
	return request(url, "POST", payload)
}

func request(url, method string, payload io.Reader) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(string(method), url, payload)

	if err != nil {
		return nil, fmt.Errorf("[-] Unable to create request: %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OFICONNECT_TOKEN")))

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("[-] Problems with request: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return body, err
}

// func (b *Bot) RetrivePersonalInformation() *PersonalInformation {
//
// }

// # documentos
// 'https://api.oficonnect.omdai.org/public/auth/datos-personales/avisos/obtener/#{oficonnect_id}'
//
// # datos personales
// 'https://api.oficonnect.omdai.org/public/auth/datos-personales/obtener/#{oficonnect_id}'
//
// # confirmados por evento?
// 'https://api.oficonnect.omdai.org/public/auth/eventos-usuario/confirmados/obtener/#{event_id}'
//
// # cursos
// 'https://api.oficonnect.omdai.org/public/auth/cursos-usuario/obtener-todos/#{oficonnect_id}/normal'
