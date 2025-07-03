package dashboard

import "github.com/dracory/dashboard/shared"

func New() *dashboard {
	return &dashboard{
		theme:    "default",
		template: shared.TEMPLATE_DEFAULT,
	}
}
