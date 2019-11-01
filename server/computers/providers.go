package computers

import "github.com/google/wire"

// Set of providers for computers components.
var Providers = wire.NewSet(NewStore, HttpHandler)