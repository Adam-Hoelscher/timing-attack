# timing-attack

## What is this?
This is a proof-of-concept for a timing attack to recover a password from a (bad) authenticator. The basic idea was insipred by the wonderful mCoding Youtube video (https://www.youtube.com/watch?v=XThL0LP3RjY) on the same. James used a relatively simple scheme to decode a password in that video. This module recreates that methodology in Go and also expands. This will (hopefully) allow learners to better understand the aspects that attackers need to us to accomplish this and what tools the defender can use to protect their system.