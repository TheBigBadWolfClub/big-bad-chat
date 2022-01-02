package internal

import (
	"net/http"
)

const LOGO = `
ww     w      w  OOOOO  LL    FFFFF c---
 ww   w ww   w  OO OO  LL    FF  ----h--
  ww w   ww w  OO OO  LL    FFFFF ----a-
   ww     ww  OOOOO  LLLLL FF  --------t
`

type WolfChatHandler struct {
}

func (n WolfChatHandler) Path() string {
	return "/api/chat"
}

func (n WolfChatHandler) Handler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(LOGO))
}
