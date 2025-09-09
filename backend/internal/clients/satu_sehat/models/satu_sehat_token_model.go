package models

type SatuSehatTokenResponseSucess struct {
	RefreshTokenExpires string   `json:"refresh_token_expires_in"`
	ApiProductList      string   `json:"api_product_list"`
	ApiProductListJson  []string `json:"api_product_list_json"`
	OrganizationName    string   `json:"organization_name"`
	DeveloperEmail      string   `json:"developer.email"`
	TokenType           string   `json:"token_type"`
	IssuedAt            string   `json:"issued_at"`
	ClientId            string   `json:"client_id"`
	AccessToken         string   `json:"access_token"`
	ApplicationName     string   `json:"application_name"`
	Scope               string   `json:"scope"`
	ExpiresIn           string   `json:"expires_in"`
	RefreshCount        string   `json:"refresh_count"`
	Status              string   `json:"status"`
}
