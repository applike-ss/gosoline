package adwords

import (
	"github.com/applike-ss/gosoap"
	"github.com/applike/gosoline/pkg/mon"
)

type MediaService interface {
	Upload(mediaType *string, data *string) (*Response, error)
}

type mediaService struct {
	logger mon.Logger
	client *gosoap.Client
}

type Response struct {
	*gosoap.Response
}

type HeaderParams map[string]interface{}

func NewMediaService(logger mon.Logger, authToken string, params gosoap.HeaderParams) (MediaService, error) {
	cli, _ := gosoap.SoapClient("https://adwords.google.com/api/adwords/cm/v201809/MediaService?wsdl")

	return &mediaService{
		logger: logger,
		client: cli,
	}, nil
}

func (m *mediaService) Upload(mediaType *string, data *string) (*Response, error) {
	params := gosoap.Params{
		"media": map[string]interface{}{
			"xsi_type": "Image",
			// "mediaId":    int64(1),
			"type": "IMAGE",
			"data": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEUAAACnej3aAAAAAXRSTlMAQObYZgAAAApJREFUCNdjYAAAAAIAAeIhvDMAAAAASUVORK5CYII=",
			// "Image": map[string]interface{}{
			// },
		},
	}

	res, err := m.client.Call("upload", params)
	response := &Response{
		res,
	}
	return response, err
}
