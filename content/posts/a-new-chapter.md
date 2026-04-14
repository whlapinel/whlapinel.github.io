+++
date = '2026-04-13T07:01:26-04:00'
draft = true
title = 'A New Chapter'
+++

I work at Amazon now. And my family and I moved to the Seattle area. But right now I'm in Florida and I'm on military leave for 6 months. Enough about that!

Before my move to Florida a couple weeks ago, I had been setting up a "home lab" - in my case, just several containerized servers running on an laptop.  So in my spare time I've been working on setting up a remote version of this. You can see the repo [here](https://github.com/whlapinel/home-server).

Much of my motivation for this project has been to continue my education regarding AWS. The project has served that purpose quite well!

I have also enjoyed learning about various popular open-source services, including:

- Actual Budget (budgeting, similar to YNAB)
- Vikunja (to-do lists)
- Vaultwarden (lightweight version of Bitwarden) (password management)
- Grav

But the most educational and really fun part is setting it all up behind a Wireguard VPN. In the EC2 security group I only expose Wireguard's port and protocol (UDP) to the public.  This way I don't have to worry much about the vulnerabilities of each service or the strength of passwords my family members might be using with those services. I connect my family to the VPN by giving them the QR code and it's all nice and easy.

It was a bit tricky making my Synology DSM 223 available to family. I'm on hotel Wifi. There's no easy way to install a Wireguard client on the DSM 223 (there are SPKs to add the module to the kernel out there, but they're outdated and didn't work). So I have it connected via ethernet to my laptop and sharing the internet connection, and port-forwarding. Only problem is when I disconnect my laptop they lose connection to the Synology. So I need a better long-term solution.

Vikunja is a really cool project and I was considering switching from Todoist. I did the migration and tried it out, but I have to say, Vikunja's UI is substantially more rough than Todoist's and after some consideration, I've decided the price of sticking with Todoist is worth it for now.

For Actual Budget, I have been trying it out and I'm really impressed. This is a real contender to YNAB, my current choice budgeting software. Most other mainstream budgeting apps (Quicken, Mint, Monarch) don't do envelope budgeting, but Actual Budget does. And its UI is quite smooth. At this point I think I'm going to make the switch, but I'll keep trying for a couple weeks before canceling my YNAB account.

That's all for now!
