# Http testing

### Example
```go
// Init router 
var e *gin.Engine = // get form your gin router 
var testRouter = httptestutil.NewRouter(e)
...
    headers := map[string]string{"Authorization": "Bearer eyJh...."}
    rr := testRouter.POST(t, "/blah", `{"name":"my name"}`, headers)
    if rr.Code != http.StatusOK {
        t.Errorf("failed response code, buddy ...")
    }
```

#### Options
```go
var defaultHeaders = map[string]string{"Authorization": "Bearer eyJh...."}
var testRouter = httptestutil.NewRouter(e).BasePath("/api/your-service").Headers(defaultHeaders)
```