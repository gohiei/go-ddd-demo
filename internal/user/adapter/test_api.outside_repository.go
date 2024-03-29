package adapter

import (
	"io"
	"net/http"

	repo "cypt/internal/user/repository"
)

func NewTestAPIOutsideRepository(client *http.Client) *TestAPIOutsideRepository {
	return &TestAPIOutsideRepository{client: client}
}

type TestAPIOutsideRepository struct {
	client *http.Client
}

var _ repo.OutsideRepository = (*TestAPIOutsideRepository)(nil)

func (r *TestAPIOutsideRepository) GetEchoData() (string, error) {
	// jsonString := `{"username": "test@homuchen.com", "password": "homuchen"}`
	// body := bytes.NewReader([]byte(jsonString))

	resp, err := r.client.Get(
		"https://api.publicapis.org/categories",
	)

	if err != nil {
		return "", err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	data := string(b)

	return data, nil
}
