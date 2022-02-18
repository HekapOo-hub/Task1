package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestUserHandler_Authenticate(t *testing.T) {
	resp, err := http.Get("http://localhost:1323/signIn?login=admin&password=1234")
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test user handler auth error %v", err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	_, err = fmt.Sscanf(string(body), "access token: %s\nrefresh token: %s", &accessToken, &refreshToken)
}

func TestUserHandler_Create(t *testing.T) {
	requestBody, err := json.Marshal(map[string]string{
		"login":       "noadmin",
		"password":    "1234",
		"re_password": "1234",
	})
	require.NoError(t, err)
	request, err := http.NewRequest(http.MethodPost,
		url+"user/create", bytes.NewBuffer(requestBody))
	require.NoError(t, err)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test user handler auth error %v", err)
		}
	}()
	responseBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "user was created", string(responseBody))
}

func TestUserHandler_Get(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, url+"user/get/noadmin", nil)
	require.NoError(t, err)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test user handler auth error %v", err)
		}
	}()
	responseBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	log.Warn(string(responseBody))
	require.Contains(t, string(responseBody), "User's info:")
}

func TestUserHandler_Update(t *testing.T) {
	requestBody, err := json.Marshal(map[string]string{
		"oldLogin": "noadmin",
		"newLogin": "updated",
		"password": "12345",
	})
	request, err := http.NewRequest(http.MethodPatch,
		url+"user/update", bytes.NewBuffer(requestBody))
	require.NoError(t, err)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-type", "application/json")
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test user handler update error %v", err)
		}
	}()
	responseBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "user was updated", string(responseBody))
}

func TestUserHandler_Delete(t *testing.T) {
	request, err := http.NewRequest(http.MethodDelete,
		url+"user/delete/updated", nil)
	require.NoError(t, err)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test user handler delete error %v", err)
		}
	}()
	responseBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "user was deleted", string(responseBody))
}

func TestUserHandler_Refresh(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet,
		url+"refresh/update", nil)
	require.NoError(t, err)
	request.Header.Set("Authorization", "Bearer "+refreshToken)
	request.Header.Set("Content-type", "application/json")
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test user handler refresh error %v", err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	_, err = fmt.Sscanf(string(body), "access token: %s\nrefresh token: %s", &accessToken, &refreshToken)
	require.NoError(t, err)
}

func TestUserHandler_LogOut(t *testing.T) {
	request, err := http.NewRequest(http.MethodDelete,
		url+"refresh/logOut", nil)
	require.NoError(t, err)
	request.Header.Set("Authorization", "Bearer "+refreshToken)
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test user handler log out error %v", err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "user logged out", string(body))
}
