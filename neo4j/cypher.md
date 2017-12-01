#### Operators  
* General `DISTINCT, ., []`
* Mathematical `+, -, *, /, %, ^`
* Comparison `=, <>, <, >, <=, >=, IS NULL, IS NOT NULL`
* Boolean `AND, OR, XOR, NOT`
* String `+`
* List `+, IN, [x], [x .. y]`
* Regular Expression `=~`
* String matching `STARTS WITH, ENDS WITH, CONTAINS`

#### Update queries
If query only performs reads, Cypher will be lazy and not actually match the pattern until you ask for the results. In an updating query, the semantics are that all the reading will be done before any writing actually happens.  

`WITH` - when you want the aggregation to happen, and that the aggregation has to be finished before Cypher can start filtering.
```
MATCH (n)-[:Friend]->(friend)
WITH n, count(friend) AS HasFriends
WHERE HasFriends >= 1
RETURN n.name, HasFriends
ORDER BY HasFriends DESC; 
```
or writing the aggregated data to the graph
```
MATCH (n:Person {name:'John'})-[:Friend]->(friend)
WITH n, count(friend) AS HasFriends
SET n.HasFriends = HasFriends
RETURN n.HasFriends
```
You can chain together as many query parts as the available memory permits.

#### Returning data
`RETURN` clause has three sub-clauses that come with it `SKIP` / `LIMIT` and `ORDER BY`.

#### Uniqueness
Single pattern with two paths or two distinct patterns

#### Expressions in general
* A hexadecimal integer literal (starting with `0x`): `0xFC3A9`
* A boolean literal: `true`, `false`, `TRUE`, `FALSE`
* A property: `n.prop`, NODE.\`weird property name\`
* A dynamic property: `n["prop"]`, `rel[n.city + n.zip]`, `map[coll[0]]`
* A parameter `$param`, `$0`
* A list of expressions: `['a', 2, n.prop, $param]`
* A regular expression: `a.name =~ 'Tob.*'`
* A case-sensitive string matching expression: `a.surname START WITH 'Sven'`, `a.surname ENDS WITH 'son'`, `a.surname CONTAINS 'son'`
* A `CASE` expression

#### Parameters

#### Operators
* General operators `DISTINCT`, `.`, `[]`
* Mathermatical operators `+ - * / % ^`
* Comparison operators `= <> < > <= >= IS NULL IS NOT NULL`
* String-specific comparison operators `STARTS WITH ENDS WITH CONTAINS`
* Boolean operators `AND OR XOR NOT`
+ String operators `+` for concatenation, `=~` for regex matching

#### NULL
Testing any value against `null` with both the `=` and the `<>` operators always is `null`. This includes `null = null` and `null <> null`. The only way to reliably test if a value `v` is `null` is by using the special `v IS NULL`, or `v IS NOT NULL` equality operators.

#### Patterns
* `(a:User)-->(b)`
* `(a:User:Admin)-->(b)`
* `(a)-[{blocked: false}]->(b)`
* `(a)-[:REL_TYPE]->(b)`
* `(a)-[*2]->(b)` equals to `(a)-->()-->(b)`