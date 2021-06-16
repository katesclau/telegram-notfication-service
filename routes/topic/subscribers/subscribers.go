package subscribers

import (
	"net/http"

	"github.com/katesclau/telegramsvc/db"
	"github.com/katesclau/telegramsvc/utils"
)

var DB *db.DBClient

func GetMethods(db *db.DBClient) map[string]func(w http.ResponseWriter, r *http.Request) {
	DB = db
	methods := make(map[string]func(w http.ResponseWriter, r *http.Request))
	methods["GET"] = getHandler
	return methods
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	key := utils.KeyFromPath(r.URL.Path, 2)

	subscribers := DB.GetTopicSubscribers(key)
	utils.BuildResponse(w, r, subscribers, http.StatusOK)
}
