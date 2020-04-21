package main

type dataClient struct {
	ID 		string `json:"id"`
    SENSOR	string `json:"sensor"`
	VALUE	string `json:"value"`
	TIME	string `json:"time"`
}

type client struct {
	ID 					string `json:"id"`
	TIME_REGISTRATION	string `json:"time_registration"`
	STATUS				string `json:"status"`
}

type allDataClient []dataClient

type allClient []client

type registration struct {
	ID	string `json:"id"`
	TIME string `json:"time"`
}

type question struct {
	STATUS	string `json:"status"`
}

type sync struct {
	TIME string `json:"time"`
	TOTAL int `json:"total"`
}

type history struct {
	SENSOR	string `json:"sensor"`
	VALUE	string `json:"value"`
	TIME	string `json:"time"`
}
