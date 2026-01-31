# Go OpenFGA Demo

A **minimal learning project** demonstrating how to use **OpenFGA** with Go via the official **go-openfga SDK**.  
This repository is intentionally small and exists only to **understand OpenFGA authorization concepts**, not production patterns.

---

## What This Demo Covers

This demo exercises the core OpenFGA APIs using a simple scenario:

- Write relationship tuples  
- Delete relationship tuples  
- Check permissions (`Check API`)  
- List accessible objects (`ListObjects API`)  

All examples use:
```
user:anne
document:Z
relation: reader
```

---

## Authorization Model Used

```json
{
  "schema_version": "1.1",
  "type_definitions": [
    { "type": "user" },
    {
      "type": "document",
      "relations": {
        "reader": { "this": {} },
        "writer": { "this": {} },
        "owner": { "this": {} }
      }
    }
  ]
}
```

Meaning:
- `user` → actor  
- `document` → resource  
- `reader`, `writer`, `owner` → permissions  

---

## Core Operations Explained

### Create Relationship Tuple

```
user:anne ── reader ──> document:Z
```

Grants read access to Anne for document Z.

---

### Check Permission (Check API)

```go
checkAPI(fgaClient, modelId)
```

Example result:
```json
{
  "allowed": true
}
```

---

### List Objects (ListObjects API)

```go
listObjectsAPI(fgaClient, modelId)
```

Example result:
```json
{
  "objects": ["document:Z"]
}
```

---

### Delete Relationship Tuple

```go
deleteRelationshipTuple(fgaClient, modelId)
```

Removes access between the user and the document.

---

## Why OpenFGA?

OpenFGA enables **relationship-based access control (ReBAC)** by modeling  
**who can do what on which resource**, keeping authorization logic outside application code.

Instead of:
```go
if user.Role == "admin" { ... }
```

You define permissions declaratively using relationships.

---

## Scope of This Repo

- ✔️ Focused on learning OpenFGA APIs  
- ✔️ Minimal, readable example  
- ❌ Not production-ready  
- ❌ No API layer, auth flow, or persistence  

---

## References

- OpenFGA Docs: https://openfga.dev  
- Go SDK: https://github.com/openfga/go-sdk  
