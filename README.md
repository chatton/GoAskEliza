My name is Cian Hatton, this repository holds an Eliza-Like chatbot.

The chatbot was my project for my Data Representation and Querying module in my 3rd year Software Development course at GMIT.

The course page can be found [here](https://data-representation.github.io/)

# Setup Instructions

Before you can run Eliza, you'll need to have [Go](https://golang.org/dl/) installed.
When you do, you can clone this repository by opening up a terminal or command line and using the folling command.

```bash
git clone https://github.com/chatton/GoAskEliza.git
```

From there, navigate into the src folder

```bash
cd GoAskEliza/src
```

You can start the web server by first compiling the code

```bash
go build .\main.go
```

and running the .exe

```bash
.\main.exe
```

or by simply running the command

```bash
go run .\main.go
```

The web server should now be up and running on port 8080.

You can start talking to Eliza by opening up a browser and navigating to

```
http://localhost:8080/
```

You can then start talking by entering text into the form at the bottom of the page and pressing Enter!

# What is Eliza?

The original [Eliza](https://en.wikipedia.org/wiki/ELIZA) program was a natural language processing computer program created in the 1960s.

Eliza simulates conversation using pattern matching with regular expressions, and some simple string manipulation.


# Problems that Arise

In the book [Paradigms of Artificial Intelligence Programming: Case Studies in Common Lisp](https://www.amazon.com/Paradigms-Artificial-Intelligence-Programming-Studies/dp/1558601910) there is a full chapter dedicated to Eliza.

In this chapter, the author talks about how Eliza seems to be "understanding" the human speaking with her at first glance, and that as long as she is spoken to in proper English, her responses will seem to demonstrate understanding. However when you proivde unusual input or oddly structured sentences, it becomes more obvious that there really is no "understanding" of what's happening in the conversation.

In my implementation, I tried to add some additional features with the plan on immitating intelligence and understanding, it's nothing more than string manipulation and pattern matching.

Some examples.

If the first sentence the user enters isn't a "greeting", Eliza will recognize this fact.

If the user says "you", in certain situations Eliza will say something like "Did you come here to talk about me?", giving the illusion of understanding in the conversation.

Eliza will never pick the same response twice, unless all possible responses have been picked already. I chose to implement this as duplicate responses stand out, i.e. break the illusion of intelligence. Non-duplicate responses don't necessarily stand out as showing more understanding or as more impressive, but it keeps the "flow" of the conversation more natural.

Eliza will occasionally "remember" a past question, and bring it up in one her generic catch all answers, again, this is just to make it seem like there is some understanding of the conversation.

In Eliza's catch all answers, she will try to change the subject to a topic that there are more specific patterns and answers for, again this soley intended to steer the direction of conversation into a place where she will have more specific responses, giving the illution of intelligence and understanding.

# Design Decisions

The Eliza struct consists of 2 interfaces, an AnswerGenerator and an AnswerPicker. I chose to user interfaces to allow the possiblity of altering her behaviour (even at runtime).

The AnswerGenerator does all the real work, it processes the string, matches patterns and gives back possible answers.
The AnswerPicker simply picks back an answer. I included 2 implementations, one that simply picks one at random, and another that will never pick duplicates unless there is no other option.

Other AnswerGenerator implementatiosn could be crearted, and nothing else would need to change.

An AnswerGenerator that will always generate the same responses for the same inputs being separated out from the 
AnswerPicker also makes it easier to write tests. It would be difficult to test the full functionality if the answer
was generated and chosen in the same package/struct.

I included multiple files that have different types of responses for different situations, this allows
you to change and add patterns and responses by simply altering the .dat files instead of having to go in and change any code.
The program will need to be restarted for any of these changes to take effect however.

All of the Eliza logic was intentionally kept completely separate from the web server aspect of the project. There's no need for the Eliza bot to know anything about the web server. This also allowed me to write the simple ***ask.go*** program, which is a small program that allows you to pass in a single question and get a single response, which I used to test Eliza functionality during development. If I needed to start a web server any time I wanted to test the logic, it would have sloved things down considerably and made the development process more cumbersome.

As there is no in-built Set data structure in Go, and multiple packages needed a data structure to quickly access and check for presence, I created a small StringSet struct. I could have just used a Map directly, but I decided that it would get enough use that it was worth implementing my own. 

I chose to use a slice of Response structs in order to represent the pattern/answer(s) pair. I did this because slices maintain order, this order allowed me to "prioritse" the more specific patterns by simply placing them at the top of the file. The more specific the pattern, the more "understanding" Eliza appears to have.

# Technical Problems That I Encountered.

I initially kept the Reponse structs in a map, this worked for a little bit, the un-ordered nature worked quite well to make things seem different each time, but as I added in more "generic" patterns, these came up more often instead of the specific ones, so I changed to an ordered datastructure instead.

At first I used a form in html and it made the request to /ask, this of course made new request to an endpoint that wasn't serving up any html. I instead used jQuery and ajax to remove the default behaviour of the form and send a request with JavaScript instead.

# Endpoints

This application currently has 3 endpoints,

1. ***/*** index.html is served when a request is made to the root resource. This serves up the actual Eliza web application that you can interact with in your browser.
2. ***/ask*** the /ask end point needs a "question" url parameter, it takes this value, passes it in to an Eliza instance, and writes the answer received to the response.
3. ***/history*** the /history returns, in JSON format, a list of all the past questions and answers of the current Eliza instance.

# Current limitations

Currently, when the web server is running, there is only a single instance of Eliza. This means that if multiple users were all connected, they would all be talking to the same Eliza, which would be "remembering" the answers of other users, she would effectively be having a conversation with all 3 users simulatiously. One way to solve this would be to maintain multiple Eliza instances and use cookies to keep track of which Eliza should be asked the question.

# Technologies used

The web server and Eliza were built using the [Go](https://golang.org/) programming language.
The web server serves a html page (using [bootstrap](http://getbootstrap.com/)), and ajax queries are made using JavaScript and [jQuery](https://jquery.com/).

# Misc.

See this [Discord Bot](https://github.com/chatton/ElizaBot) that uses this Eliza web server.

# References

This implementation of Eliza gave me many ideas in terms of the reflection map and also some of the pattern/response combinations that I implementated here. https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/

Chapter 15 of [Paradigms of Artificial Intelligence Programming: Case Studies in Common Lisp](https://www.amazon.com/Paradigms-Artificial-Intelligence-Programming-Studies/dp/1558601910) gave a good overview of the problems that arised in the original version.

This Eliza implementation http://www.manifestation.com/neurotoys/eliza.php3 gave me some ideas for enhancements in how my version could seem more "real" and "intelligent"

I consulted both the [jQuery](https://api.jquery.com/) and [bootstrap](https://getbootstrap.com/docs/3.3/getting-started/) documentation frequenty during development.
