package helpers

var (
	environment string
	redirectURL string
)

func Start(env string, redirectHost string) {
	environment = env
	redirectURL = redirectHost
}
