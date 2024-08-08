package wspush

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/AnnonaOrg/annona_bot/core/service"
	log "github.com/sirupsen/logrus"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}

	_, channelStr, ok := strings.Cut(r.URL.Path, "/ws/push_v2/")
	if len(channelStr) <= 0 || !ok {
		log.Debugf("收到非法推送: %s", r.URL.Path)
		return
	}
	log.Debugf("r.URL.Path: %s", r.URL.Path)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		log.Errorf("io.ReadAll(r.Body): %v", err)
		return
	}

	if err := service.PushMsgData(body); err != nil {
		fmt.Fprintf(w, "err: %v", err)
		log.Errorf("PushMsgData(%s): %v", string(body), err)
		return
	}
	fmt.Fprintf(w, "ok")
}
