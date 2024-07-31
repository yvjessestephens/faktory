// Generated by ego.
// DO NOT EDIT

//line dead.ego:1

package webui

import "fmt"
import "html"
import "io"
import "context"

import (
	"net/http"

	"github.com/yvjessestephens/faktory/client"
)

func ego_dead(w io.Writer, req *http.Request, key string, dead *client.Job) {

//line dead.ego:12
	_, _ = io.WriteString(w, "\n\n")
//line dead.ego:13
	ego_layout(w, req, func() {
//line dead.ego:14
		_, _ = io.WriteString(w, "\n\n")
//line dead.ego:15
		ego_job_info(w, req, dead)
//line dead.ego:16
		_, _ = io.WriteString(w, "\n\n<h3>")
//line dead.ego:17
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Error"))))
//line dead.ego:17
		_, _ = io.WriteString(w, "</h3>\n<div class=\"table-responsive\">\n  <table class=\"error table table-bordered table-striped table-light\">\n    <tbody>\n      <tr>\n        <th>")
//line dead.ego:22
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "ErrorClass"))))
//line dead.ego:22
		_, _ = io.WriteString(w, "</th>\n        <td>\n          <code>")
//line dead.ego:24
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(dead.Failure.ErrorType)))
//line dead.ego:24
		_, _ = io.WriteString(w, "</code>\n        </td>\n      </tr>\n      <tr>\n        <th>")
//line dead.ego:28
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "ErrorMessage"))))
//line dead.ego:28
		_, _ = io.WriteString(w, "</th>\n        <td>")
//line dead.ego:29
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(dead.Failure.ErrorMessage)))
//line dead.ego:29
		_, _ = io.WriteString(w, "</td>\n      </tr>\n      ")
//line dead.ego:31
		if dead.Failure.Backtrace != nil {
//line dead.ego:32
			_, _ = io.WriteString(w, "\n        <tr>\n          <th>")
//line dead.ego:33
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "ErrorBacktrace"))))
//line dead.ego:33
			_, _ = io.WriteString(w, "</th>\n          <td>\n            <code>\n            ")
//line dead.ego:36
			for _, line := range dead.Failure.Backtrace {
//line dead.ego:37
				_, _ = io.WriteString(w, "\n              ")
//line dead.ego:37
				_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(line)))
//line dead.ego:37
				_, _ = io.WriteString(w, "<br/>\n            ")
//line dead.ego:38
			}
//line dead.ego:39
			_, _ = io.WriteString(w, "\n            </code>\n          </td>\n        </tr>\n      ")
//line dead.ego:42
		}
//line dead.ego:43
		_, _ = io.WriteString(w, "\n    </tbody>\n  </table>\n</div>\n\n<form class=\"form-horizontal\" action=\"")
//line dead.ego:47
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(root(req))))
//line dead.ego:47
		_, _ = io.WriteString(w, "/morgue/")
//line dead.ego:47
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(key)))
//line dead.ego:47
		_, _ = io.WriteString(w, "\" method=\"post\">\n  ")
//line dead.ego:48
		_, _ = fmt.Fprint(w, csrfTag(req))
//line dead.ego:49
		_, _ = io.WriteString(w, "\n  <div class=\"pull-left\">\n    <a class=\"btn btn-default\" href=\"")
//line dead.ego:50
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(relative(req, "/morgue"))))
//line dead.ego:50
		_, _ = io.WriteString(w, "\">")
//line dead.ego:50
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "GoBack"))))
//line dead.ego:50
		_, _ = io.WriteString(w, "</a>\n    <button class=\"btn btn-primary btn-sm\" type=\"submit\" name=\"action\" value=\"retry\">")
//line dead.ego:51
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "RetryNow"))))
//line dead.ego:51
		_, _ = io.WriteString(w, "</button>\n    <button class=\"btn btn-danger btn-sm\" type=\"submit\" name=\"action\" value=\"delete\">")
//line dead.ego:52
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(t(req, "Delete"))))
//line dead.ego:52
		_, _ = io.WriteString(w, "</button>\n  </div>\n</form>\n")
//line dead.ego:55
	})
//line dead.ego:56
	_, _ = io.WriteString(w, "\n")
//line dead.ego:56
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
