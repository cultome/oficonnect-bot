package oficonnectbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type Bot struct {
	OfiConnectID string
	Client       *Client
}

type MarshalInformation struct {
	ID             int    `json:"id"`
	CountryID      string `json:"id_pais"`
	Name           string `json:"nombres"`
	LastName       string `json:"apellido_paterno"`
	MotherLastName string `json:"apellido_materno"`
	Role           string `json:"role"`
	Email          string `json:"email"`
	Birthday       string `json:"nacimiento"`
	Sex            string `json:"sexo"`
	HomePhone      string `json:"tel_domicilio"`
	PersonalPhone  string `json:"tel_personal"`
	English        string `json:"ingles"`
	Resident       int    `json:"residente"`
	PhotoURL       string `json:"url_foto"`
	Languages      string `json:"idiomas"`
	CreatedAt      string `json:"created_at"`
	BloodType      string `json:"tipo_sangre"`
	Allergies      string `json:"alergias"`
	Level          *Level
}

type Level struct {
	ID          int    `json:"id"`
	UserLevelID int    `json:"id_usuario_nivel"`
	LevelID     int    `json:"id_nivel"`
	Type        string `json:"tipo"`
	Name        string `json:"nombre"`
	Stars       int    `json:"estrellas"`
}

type PersonalInformationResponse struct {
	Status              string                `json:"status"`
	PersonalInformation []*MarshalInformation `json:"datos_personales"`
	Level               []*Level              `json:"nivel"`
}

type Event struct {
	ID           int    `json:"id"`
	Active       int    `json:"activo"`
	Open         int    `json:"abierto"`
	EventID      int    `json:"id_evento"`
	EventName    string `json:"nombre_evento"`
	CategoryName string `json:"nombre_categoria"`
	DivisionName string `json:"nombre_division"`
	Confimed     int    `json:"confirmado"`
	Quota        string `json:"cupo"`
	FileURL      string `json:"url_file"`
	Role         string `json:"role"`
	PositionID   int    `json:"id_puesto"`
	PositionName string `json:"nombre_puesto"`
	UserID       int    `json:"id_usuario"`
	UpdatedAt    string `json:"fecha_actualizado"`
	CreatedAt    string `json:"fecha_creado"`
}

type EventsResponse struct {
	Status string   `json:"status"`
	Events []*Event `json:"eventos"`
}

type RegistrationRequest struct {
	ID       int    `json:"id"`
	EventID  int    `json:"id_evento"`
	UserID   int    `json:"id_usuario"`
	Confimed string `json:"confirmado"`
}

type RegistrationResponse struct {
	Status         string `json:"status"`
	Message        string `json:"message"`
	Limit          int    `json:"limite"`
	TotalConfirmed int    `json:"total_confirmados"`
}

type Confirms struct {
	Count string `json:"confirmados"`
}

type QuotaResponse struct {
	Status        string      `json:"status"`
	Confirmations []*Confirms `json:"confirmados"`
}

func BuildBot(oficonnectID string) *Bot {
	return &Bot{
		OfiConnectID: oficonnectID,
		Client:       BuildClient(),
	}
}

func (b *Bot) RetriveEvents() ([]*Event, error) {
	url := fmt.Sprintf("https://api.oficonnect.omdai.org/public/auth/eventos-usuario/obtener/%s", b.OfiConnectID)
	var response EventsResponse

	err := b.Client.Get(url, &response)

	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	if response.Status != "success" {
		return nil, fmt.Errorf("[-] Failed response: %+v", response)
	}

	return response.Events, nil
}

func (b *Bot) RegisterForEvent(evt *Event) (*RegistrationResponse, error) {
	url := "https://api.oficonnect.omdai.org/public/auth/eventos-usuario/evento/confirmar"
	var response RegistrationResponse

	payload, err := json.Marshal(RegistrationRequest{
		ID:       evt.ID,
		EventID:  evt.EventID,
		UserID:   evt.UserID,
		Confimed: "1",
	})

	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	err = b.Client.Post(url, bytes.NewBuffer(payload), &response)

	return &response, err
}

func (b *Bot) RetrivePersonalInformation() (*MarshalInformation, error) {
	url := fmt.Sprintf("https://api.oficonnect.omdai.org/public/auth/datos-personales/obtener/%s", b.OfiConnectID)
	var response PersonalInformationResponse

	err := b.Client.Get(url, &response)

	if err != nil {
		return nil, fmt.Errorf("[-] Unable to read body: %s", err.Error())
	}

	info := response.PersonalInformation[0]
	info.Level = response.Level[0]

	return info, nil
}

func (b *Bot) RetriveConfirmationsByEvent(eventID int) (int, error) {
	url := fmt.Sprintf("https://api.oficonnect.omdai.org/public/auth/eventos-usuario/confirmados/obtener/%d", eventID)
	var response QuotaResponse

	err := b.Client.Get(url, &response)

	if err != nil {
		return -1, fmt.Errorf("[-] Unable to parse response body: %s", err.Error())
	}

	count, err := strconv.Atoi(response.Confirmations[0].Count)

	if err != nil {
		return -1, fmt.Errorf("[-] Invalid number in confirmations: %s", err.Error())
	}

	return count, nil
}

// # documentos
// 'https://api.oficonnect.omdai.org/public/auth/datos-personales/avisos/obtener/#{oficonnectID}'
// # cursos
// 'https://api.oficonnect.omdai.org/public/auth/cursos-usuario/obtener-todos/#{oficonnectID}/normal'
