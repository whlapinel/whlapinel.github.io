+++
date = '2026-05-14T07:01:26-04:00'
draft = false
title = 'Time to Focus'
+++

## Updates

I've been working on learning Actual Budget and I ultimately decided to finalize the switch and canceled my YNAB account. I must say, Actual is quite a nice piece of work. It has a simple, intuitive, and snappy interface, and I think its capabilities actually exceed YNAB's in some places. Don't quote me on this, but for example I believe you can only have one spending target per budget category in YNAB, but in Actual, the equivalent concept of templates allows as many as you want. Quite powerful. Templates are an experimental feature but I've found them to be fully functional, bug-free and very useful. The connection between Schedules and templates is quite handy as well.

I have been working on creating a "personal assistant" using Claude Code that has access to my Actual instance via an MCP server (I had my assistant make one, but there are lots of other versions out there), as well as my Todoist account. It also has access to my email and calendar. Some of the helpful things I've had my assistant do just in the last few days:

- View my transactions going back 1 year (downloaded as CSV from bank) and identify all recurring transactions in order to create templates and schedules in Actual. Very helpful!
- Reconcile the transaction registers in Actual for all accounts. Also very helpful!
  - this was a bit harder for Sonnet 4.6 than I expected, but it was still much faster than I would have done on my own. I had all kinds of issues with duplicates and my credit card (inherently more difficult than the checking account) hadn't been reconciled in months.
- Create watch events in my calendar extrapolating the pattern and projecting into the future. Again, quite helpful!
  - For my Navy drudgery I have watch periodically and it follows a version of the Panama schedule, and it would have been a fair amount of work to create the calendar events myself. Since the times and dates don't follow a pattern amenable to creating a simple event series, I'd have to create each event manually. Very quickly, my "assistant" completed the task with just a couple sentences of guidance: something like "please extrapolate the pattern from my watch schedule in my calendar and project into the future, ending September 30th. We're following a version of the Panama schedule, where we have 2 weeks on one shift and then rotate back one shift."

Another thing my assistant helped with is completing my very first OSS pull request for the Actual Budget API! Just needed to add a couple API endpoints for manipulating the category notes (the current interface for templates involves writing strings according to a certain syntax in category notes). On that note, contributing to Actual was really easy and fun, so they get +100 points for that as well!

It's always nice when a learning project turns out to be pragmatic and immediately useful!  Doesn't often work out that way.

I haven't found much reason yet for my assistant to use the Todoist API, but I have the MCP server set up and tested.

## Time to Focus

Ok so today I decided that I have been kind of flailing around trying to decide what to learn and do in my limited spare time on a day-to-day basis, and it's begun to dawn on me as well that my personal projects, as fun and helpful as they've been so far, might not do the best to prepare me for returning to my job at Amazon, which is obviously my biggest learning priority.

So I asked Claude Code to help me out by giving me a learning plan. I've posted it here in [Learning Plan]({{<relref "projects/learning-plan">}}). I work on secure communications and would like to understand the fundamentals at a more intuitive level. So I provided that general intent and Claude helped me create a general plan, and we refined it down to the day-by-day level.

My teacher-brain woke up once I started seeing slides and lesson plans and I had to hold myself back from reviving my course-manager project. 😅

The hard part will be following this plan, because it is quite ambitious, budgeting 2-3 hours a day for 7 days a week, no off-days. My wife and kids are here visiting me so between work and home it will be tough (as always) to find the time. We'll see how it goes!  I'll try and post here once a week about my progress.
