package route

import (
	"github.com/PaoloProdossimoLopes/go-library/middleware"
	"github.com/PaoloProdossimoLopes/go-library/wellcome"
)

func Init() {
	middleware.Get("/", wellcome.GetWellcomeHandler)
}
