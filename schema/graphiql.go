package schema

import (
	"net/http"
)

var page = []byte(`
	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
			<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
		</head>
		<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
			<div id="graphiql" style="height: 100vh;">Loading...</div>
			<script>
				function graphQLFetcher(graphQLParams) {
					return fetch("/api?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7IkZpcnN0TmFtZSI6IkFtaXQiLCJMYXN0TmFtZSI6Ik1pdHRhbCIsIkVtYWlsSUQiOiJtaXR0bWFpbGJveEBnbWFpbC5jb20iLCJVc2VybmFtZSI6ImFrbWl0dGFsMzIxNCIsIlBhc3N3b3JkIjoiJDJhJDEwJHhGdldGSUEySjFCOWFLS1JQU1JOci53eE5kOGFBeTBjcnh3QmJSYVFNWWJXNmNYcUdGVW9hIn0sImV4cCI6MTUxOTE5NTAwMywiaW90IjoxNTE5MDIyMjAzLCJpc3MiOiJSZWNydUlUIn0.C77BASQiwv8plz4M8yh1guWQZuoq2MiB5LL1RnQPSi4", {
						method: "post",
						body: JSON.stringify(graphQLParams),
						credentials: "include",
					}).then(function (response) {
						return response.text();
					}).then(function (responseBody) {
						try {
							return JSON.parse(responseBody);
						} catch (error) {
							return responseBody;
						}
					});
				}
				ReactDOM.render(
					React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
					document.getElementById("graphiql")
				);
			</script>
		</body>
	</html>
	`)

func GraphiQLHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write(page)
}
