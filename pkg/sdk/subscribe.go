package sdk

import "net/http"

type MBClient struct {
	host string
}

func (mbClient *MBClient) Subscribe(topic string, consumer string, handlerPath string) error {
	urlString := mbClient.host + "/register?topic=" + topic
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Consumer", consumer)
	req.Header.Set("X-Handler-Path", handlerPath)

	client := &http.Client{}

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
