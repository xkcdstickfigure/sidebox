# Sidebox
Create instant email inboxes for signing up to things

<img alt="screenshot of extension" src="https://sidebox.net/screenshot.png" width="192" height="228" />

### Why should I use this?
- Make multiple accounts on the same service
- Stop companies knowing your real email and tracking you across the internet with it
- Completely avoid spam - messages are in a separate email inbox that you can just delete

### Why is this better than services like 10minutemail?
You can create multiple email inboxes and they last forever (unless you delete them). This means you won't get locked out of accounts you make if you forget the email or password, or need an email 2fa code.

<!-- zero width space in the email below to avoid linking -->
### Why is this better than username+something​@gmail.com?
Companies can still easily figure out your normal email address, and it's not as easy to block them. With Sidebox, all they get is a random identifier, and once you delete it, it's useless.

### Repository structure
The main extension code is in `src/shared`. The `build.sh` script creates dev (for running the API locally) and prod (using the real API) builds of the extension for Firefox and Chromium. There's no React or TypeScript or anything fancy, it's just vanilla HTML/CSS/JS. The API is written in Go and is in [another repository](https://github.com/xkcdstickfigure/sidebox-api).
