package document

import (
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	ctx "tabs/internal/context"
	"tabs/internal/info"
	vi "tabs/internal/views"
)

// Document represents a document, used for wrapping content in.
type Document struct {
	Context *ctx.Context
	Views   *vi.Views
}

// NewDocument creates a new instance of Document.
func NewDocument(context *ctx.Context, views *vi.Views) *Document {
	return &Document{Context: context, Views: views}
}

// CreateDocument creates an HTML document.
// The document includes the necessary scripts and stylesheets for the page to function properly.
// The content area is a slot, which can be updated dynamically using HTMX.
func (d *Document) CreateDocument(info *info.Info, content ...Node) Node {
	return HTML5(HTML5Props{
		Title:    info.Title,
		Language: "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("/public/styles/generated/output.css"), Type("text/css"), Defer()),
			Script(Src("/public/scripts/document.js"), Defer()),
			Script(Src("/public/scripts/lib/htmx.min.js"), Defer()),
			Script(Src("/public/scripts/lib/alpinejs.min.js"), Defer()),
		},
		Body: []Node{
			Body(
				Div(Class("h-[80vh]"), d.Views.Navbar.CreateNavbar(info),
					Div(Class("h-full pt-5 pb-5 pl-20 pr-20"),
						Div(Class("h-full rounded-lg bg-base-200"),
							Main(Class("h-full p-5"), ID("content"),
								d.CreatePartial(content...)))))),
		},
	})
}

// CreatePartial creates a container for dynamic content.
// It is used when creating the document and for a snippet when swapping content with HTMX.
// The included script offers the minimally required functionality for one snippet.
func (d *Document) CreatePartial(content ...Node) Node {
	return Div(Group(content), Script(ID("partial-script"), hx.SwapOOB("true"), Raw(partialScript)))
}

// partialScript represents a JavaScript code snippet.
const partialScript string = `
// After loading the DOM
document.addEventListener("DOMContentLoaded", () => refresh(), {once: true});

/**
 * Refreshes the web page if there is no <main> element.
 * Duplicating a tab in the browser fetches the most recent HTML returned from the server.
 * However, with HTMX, these can be snippets (e.g., for swapping).
 * A snippet by itself (no linked files etc.) is not useful, hence, a refresh is done.
 * @return {void} Nothing is returned.
 */
function refresh(){
    if (!document.querySelector("main")) {
        location.reload();
    }
}`
