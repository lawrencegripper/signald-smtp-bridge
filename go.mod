module github.com/lawrencegripper/signald-smtp-bridge

go 1.16

require (
	github.com/chromedp/cdproto v0.0.0-20210706234513-2bc298e8be7f
	github.com/chromedp/chromedp v0.7.3
	github.com/emersion/go-sasl v0.0.0-20220912192320-0145f2c60ead
	github.com/emersion/go-smtp v0.15.1-0.20221021114529-49b17434419d
	github.com/google/uuid v1.3.0
	github.com/marcospgmelo/parsemail v1.3.1-0.20201020162348-38663e9311e7
	gitlab.com/signald/signald-go v0.6.0
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)

replace github.com/marcospgmelo/parsemail => github.com/lawrencegripper/parsemail v1.3.1
