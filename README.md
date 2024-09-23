# Synapse SMTP Bridge

> Note: This is a pet project, just hacking for fun.

This is an SMTP server which bridges to Synapse (Matrix). 

Emails can be sent to it with `FROM: <SendingAccount>@matrix.bridge` and `TO: <RecipientNumberOrGroupID>@matrix.bridge`. If you don't control the sending email address you can set `SEND_FROM` to specify the matrix username that should be used when sending messages.

Plain text emails are sent normally via matrix. HTML Emails are rendered using Headless Chrome and `chromedp` to a PDF and that PDF is attached to the matrix message.

![image](https://user-images.githubusercontent.com/1939288/125082304-f6906b80-e0be-11eb-9050-35c00d30b091.png)

# Environment Variables for config

- `SEND_FROM` set to the matrix username to send from when `@matrix.bridge` email isn't used as the from address
- `SEND_TO` set to the matrix username/number to send to when `@matrix.bridge` isn't set in the recipient email address 
- `SMTP_USERNAME`
- `SMTP_PASSWORD`
- `SMTP_ALLOW_ANNON` set to `TRUE` to enable anonymous access
- `DEBUG` set to `TRUE` to see full data of incoming mail in stdout
- `SYNAPSE_SERVER_URL` set to the URL of the Synapse server
- `SYNAPSE_TOKEN` set to the token for the Synapse server

# Deploying on Kubernetes

1. Open `Deployment.yaml`
1. Edit the `/path/to/your/synapse/config/folder` value at the bottom of the file. This should point to the configuration directory of your `synapse`. I got this onto the server by configuring `synapse` locally then copying the data. 
> Note: This is a host mount so will work nicely on single node clusters, for multinode clusters a PVC or similar will be required.
1. `kubectl apply -f ./Deployment.yaml`

# Getting started dev

Fill out `.env-sample` and copy to `.env`

Start headless chrome
```
docker run -d --network=host --rm --name headless-shell chromedp/headless-shell --remote-debugging-address=0.0.0.0 --remote-debugging-port=9222 --disable-gpu --headless
```

Start the SMTP server
```
make run
```

Fire test emails
```
make test
```
