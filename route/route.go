package route

import (
	"github.com/PaoloProdossimoLopes/go-library/server"
	"github.com/PaoloProdossimoLopes/go-library/wellcome"
)

func Init() {
	server.Get("/", wellcome.GetWellcomeHandler)
	server.Post("/", wellcome.GetWellcomeHandler)
}
