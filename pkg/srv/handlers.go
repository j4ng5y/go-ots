package srv

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type createSecretRequest struct {
	Secret     string `json:"secret"`
	Passphrase string `json:"passphrase"`
	TTL        int    `json:"ttl"`
	Recipient  string `json:"recipient"`
}

func (req *createSecretRequest) unmarshal(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, req)
}

type createSecretResponse struct {
	UserID             string    `json:"user_id"`
	MetadataKey        string    `json:"metadata_key"`
	SecretKey          string    `json:"secret_key"`
	TTL                int       `json:"ttl"`
	MetadataTTL        int       `json:"metadata_ttl"`
	SecretTTL          int       `json:"secret_ttl"`
	Recipient          string    `json:"recipient"`
	CreateAt           time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	PassphraseRequired bool      `json:"passphrase_required"`
}

func (resp *createSecretResponse) marshal() ([]byte, error) {
	return json.Marshal(resp)
}

type generateSecretRequest struct{}

func (req *generateSecretRequest) unmarshal(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, req)
}

type generateSecretResponse struct{}

func (resp *generateSecretResponse) marshal() ([]byte, error) {
	return json.Marshal(resp)
}

type retrieveSecretRequest struct{}

func (req *retrieveSecretRequest) unmarshal(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, req)
}

type retrieveSecretResponse struct{}

func (resp *retrieveSecretResponse) marshal() ([]byte, error) {
	return json.Marshal(resp)
}

type retrieveMetadataRequest struct{}

func (req *retrieveMetadataRequest) unmarshal(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, req)
}

type retrieveMetadataResponse struct{}

func (resp *retrieveMetadataResponse) marshal() ([]byte, error) {
	return json.Marshal(resp)
}

type burnSecretRequest struct{}

func (req *burnSecretRequest) unmarshal(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, req)
}

type burnSecretResponse struct{}

func (resp *burnSecretResponse) marshal() ([]byte, error) {
	return json.Marshal(resp)
}

func (S *Server) createHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func (S *Server) generateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func (S *Server) retrieveHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func (S *Server) retrieveMetadataHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func (S *Server) burnHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
