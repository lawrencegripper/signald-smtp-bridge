module github.com/lawrencegripper/signald-smtp-bridge

go 1.16

require (
	github.com/armon/consul-api v0.0.0-20180202201655-eb2c6b5be1b6 // indirect
	github.com/chromedp/cdproto v0.0.0-20210706234513-2bc298e8be7f
	github.com/chromedp/chromedp v0.7.3
	github.com/emersion/go-sasl v0.0.0-20200509203442-7bfe0ed36a21
	github.com/emersion/go-smtp v0.15.0
	github.com/google/uuid v1.3.0
	github.com/marcospgmelo/parsemail v1.3.1-0.20201020162348-38663e9311e7
	github.com/ugorji/go v1.1.4 // indirect
	github.com/xordataexchange/crypt v0.0.3-0.20170626215501-b2862e3d0a77 // indirect
	gitlab.com/signald/signald-go v0.0.0-20211107182225-dd56e3d6f746
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
)

replace github.com/marcospgmelo/parsemail => github.com/lawrencegripper/parsemail v1.3.1
