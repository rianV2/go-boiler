// +build unit

package controller_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/icrowley/fake"
	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	err := config.Load()
	if err != nil {
		fmt.Printf("Config error: %s\n", err.Error())
		os.Exit(1)
	}

	err = initLogging()
	if err != nil {
		fmt.Printf("Logging error: %s\n", err.Error())
		os.Exit(1)
	}

	retCode := m.Run()
	os.Exit(retCode)
}

func initLogging() error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.StandardLogger()
	level, err := logrus.ParseLevel(config.Instance().LogLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)

	return err
}

func setupHttpHandler(t *testing.T) http.Handler {
	server := web.NewHttpServer(nil, config.Instance(), &mocks.MessageBus{}, &mocks.Account{}, &mocks.Queue{}, &mocks.Limit{}, &mocks.Worker{})
	handler, err := server.GetHandler()
	if err != nil {
		t.Fatal(err.Error())
	}
	return handler
}

func setupHttpServerBuilder(t *testing.T) *web.HttpServerBuilder {
	builder := web.NewHttpServerBuilder(nil, config.Instance(), &mocks.MessageBus{}, &mocks.Account{}, &mocks.Queue{}, &mocks.Limit{}, &mocks.Worker{})
	return builder
}

func performRequestJSON(r http.Handler, method string, path string, body interface{}, headers map[string]string) (*httptest.ResponseRecorder, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return performRequest(r, method, path, strings.NewReader(string(b)), headers, nil)
}

func performRequest(r http.Handler, method string, path string, body io.Reader, headers map[string]string, queryStrings map[string]string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	if len(queryStrings) > 0 {
		q := req.URL.Query()
		for k, v := range queryStrings {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w, nil
}

func printOnFailed(t *testing.T) func(body string) {
	return func(body string) {
		if t.Failed() {
			t.Logf("Response: %#v", body)
		}
	}
}

func fakeJwtToken(t *testing.T, data *model.User) (string, model.User) {

	if data == nil {
		data = fakeUser(t)
	}

	jwtClaims := struct {
		jwt.StandardClaims
		ID       string   `json:"user_id" binding:"required"`
		Email    string   `json:"email" binding:"required"`
		Scope    []string `json:"scope" binding:"required"`
		Tenant   string   `json:"tenant" binding:"required"`
		Role     string   `json:"role"`
		IsClient bool     `json:"is_client"`
	}{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
		ID:       data.ID,
		Email:    data.Email,
		Scope:    data.Scope,
		Role:     data.Role,
		Tenant:   data.Tenant,
		IsClient: data.IsClient,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		t.Fatalf("Failed generating access token")
	}
	return accessToken, *data
}

func fakeUser(t *testing.T) *model.User {
	user := &model.User{
		ID:       fake.DigitsN(10),
		Email:    fake.EmailAddress(),
		Scope:    []string{"permission1", "permission2"},
		Role:     fake.Word(),
		Tenant:   fake.CharactersN(10),
		IsClient: false,
	}

	return user
}
