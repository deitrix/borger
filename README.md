# Borger

Generic database access for Go.

## Getting Started

### Define your type

```go
type Car struct {
    ID          string
    Make        string
    Model       string
    Description string
}
```

### Create a `Table` for your type

```go
var Cars = borger.Table[Car]{
	Name: "cars",
	Columns: func(car *Car) map[string]any {
		return map[string]any{
			"id":          &car.ID,
			"make":        &car.Make,
			"model":       &car.Model,
			"description": &car.Description,
		}
	},
}
```

### Set the global database (I know. Ew)

```go
db, err := sql.Open("...", "...")
...
borger.DB = db
```

### Select

```go
cars, err := Cars.Select("id", "make", "model")
if err != nil {
    log.Fatalf("failed to select cars: %v", err)
}
```