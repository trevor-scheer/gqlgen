package models

import "github.com/trevor-scheer/gqlgen/integration/server/remote_api"

type Viewer struct {
	User *remote_api.User
}
