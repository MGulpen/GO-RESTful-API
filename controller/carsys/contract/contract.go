package contract

//CurrentVersion is a struct for the json obj from http://nl.carsys.online/version.json
type CurrentVersion struct {
	BuiltFromBranch string `json:"Built from branch"`
	CommitID        string `json:"Commit id"`
	BuildDate       string `json:"Build date"`
	BuildNumber     string `json:"Build number"`
}
