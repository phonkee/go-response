# go-response

This package is small helper to write json messages to response writer.

# Redesign

currently api is under re-design and will be simplified and released
under `/v2` version
Let's try some examples of new api

```go
// Blank response
New(http.StatusOK).Write(r, w)

// Write json result
Result(map[string]string{}).Write(r, w)

```

# author

Peter Vrba <phonkee@pm.me>
