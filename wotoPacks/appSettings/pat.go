package appSettings

import "github.com/Araragi10/SAO_k3/wotoPacks/interfaces"

func (_s *AppSettings) GetPatClient() interfaces.WClient {
	return _s.patClient
}

func (_s *AppSettings) SetPatClient(client interfaces.WClient) {
	if _s.patClient != client {
		_s.patClient = client
	}
}
