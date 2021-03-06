#### Nodes
```
(node) (p) (t) ()

(p:Person) (t:Thing)

MATCH (node:Label) RETURN node.property

MATCH (n:Class)-->(n2:Class)
WHERE n.propertyA = {value}
RETURN n2.propertyA, n2.propertyB
```

#### Relationships

* relationship-types like -[:KNOWS|:LIKE]->
* a variable name -[rel:KNOWS]-> before the colon
* additional properties -[{since:2010}]->
* structural information for paths of variable length -[:KNOWS*..4]->

```
MATCH (n1:Label1)-[rel:TYPE]->(n2:Label2)
WHERE rel.property > {value}
RETURN rel.property, type(rel)
```

#### Patterns

* friend-of-a-friend (user)-[:KNOWS]-(friend)-[:KNOWS]-(foaf)
* shortest path: path = shortestPath( (user)-[:KNOWS*..5]-(other) )
collaborative filtering (user)-[:PURCHASED]->* (product)<-[:PURCHASED]-()-[:PURCHASED]->(otherProduct)
* tree navigation (root)<-[:PARENT*]-(leaf:Category)-[:ITEM]->(data:Product)

#### Create

* `CREATE (you:Person {name:"You"}) RETURN you`
* create single node or more complex
  structures
  `MATCH (you:Person {name:"You"}) 
  CREATE (you)-[like:LIKE]->(neo:Database {name:"Neo4j"})
  RETURN you, like, neo`
* `MATCH (you:Person {name:"You"}) 
  FOREACH (name in ["Johan","Rajesh","Anna"] | CREATE (you)-[:FRIEND]-(:Person {name:name}))` FOREACH allows you to execute update operations for each element of a list
* Find Your Friends `MATCH (you:Person {name:"You"})-[:FRIEND]->(yourFriends) RETURN you, yourFriends`
* Create Second Degree Friends and Expertise `MATCH (neo:Database {name:"Neo4j"}) MATCH (anna:Person {name:"Anna"}) CREATE (anna)-[:FRIEND]->(:Person:Expert {name:"Amanda"})-[:WORKED_WITH]->(neo)` 
* Find Someone in your Network Who Can Help You Learn Neo4j `Match (you:Person {name:"You"}) MATCH (expert)-[:WORKED_WITH]->(db:Database {name:"Neo4j"}) MATCH path = shortestPath((you)-[:FRIEND*..5]-(expert)) RETURN db.expert, path`

#### BlingBling

* As relationships are stored efficiently, two nodes can share any number or type of relationships without sacrificing performance.
* Although relationships are directed, relationships can always be navigated regardless of direction.
* Since a relationship always has a start and end node, you can't delete a node without also deleting its associated relationships. To put it in another way, an existing relationship will never point to a non-existing endpoint.
  
#### User Defined Procedures and Functions

* Extending Cypher
* Functions  
  1. simple computations / conversions and return a single value
  2. can be used in any expression or predicate
* Procedures  
  1. are more complex operations and generate streams of results
  2. must be used with CALL clause and YIELD their result columns
* Functions & Procedures  
  * They are generate, fetch or compute data to make is avaliable to later processing steps in your Cypher query
* Listing & Using Functions & Procedures  
  * list available functions & procedures `CALL dbms.functions()`, `CALL dbms.procedures()`
  * each procedure returns one ore more columns of data. With `YIELD` these columns can be selected and also aliased and are then available in you Cypher statement.
  ```
  CALL dbms.procedures()
  YIELD name, signature, description 
  WITH split(name, ".") AS parts
  RETURN parts[0..-1] AS package, count(*) AS count, collect(parts[-1]) AS names
  ORDER BY count DESC
  ```
* Deploying Procedures & Functions
Build your own procedures or downloaded from an community project, they are packaged in a jar-file. -> `$NEO4J_HOME/plugins` Neo4j server and restart.
* Procedure and Function Gallery  
  * graph algorithms
  * index operations
  * database/api integration
  * graph refactorings
  * import and export
  * spatial index lookup
  * rdf import and export
  * and many more

#### Data Modeling
* A Node can have Relationships, Properties and Labels
* A Label is a named graph construct that is used to group nodes into sets -> :PERSON. All nodes labeled with the same label belongs to the same set. Many database queries can work with these sets instead of the whole graph, making queries easies to write and more efficient.
* A NODE may be labeled with any number of labels, including none, making labels an optional addition to the graph.
