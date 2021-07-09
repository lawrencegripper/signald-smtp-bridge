# Signald SMTP Bridge

> Note: This is a pet project, just hacking for fun.

This is an SMTP server which bridges to Signal Messenger. Emails can be send to it with `FROM: <SendingAccount>@signal.bridge` and `TO: <RecipientNumberOrGroupID>@signal.bridge`.

Plain text emails are sent normally via signal. HTML Emails are rendered using Headless Chrome and `chromedp` to a PDF and that PDF is attached to the signal message.

![image](https://user-images.githubusercontent.com/1939288/125082304-f6906b80-e0be-11eb-9050-35c00d30b091.png)


# Getting started

Fill out `.env-sample` and copy to `.env`

Start headless chrome
```
docker run -d --network=host --rm --name headless-shell chromedp/headless-shell --remote-debugging-address=0.0.0.0 --remote-debugging-port=9222 --disable-gpu --headless
```

Start signald - assuming signalctrl already used to link account
```
make signald
```

Start the SMTP server
```
make run
```

Fire test emails
```
make test
```
