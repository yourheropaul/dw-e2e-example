# An end-to-end test example

A forward-looking example of how we might specify and automatically end-to-end test our software applications in the future.

## What is this?

An **example monorepo**, containing a very basic frontend application and a minimal framework for a Content Management System that drives the content, complete with an **end-to-end test runner** and a single user story.

![image](https://user-images.githubusercontent.com/50168/184842493-d53e9cd2-9c75-4ca2-a3b9-ef3be5cb3922.png)

The CMS server stars with a mock, hard-coded database and serves its data on port `5001` as JSON. The frontend application serves HTTP requests on port `5000`, and on request reads from the CMS, passes some of the data through some HTML templates, and returns the result.

This is designed to simulate the next version of the Daily Wire web stack: a more-or-less static frontend and a dynamic backend.

The focus of this demo is the top of the **landing page of dailywire.com**, and specifically **only the Top and Featured stories**. 

### How do I run it locally?

#### Servers

Start the servers with `make run-servers`. You can then visit http://localhost:5000 and, hopefully, see some slightly broken-looking content:

![image](https://user-images.githubusercontent.com/50168/184844042-d4ec46e4-5d11-40f2-9355-670b0569bfe8.png)

http://localhost:5001 should be serving a bunch of JSON.

#### End-to-end test runner

`make e2e` should start the test runner. If successful, the output will look like this:

![image](https://user-images.githubusercontent.com/50168/184847676-2a0fc2bd-8787-4450-a0cc-8451540b56f1.png)

## Context: why do this at all?

The Daily Wire, as a software project, is currently comprised of a a frontend server, one or two CMS applications (depending on how you count them), a set of backend services, some mobile and specialist applications, a spin-off (Jeremy's Razors) and a maze of support applications that handle billing, analytics, video, audio nd so forth.

**We want all this stuff to work as a whole**. It's not really practial for a human, not matter how accomplished at quality assurance, to run through every single flow of the applications and verify all behaviour on every code change. What we need is some degree of automated testing; in other words, we want our computers to take over as much of the software validation as they can. Repeated, specific task execution is, afterall, what they're better than humans at doing.

Due to its organic (read: uncontrolled) technical growth, The Daily Wire's software stack is currently distributed over several different version control solutions. This presents quite a challenge for automated testing: Since all sorts of data transit between the various system components in a variety of structures (and a mixture of well-formedness and, well, less-well-formedness) it's not feasible to simulate or mock out the responses and exchanges of every part of the system &dash; Postman just won't be enough in the long term.

Even if it were, each of the components is part of a harmonious, self-supporting system: the function of each component influences, informs or directs one or more of the other components. It doesn't follow that they will all work together just because they all work separately.

## Proposal: A better way

### Principle 1: Human-centricity

Since the entire Daily Wire ecosystem is intended for the use and benefit of humans, it makes sense to think about the whole thing from a human perspective. There is a value we want to deliver to our various users, and we should be able to easily articulate what it is. If we can't, then what are we doing?

People using software systems seldom think in terms of the technical properties required to get their work done. 

It would be a rare individual who, when wanting to see their bank balance online, would think, _What I need is some database servers backed onto a reliable messaging system, which can pass information though a secure API that a frontend web application, via TLS,  can query to tell me the number of dollars in my account_.  

Most people might say to themselves something more like, _I'd like to see my bank balance, it would be convenient if I could do it instantly on my phone and, because I want to keep it private, I don't mind identifying myself_. 

Because we're engineers, it's very easy to fall into thinking in purely technical terms. We must resist this, and flip it on its head. Let us instead adopt the following maxims:

> **if our product does everything it's supposed to do for all the people we care about, then it works &mdash; and if it doesn't do that, it doesn't work.**

and:

> **No amount of technical cleverness or correctness matters if there's no benefit to the users** (and bear in mind that we are ourselves users).

To that end, I've started work on a [BDD-inspired behavioural specification](https://docs.google.com/spreadsheets/d/1s5_zTN50Y3JJNZmsyn5cALOPnD1XEUiuMxXqLszb05o/edit#gid=270614781) which breaks down all currently known types of human users by type and motivation to use the Daily Wire system, catalogues all of their specific desires and will, eventually, contain a full set of explicit use cases that will inform and direct our technical efforts (including this [trivial example](https://docs.google.com/spreadsheets/d/1s5_zTN50Y3JJNZmsyn5cALOPnD1XEUiuMxXqLszb05o/edit#gid=1717537016), which is implemented in this repository). 

I further propose that we use this specification as our living, evolving, canonical document for what we're building and why. It's the recording of a continuing conversation about what we're doing and why we're doing it.

This currently lives in Google Docs, but it should eventually live in version control. For more details on why that is, please see my cure-for-insomnia [whitepaper on how specifications should be managed](https://endian.io/theory/adject-clarity-five-integral-c-words-of-maximally-useful-product-specification-for-teams/).

### Principle 2: Software design and testing is downstream of specification

If we think about what uses want and desire first, design and implemented (both of the direct user experience and the underlying technology) becomes much more straightforward. When we know what we want to achieve and what a successful version of that achievement looks like, testing becomes trivial: the software either meets the requirements or it doesn't. Red light or green light.

This approach affords us some other advantages too.

First, it makes the 10,000 foot view of product management pretty easy. What features can we tell users about in our marketing materials? All the things that have passing testing, in green. What's being worked on right now, from a user perspective? All the stuff with partially passing testing, in orange.

Second, non-technical people can easily contribute the specification. It's not necessary to understand the complexities of crypto math; it's only necessary to know what users want to see.

Third, we don't have to insist on tests that cover 100% of the code. In a legacy system like Garizon, there's a bunch of stuff we don't need immediately, and we don't care if it doesn't work at the moment. This makes test-writing much more cost-efficient.

And yes, visual UX flows and graphic designs are a form of specification. We need them too, but they should always be secondary to thinking about what we ultimately want to deliver and to whom.
