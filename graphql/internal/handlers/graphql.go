package handlers

import (
	"fmt"
	"net/http"

	"github.com/abbeyhrt/keep-up/graphql/internal/database"
	"github.com/abbeyhrt/keep-up/graphql/internal/resolver"
	"github.com/abbeyhrt/keep-up/graphql/internal/schema"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// GraphQLHandler will be the handler for all graphql queries
func GraphQLHandler(store database.DAL) *relay.Handler {
	schema := graphql.MustParseSchema(schema.Schema, resolver.New(store))
	return &relay.Handler{Schema: schema}
}

func GraphiqlHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(page)
	})
}

var graphiqlVersion = "0.11.2"
var page = []byte(fmt.Sprintf(`<!--
The request to this GraphQL server provided the header "Accept: text/html"
and as a result has been presented GraphiQL - an in-browser IDE for
exploring GraphQL.
If you wish to receive JSON, provide the header "Accept: application/json" or
add "&raw" to the end of the URL within a browser.
-->
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8" />
		<title>GraphiQL</title>
		<meta name="robots" content="noindex" />
		<style>
			html, body {
				height: 100%%;
				margin: 0;
				overflow: hidden;
				width: 100%%;
			}
		</style>
		<link href="//cdn.jsdelivr.net/npm/graphiql@%[1]s/graphiql.css" rel="stylesheet" />
		<script src="//cdn.jsdelivr.net/fetch/0.9.0/fetch.min.js"></script>
		<script src="//cdn.jsdelivr.net/react/15.4.2/react.min.js"></script>
		<script src="//cdn.jsdelivr.net/react/15.4.2/react-dom.min.js"></script>
		<script src="//cdn.jsdelivr.net/npm/graphiql@%[1]s/graphiql.min.js"></script>
	</head>
	<body>
		<div id="root" style="height: 100vh;">Loading...</div>
		<script>
			// Collect the URL parameters
			var parameters = {};
			window.location.search.substr(1).split('&').forEach(function (entry) {
			var eq = entry.indexOf('=');
			if (eq >= 0) {
				parameters[decodeURIComponent(entry.slice(0, eq))] =
				decodeURIComponent(entry.slice(eq + 1));
			}
			});
			// Produce a Location query string from a parameter object.
			function locationQuery(params) {
				return '?' + Object.keys(params).filter(function (key) {
					return Boolean(params[key]);
				}).map(function (key) {
					return encodeURIComponent(key) + '=' +
					encodeURIComponent(params[key]);
				}).join('&');
			}
			// Derive a fetch URL from the current URL, sans the GraphQL parameters.
			var graphqlParamNames = {
				query: true,
				variables: true,
				operationName: true
			};
			var otherParams = {};
			for (var k in parameters) {
				if (parameters.hasOwnProperty(k) && graphqlParamNames[k] !== true) {
					otherParams[k] = parameters[k];
				}
			}
			var fetchURL = locationQuery(otherParams);
			// Defines a GraphQL fetcher using the fetch API.
			function graphQLFetcher(graphQLParams) {
			return fetch("/graphql", {
				method: 'post',
				headers: {
					'Accept': 'application/json',
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(graphQLParams),
				credentials: 'include',
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
			// When the query and variables string is edited, update the URL bar so
			// that it can be easily shared.
			function onEditQuery(newQuery) {
				parameters.query = newQuery;
				updateURL();
			}
			function onEditVariables(newVariables) {
				parameters.variables = newVariables;
				updateURL();
			}
			function onEditOperationName(newOperationName) {
				parameters.operationName = newOperationName;
				updateURL();
			}
			function updateURL() {
				history.replaceState(null, null, locationQuery(parameters));
			}
			// Render <GraphiQL /> into the body.
			ReactDOM.render(
				React.createElement(GraphiQL, {
					fetcher: graphQLFetcher,
					onEditQuery: onEditQuery,
					onEditVariables: onEditVariables,
					onEditOperationName: onEditOperationName,
				}),
				document.getElementById("root")
			);
		</script>
	</body>
</html>
`, graphiqlVersion))
