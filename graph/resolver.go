package graph
import "github.com/go-pg/pg"

// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *pg.DB
}
