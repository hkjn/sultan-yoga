sultan-yoga
===========
Source for Sultan Yoga.

Excruciatingly gentle instructions
===========
Bringing up a local development server from this directory:
$ dev_appserver.py . --port 1234 --host=0.0.0.0

The dev server will be accessible from all IPs, so make sure that the
port on the machine hosting the server is not accessible from the
public Internet.

If the hostname is 'zero-one', the web page is then found at:
http://zero-one:1234/

Any changes to the server code will be picked up automatically.

Press ^C (Ctrl+C) in the terminal running the server to stop it.

