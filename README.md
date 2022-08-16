# An end-to-end test example

## What is this?

An example monorepo, containing a very basic frontend application and a minimal framework for a Content Management System that drives the content, complete with an end-to-end test runner and a single user story.

![image](https://user-images.githubusercontent.com/50168/184842493-d53e9cd2-9c75-4ca2-a3b9-ef3be5cb3922.png)

The CMS server stars with a mock, hard-coded database and serves its data on port `5001` as JSON. The frontend application serves HTTP requests on port `5000`, and on request reads from the CMS, passes some of the data through some HTML templates, and returns the result.

This is designed to simulate the next version of the Daily Wire web stack: a more-or-less static frontend and a dynamic backend.

The focus of this demo is the top of the landing page of dailywire.com, and specifically only the Top and Featired stories. When running, the application looks a bit like this:

![image](https://user-images.githubusercontent.com/50168/184844042-d4ec46e4-5d11-40f2-9355-670b0569bfe8.png)

### How do I run it locally?

- Start the servers with `make run-servers`. You can then visit `http://localhost:5000` and, hopefully, see the slightly broken-looking content pictured above.

## Context
