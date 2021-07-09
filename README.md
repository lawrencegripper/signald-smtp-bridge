# Signald SMTP Bridge

> Note: This is a pet project, just hacking for fun.

This is an SMTP server which bridges to Signal Messenger. Emails can be send to it with `FROM: <SendingAccount>@signal.bridge` and `TO: <RecipientNumberOrGroupID>@signal.bridge`.

Plain text emails are sent normally via signal. HTML Emails are rendered using Headless Chrome and `chromedp` to a PDF and that PDF is attached to the signal message.

If you don't know what `signald` is yet best to stop and learn how that works first here: https://github.com/thefinn93/signald

# Deploying on Kubernetes

1. Open `Deployment.yaml`
1. Edit the `/path/to/your/signald/config/folder` value at the bottom of the file. This should point to the configuration directory of your `signald`. I got this onto the server by configuring `signald` locally then copying the data. 
> Note: This is a host mount so will work nicely on single node clusters, for multinode clusters a PVC or similar will be required.
1. `kubectl apply -f ./Deployment.yaml`

# Getting started dev

Fill out `.env-sample` and copy to `.env`

Start headless chrome
```
docker run -d --network=host --rm --name headless-shell chromedp/headless-shell --remote-debugging-address=0.0.0.0 --remote-debugging-port=9222 --disable-gpu --headless
```

Start signald - assuming `signalctrl` already used to link account
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