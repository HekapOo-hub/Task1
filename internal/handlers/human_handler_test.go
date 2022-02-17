package handlers

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHumanHandler_Create(t *testing.T) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"name": "human",
		"male": false,
		"age":  12,
	})
	require.NoError(t, err)
	request, err := http.NewRequest(http.MethodPost,
		"http://localhost:1323/human/create", bytes.NewBuffer(requestBody))
	require.NoError(t, err)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test human handler create error %v", err)
		}
	}()
	responseBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "human info was created", string(responseBody))
}

func TestHumanHandler_Update(t *testing.T) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"oldName": "human",
		"name":    "updated",
		"male":    false,
		"age":     122,
	})
	require.NoError(t, err)
	request, err := http.NewRequest(http.MethodPatch,
		"http://localhost:1323/human/update", bytes.NewBuffer(requestBody))
	require.NoError(t, err)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test human handler update error %v", err)
		}
	}()
	responseBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "human info was updated", string(responseBody))
}

func TestHumanHandler_Get(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet,
		"http://localhost:1323/human/get/updated", nil)
	require.NoError(t, err)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test human handler get error %v", err)
		}
	}()
	responseBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Contains(t, string(responseBody), "Age:122")
}

func TestHumanHandler_Delete(t *testing.T) {
	request, err := http.NewRequest(http.MethodDelete,
		"http://localhost:1323/human/delete/updated", nil)
	require.NoError(t, err)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := (&http.Client{}).Do(request)
	require.NoError(t, err)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Warnf("test human handler delete error %v", err)
		}
	}()
	responseBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "human's info was deleted", string(responseBody))
}
