package main

type credential struct {
	CREDENTIAL_PASS	string `json:credential_pass`
	USED_ID			string `json:used_id`
}

type dataClient struct {
	ID     string `json:"id"`
	SENSOR string `json:"sensor"`
	VALUE  string `json:"value"`
	TIME   string `json:"time"`
}

type client struct {
	ID                string `json:"id"`
	TIME_REGISTRATION string `json:"time_registration"`
	STATUS            string `json:"status"`
}

type allCredential []credential

type allDataClient []dataClient

type allClient []client

type registration struct {
	ID   string `json:"id"`
	TIME string `json:"time"`
}

type question struct {
	STATUS string `json:"status"`
}

type Sync struct {
	TIME  string `json:"time"`
	TOTAL int    `json:"total"`
}

type history struct {
	SENSOR string `json:"sensor"`
	VALUE  string `json:"value"`
	TIME   string `json:"time"`
}
