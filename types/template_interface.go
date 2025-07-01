package types

type TemplateInterface interface {
	ToHTML(dashboard DashboardInterface) string
}
