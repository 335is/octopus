package octopus

import (
	"encoding/json"
)

// Feed describes a feed
type Feed struct {
	ID                                string   `json:"Id"`
	Name                              string   `json:"Name"`
	FeedType                          string   `json:"FeedType"`
	ApiVersion                        string   `json:"ApiVersion"`
	RegistryPath                      string   `json:"RegistryPath"`
	FeedUrl                           string   `json:"FeedUrl"`
	Username                          string   `json:"Username"`
	Password                          string   `json:"Password"`
	PackageAcquisitionLocationOptions []string `json:"PackageAcquisitionLocationOptions"`
	SpaceId                           string   `json:"SpaceId"`
	LastModifiedOn                    string   `json:"LastModifiedOn"`
	LastModifiedBy                    string   `json:"LastModifiedBy"`
}

// GetFeeds gets all Octopus feeds
func (c *Client) GetAllFeeds() ([]Feed, error) {
	feeds := []Feed{}

	resp, err := c.DoGetRequest("feeds/all")
	if err != nil {
		return feeds, err
	}

	err = json.NewDecoder(resp.Body).Decode(&feeds)
	if err != nil {
		return feeds, err
	}

	return feeds, nil
}
