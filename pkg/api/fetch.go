package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func fetch(url string, jsonResp interface{}) error {
	if c == nil {
		InitCache()
	}

	// try to get body from cache before actually fetching it
	cachedBody, cached := c.Get(url)
	if cached {
		err := json.Unmarshal(cachedBody.([]byte), jsonResp)
		if err != nil {
			// unmarshal returns error,
			// because the cached result is a plain text when
			// the resource is not found
			return errors.New("NOT_FOUND")
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			setCache(url, body) // also save not found result to cache
			return errors.New("NOT_FOUND")
		} else {
			return errors.New("UNKNOWN")
		}
	}

	err = json.Unmarshal(body, jsonResp)
	if err != nil {
		return err
	}

	setCache(url, body)
	return nil
}
