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
  