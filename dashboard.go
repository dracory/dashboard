package dashboard

import "github.com/dracory/dashboard/types"

type dashboard struct {
	content string
}

var _ types.DashboardInterface = (*dashboard)(nil)

func (d *dashboard) ToHTML() string {
	return d.content
}

func (d *dashboard) GetContent() string {
	return d.content
}

func (d *dashboard) SetContent(content string) {
	d.content = content
}
