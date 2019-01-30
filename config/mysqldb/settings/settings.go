package settings

const (
	User             string = "root"
	Password         string = "1143"
	IP               string = "localhost"
	Port             string = "3306"
	Database         string = "/vehicle_db"
	ConnectionString string = User + ":" + Password + "@tcp(" + IP + ":" + Port + ")/" + Database
)
