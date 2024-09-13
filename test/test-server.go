package test

const serverAliveMessage string = "Server is alive and kickin' !"

func ServerAlive() string {
	return serverAliveMessage
}

func ServerStats() string {
	return "Server alive since 9 hours"
}