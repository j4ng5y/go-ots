package srv

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"k8s.io/klog"
)

// User is a general struct that models user information
type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	EmailAddress string `json:"email_address"`
}

// Secret is a general struct that models secrets information
type Secret struct {
	ID   string `json:"id"`
	Data struct {
		Secret     string    `json:"secret"`
		Passphrase string    `json:"passphrase"`
		ExpiresAt  time.Time `json:"expires_at"`
	} `json:"data"`
	Metadata struct {
		PassphraseRequired bool      `json:"passphrase_required"`
		CreatedAt          time.Time `json:"create_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		OpenedAt           time.Time `json:"opened_at"`
		DestroyedAt        time.Time `json:"destroyed_at"`
	} `json:"metadata"`
}

type createSecretRequest struct {
	Generate   bool   `json:"generate"`
	Secret     string `json:"secret"`
	Passphrase string `json:"passphrase"`
	ExpiresIn  string `json:"expires_at"`
}

func (req *createSecretRequest) unmarshal(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, req)
}

type createSecretResponse struct {
	User User   `json:"user"`
	Data Secret `json:"data"`
}

func (resp *createSecretResponse) marshal() ([]byte, error) {
	return json.Marshal(resp)
}

type getSecretRequest struct {
	ID string `json:"id"`
}

func (req *getSecretRequest) unmarshal(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, req)
}

type getSecretResponse struct {
	User User   `json:"user"`
	Data Secret `json:"data"`
}

func (resp *getSecretResponse) marshal() ([]byte, error) {
	return json.Marshal(resp)
}

type deleteSecretRequest struct {
	ID string `json:"id"`
}

func (req *deleteSecretRequest) unmarshal(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, req)
}

type deleteSecretResponse struct {
	User User   `json:"user"`
	Data Secret `json:"data"`
}

func (resp *deleteSecretResponse) marshal() ([]byte, error) {
	return json.Marshal(resp)
}

func (S *Server) createHandler(w http.ResponseWriter, r *http.Request) {
	var (
		req  = new(createSecretRequest)
		resp = new(createSecretResponse)
	)

	w.Header().Add("Content-Type", "application/json")

	if err := req.unmarshal(r.Body); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp.User.ID = "test"
	resp.User.Username = "test_user"
	resp.User.EmailAddress = "test@test.com"
	resp.Data.ID = "test"
	resp.Data.Data.Secret = "secret"
	resp.Data.Data.Passphrase = "passphrase"
	resp.Data.Data.ExpiresAt = time.Now().Add(24 * time.Hour)
	resp.Data.Metadata.PassphraseRequired = false
	resp.Data.Metadata.CreatedAt = time.Now()
	resp.Data.Metadata.UpdatedAt = time.Now()

	respJSON, err := resp.marshal()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respJSON)
	if err != nil {
		klog.Error(err)
	}
}

func (S *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	var (
		req  = new(getSecretRequest)
		resp = new(getSecretResponse)
	)

	w.Header().Add("Content-Type", "application/json")

	if err := req.unmarshal(r.Body); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp.User.ID = "test"
	resp.User.Username = "test_user"
	resp.User.EmailAddress = "test@test.com"
	resp.Data.ID = "test"
	resp.Data.Data.Secret = "secret"
	resp.Data.Data.Passphrase = "passphrase"
	resp.Data.Data.ExpiresAt = time.Now().Add(24 * time.Hour)
	resp.Data.Metadata.PassphraseRequired = false
	resp.Data.Metadata.CreatedAt = time.Now()
	resp.Data.Metadata.UpdatedAt = time.Now()

	respJSON, err := resp.marshal()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respJSON)
	if err != nil {
		klog.Error(err)
	}
}

func (S *Server) deleteHandler(w http.ResponseWriter, r *http.Request) {
	var (
		req  = new(deleteSecretRequest)
		resp = new(deleteSecretResponse)
	)

	w.Header().Add("Content-Type", "application/json")

	if err := req.unmarshal(r.Body); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp.User.ID = "test"
	resp.User.Username = "test_user"
	resp.User.EmailAddress = "test@test.com"
	resp.Data.ID = "test"
	resp.Data.Data.Secret = "secret"
	resp.Data.Data.Passphrase = "passphrase"
	resp.Data.Data.ExpiresAt = time.Now().Add(24 * time.Hour)
	resp.Data.Metadata.PassphraseRequired = false
	resp.Data.Metadata.CreatedAt = time.Now()
	resp.Data.Metadata.UpdatedAt = time.Now()

	respJSON, err := resp.marshal()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respJSON)
	if err != nil {
		klog.Error(err)
	}
}
