package dashboard

func New() *dashboard {
	return &dashboard{
		theme:    "default",
		template: TEMPLATE_DEFAULT,
	}
}
