# GoAssessment

<h1>Assesment 01 :</h1>
Create 3 micro-services
   1. Available tickets
   2. Adding tickets to cart
   3. Allocating tickets to user once paymen is completed.

<h2>Some points on solution:</h2>
<p>I have created another service called <b>TicketStore</b>, so that it acts similar to a database. This enabled me to implement appropriate locks when multiple people are accessing records either ro read or write. </p>
<p>We could replace <b>TicketStore</b> with a layer directly interacting with a database. In such a case we can user locking system provided by the database. We could also use primary key combination to prevent multiple records being added to PreBooking table.</p>


<h1>Assessment 02 : </h1>
Build a sample Go app where you have a list of users, each with properties like Name, email, mobile,
etc. Each user also has their preferred way of getting notification from app for. e.g. via email, SMS,
call.

<h2>Some points on solution:</h2>
<p>Implemented Notification using interface feature of golang.</p>


// Naveen GM
This is my new commit to that I can test the Orbit web-hook integrations.
