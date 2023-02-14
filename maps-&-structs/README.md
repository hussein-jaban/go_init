## Summary

- ### Maps

  - Collections of value types that are accessed via keys
  - Created via literals or via make function
  - Members accessed via [key] syntax
    - myMap ["key"] = “value”
  - Check for presence with "value, ok" form of result
  - Multiple assignments refer to same underlying data

- ### Structs

  - Collections of disparate data types that describe a single concept
  - Keyed by named fields
  - Normally created as types, but anonymous structs are allowed
  - Structs are value types
  - No inheritance, but can use composition via embedding
  - Tags can be added to struct fields to describe field
