package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	g "github.com/AllenDang/giu"
	h "github.com/jPomeranz/mailman/internal/http"
)

const windowWidth = 1000
const windowHeight = 400

var (
	targetURL       string        = "https://"
	possibleMethods []string      = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodHead, http.MethodOptions, http.MethodTrace, http.MethodConnect}
	timeout         time.Duration = 10 * time.Second
	methodMenuIndex int32
	method          string = http.MethodGet
	statusCode      string
	headerKeys      []string   = make([]string, 0)
	headerValues    [][]string = make([][]string, 0)
)

func execute() {
	res, err := h.JSONRequest(targetURL, method, nil, timeout)
	displayResult(res, err)
}

func displayResult(res *http.Response, err error) {
	if err != nil {
		// TODO: Handle errors
		urlErr := err.(*url.Error)

		if urlErr.Timeout() {
			fmt.Println("Error: timeout")
		}

		log.Fatal("Error:", err)
	} else {
		statusCode = res.Status
		for k, v := range res.Header {
			headerKeys = append(headerKeys, k)
			headerValues = append(headerValues, v)
		}
	}

}

func methodChanged() {
	method = possibleMethods[methodMenuIndex]
}

func loop() {
	var headerRows []*g.RowWidget
	for i := range headerKeys {
		headerRows = append(headerRows, g.Row(g.LabelWrapped(headerKeys[i]), g.LabelWrapped(strings.Join(headerValues[i][:], ","))))
	}

	g.SingleWindow("Main", g.Layout{
		g.SplitLayout("Split", g.DirectionHorizontal, true, 500,
			g.Layout{
				g.Label("Request"),
				g.Line(
					g.Label("URL"),
					g.InputText("##URLInput", 400, &targetURL),
					g.Button("Execute", execute),
				),
				g.Line(
					g.Label("Method"),
					g.Combo("##MethodCombo", possibleMethods[methodMenuIndex], possibleMethods, &methodMenuIndex, 100, 0, methodChanged),
				),
			},
			g.Layout{
				g.Label("Response"),
				g.TabBar("Response", g.Layout{
					g.TabItem("Info", g.Layout{
						g.Table("InfoTable", true, g.Rows{
							g.Row(g.Label("Status Code"), g.LabelWrapped(statusCode)),
						}),
					}),
					g.TabItem("Headers", g.Layout{
						g.Table("HeaderTable", true, headerRows),
					}),
				}),
			},
		),
	})
}

func main() {
	wnd := g.NewMasterWindow("", windowWidth, windowHeight, 0, nil)
	wnd.Main(loop)
}
