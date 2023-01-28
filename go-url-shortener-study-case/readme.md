# URL Shortener Service

## Introduction

Sometimes when we browsing through social media, we want to share a link to other sites. However, some social media put restriction on character count for each post. Matters come to worse if the link we want to share doesn't care about SEO / URL readability. Which is why there are services to shorten URLs.

Basically, URL shortener works by:
1. We put a long url, then we get a short URL
2. If we put the short URL in our browser, we will be redirected to the long url

Sometimes these URL shortener not just redirecting users, but also collected data such as number of links visited.

## Constraints

1. Assume your team are mostly comfortable code using Redis, Go and PostgreSQL.
2. All of the URL are not evenly accessed, some of them are more frequently accessed than others.
3. Approximately, there will be 100 million new URL created every months.
4. The system needs to store the visited URL count.
5. The ratio of read:write approximately will be 100:1
6. The URL lifetime is only 1 year after creation. After that, it will be deleted.

## Exercise

1. What are the most suitable length and character types for the random URL based on the constraint?
    - consider that the URL needs to be URL safe
    - consider that the URL needs to be easily remembered, so you need to have the shortest length possible

2. In the sample code, the redirects use 301 (Permanently moved to a new location). Is this the correct HTTP response code? why?

3. If you use caching, which evictions policy may be suitable for the problem? how would you test it?

4. Explore the starter code, implement your design into code. Here are some things you can try:
    - Does the database table needs sharding (even if you use other than PostgreSQL)?
    - To store the visited URL count, we use asynchronous job. One job for every single URL visited. Is this the best approach given the possible traffic?

5. In this starter code, we have to retrieve the long URL and count the visit, please order from the most optimal approach, why?
    1. using job for view count and cache to read data
    2. using cache to read data only
    3. using job for view count only
    4. don't use job or cache

## Starter Code

In this exercise, you have been provided some starter code along with some docker-compose configuration. This can be a starting point if you want to explore more about this problem and want to try code it by yourself!

In this code 

There are 3 endpoints:
1. POST /items

```
Via this endpoint, user may put a new long URL

If user provides the *short_url* in body parameter, it will be used as the short_url. If the parameter left empty, then the system will generate a random short URL.

```

2. GET /items/{slug}
```
Getting the details of url such as:
- Title
- Content
- Slug
- Long URL
```

3. GET /items/job/{slug}
```
Getting the details of url such as, dispatch view count with Job:
- Title
- Content
- Slug
- Long URL
```

4. GET /items/cached/{slug}
```
Getting the details of url such as, get item data via cache:
- Title
- Content
- Slug
- Long URL
```

4. GET /items/cached-job/{slug}
```
Getting the details of url such as, get item data via cache and dispatch view count with job:
- Title
- Content
- Slug
- Long URL
```


5. GET /redirect/{slug}
```
Redirect the users if the slug is found, otherwise returns 404
```


## Notes
- This exercise intends to introduce you to system design, in reality it will be much more complex and harder to test and measure
- The starter code serves as starting point, however you may create your own starter code
- Stay away from libraries / frameworks which still don't have v1.y.z, since there will be breaking change and might cause you a headache in the future. Asyncq library used in this exercise still v0.y.z, but we still haven't found a better library to implement a simple task queue.