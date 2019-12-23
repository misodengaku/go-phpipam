package phpipam

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Subnet struct from phpipam
type Subnet struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		ID                    string      `json:"id"`
		Subnet                string      `json:"subnet"`
		Mask                  string      `json:"mask"`
		SectionID             string      `json:"sectionId"`
		Description           string      `json:"description"`
		FirewallAddressObject interface{} `json:"firewallAddressObject"`
		VrfID                 string      `json:"vrfId"`
		MasterSubnetID        string      `json:"masterSubnetId"`
		AllowRequests         string      `json:"allowRequests"`
		VlanID                string      `json:"vlanId"`
		ShowName              string      `json:"showName"`
		Device                string      `json:"device"`
		Permissions           string      `json:"permissions"`
		PingSubnet            string      `json:"pingSubnet"`
		DiscoverSubnet        string      `json:"discoverSubnet"`
		DNSrecursive          string      `json:"DNSrecursive"`
		DNSrecords            string      `json:"DNSrecords"`
		NameserverID          string      `json:"nameserverId"`
		ScanAgent             string      `json:"scanAgent"`
		IsFolder              string      `json:"isFolder"`
		IsFull                string      `json:"isFull"`
		Tag                   string      `json:"tag"`
		EditDate              interface{} `json:"editDate"`
		Links                 []struct {
			Rel     string   `json:"rel"`
			Href    string   `json:"href"`
			Methods []string `json:"methods"`
		} `json:"links"`
	} `json:"data"`
	Message string `json:"message"`
}

type SubnetAddresses struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    []struct {
		ID                    string      `json:"id"`
		SubnetID              string      `json:"subnetId"`
		IP                    string      `json:"ip"`
		IsGateway             interface{} `json:"is_gateway"`
		Description           interface{} `json:"description"`
		Hostname              string      `json:"hostname"`
		Mac                   interface{} `json:"mac"`
		Owner                 interface{} `json:"owner"`
		Tag                   string      `json:"tag"`
		DeviceID              string      `json:"deviceId"`
		Port                  interface{} `json:"port"`
		Note                  interface{} `json:"note"`
		LastSeen              interface{} `json:"lastSeen"`
		ExcludePing           string      `json:"excludePing"`
		PTRignore             string      `json:"PTRignore"`
		PTR                   string      `json:"PTR"`
		FirewallAddressObject interface{} `json:"firewallAddressObject"`
		EditDate              interface{} `json:"editDate"`
		Links                 []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
	} `json:"data"`
	Message string `json:"message"`
}

type SubnetUsage struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		Used             int     `json:"used"`
		Maxhosts         int     `json:"maxhosts"`
		Freehosts        int     `json:"freehosts"`
		FreehostsPercent float64 `json:"freehosts_percent"`
	} `json:"data"`
	Message string `json:"message"`
}

// GetSubnet Client pointer method to get all phpipam subnet data using subnetID
// string, returns Subnet struct and error
func (c *Client) GetSubnet(subnetID string) (Subnet, error) {
	var subnetData Subnet
	req, _ := http.NewRequest("GET", c.ServerURL+"/api/"+c.Application+"/subnets/"+subnetID+"/", nil)
	body, err := c.Do(req)
	if err != nil {
		return subnetData, err
	}
	err = json.Unmarshal(body, &subnetData)
	if err != nil {
		return subnetData, err
	}
	if subnetData.Code != 200 {
		return subnetData, errors.New(subnetData.Message)
	}
	return subnetData, nil
}

// GetSubnet Client pointer method to get all phpipam subnet data using subnetID
// string, returns Subnet struct and error
func (c *Client) GetSubnetAddresses(subnetID string) (SubnetAddresses, error) {
	var subnetData SubnetAddresses
	req, _ := http.NewRequest("GET", c.ServerURL+"/api/"+c.Application+"/subnets/"+subnetID+"/addresses/", nil)
	body, err := c.Do(req)
	if err != nil {
		return subnetData, err
	}
	err = json.Unmarshal(body, &subnetData)
	if err != nil {
		return subnetData, err
	}
	if subnetData.Code != 200 {
		return subnetData, errors.New(subnetData.Message)
	}
	return subnetData, nil
}

// GetSubnet Client pointer method to get all phpipam subnet data using subnetID
// string, returns Subnet struct and error
func (c *Client) GetSubnetUsage(subnetID string) (SubnetUsage, error) {
	var subnetData SubnetUsage
	req, _ := http.NewRequest("GET", c.ServerURL+"/api/"+c.Application+"/subnets/"+subnetID+"/usage/", nil)
	body, err := c.Do(req)
	if err != nil {
		return subnetData, err
	}
	err = json.Unmarshal(body, &subnetData)
	if err != nil {
		return subnetData, err
	}
	if subnetData.Code != 200 {
		return subnetData, errors.New(subnetData.Message)
	}
	return subnetData, nil
}
