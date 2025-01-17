# fetch
Adam Payne's submission for the fetch rewards coding challenge.

## Process
I noticed that you included the api.yml OpenAPI description file, so I assumed this was a hint that I should be using a code generator tool to stub out the server
interface rather than code it all by hand. I found the oapi-codegen package and used it in strict mode to generate the interface boilerplate code.
Then I began filling in the stubs. I noticed that the instructions never call for retrieving the receipt data, only the calculated points.
This made it easier, because I could map from a generated UUID string to an integer for the points, and not have to store the receipt as a data record.

I tried to optimize the point calculation by using simple builtin operators where possible, avoiding regexp. I also tried to keep the error handling simple by
just not awarding points if, for example, parsing one of the fields to a float, date, or time failed.

Overall, this was my first non-toy application in Go. Years ago, I played around with the select/channel processing, but did not develop anything too complicated. 
This was a fun challenge, and I thank you all for giving me the opportunity!
