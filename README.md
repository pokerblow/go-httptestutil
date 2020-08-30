# Http testing

### Example
```go
// Init router 
var e *gin.Engine = // get form your gin router 
var testRouter = httptestutil.NewRouter(e)
...
    postBody := `{"name":"my name"}`
    headers := map[string]string{"Authorization": "Bearer eyJh...."}
    rr := testRouter.Request(t, "POST", "/blah", &postBody, headers)
    if rr.Code != http.StatusOK {
        t.Errorf("failed response code, buddy ...")
    }
```