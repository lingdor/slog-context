# slog-context
slog no support for context value, not to be contacted modules in traceid. so support that package.

# context for slog

```go
package gin

import(
    "github.com/lingdor/slog-context/slog2"
)

func Request1(){

logger :=slog.With("traceid",123)
    ctx:=context.Background()
    ctx=slog2.NewContext(ctx,logger)
    gorm.ReadDB(ctx)
}
```

```go
package  gorm

import(
    "github.com/lingdor/slog-context/slog2"
)

func ReadDB(ctx context.Context){
    logger:=slog2.FromContext(ctx)
    logger.Info("db is ready")
}
// or

func ReadDB(ctx context.Context){
    logger:=slog2.FromContextOrDefault(ctx)
    logger.Info("db is ready")
}

```