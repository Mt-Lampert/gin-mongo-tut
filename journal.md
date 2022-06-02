# Important project data

- Name of the Database we're using: __GinMongoTut__
  - collections: __podcasts__

## 2022-06-02

### 6:34

What's next? Adding a complete podcast  to the database. This time without episodes,
so the `[]episodes` array will be empty. Here's the plan

1. Write a JSON file with one entry that provides a valid payload for the database
0. Provide a route for adding a new podcast, without sending it to the database.
   Instead, we will return the in-going payload together with a message.
0. After this has successful, we will take care about inserting the new podcast 
   into the database

### 10:14

It worked out! Had some hassle because of datatypes not congruent with JSON types which
are always strings, but as the Bard said, "All's well that ends well!", and I got it
all working.



## 2022-06-01

### 6:50

I have no idea -- yet -- how to update episodes in the episode array of the podcast.
Can't believe that the only way is extracting them to a collection of their own and 
cross-reference them to podcasts. And this bothers me a lot. Feeling incompetent about
something important.

### 8:18

After reading MongoDB documentation, I no longer feel bothered. I found out that the
makers of MongoDB have thought about that and offer solutions.

##



## 2022-05-31

### 06:52

Yesterday went a lot better than I thought. Creating the global `mgH` MongoDB handler for setting up a method set was a very good idea. And it worked out extremely well.

### 08:35

Did the first MongoDB query using custom Golang data structures. It went shockingly well.
Both quering the database AND converting the results into JSON for sending. Easy and tidy
as pie.
