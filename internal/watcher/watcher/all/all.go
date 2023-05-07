package all

import (
	"demo/internal/watcher/watcher"
	"demo/internal/watcher/watcher/user"
)

func init() {
	watcher.Register("user", &user.UserWatcher{})
}
