module github.com/lawrencegripper/signald-smtp-bridge

go 1.22

require (
	github.com/chromedp/cdproto v0.0.0-20210706234513-2bc298e8be7f
	github.com/chromedp/chromedp v0.7.3
	github.com/emersion/go-sasl v0.0.0-20220912192320-0145f2c60ead
	github.com/emersion/go-smtp v0.20.2
	github.com/google/uuid v1.3.0
	gitlab.com/signald/signald-go v0.6.1
)

require (
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.1.0-rc.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)

replace github.com/marcospgmelo/parsemail => github.com/lawrencegripper/parsemail v1.3.1
