package models

type Authorize struct {
	ID       int    `json:"id"`
	AuthCL   string `json:"authCL"`
	AuthPV   string `json:"authPV"`
	MainAuth string `json:"mainAuth"`
}
