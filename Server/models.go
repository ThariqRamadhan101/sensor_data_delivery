package main

<<<<<<< HEAD
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

=======
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

>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
type allDataClient []dataClient

type allClient []client

type registration struct {
<<<<<<< HEAD
	ID   string `json:"id"`
=======
	ID	string `json:"id"`
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
	TIME string `json:"time"`
}

type question struct {
<<<<<<< HEAD
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
=======
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
>>>>>>> 1faa28365bc6d90da37a612597569c569b53f5a3
}
