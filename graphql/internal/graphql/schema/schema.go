package schema

//Schema to be used in the graphql handler
var Schema = `
	type Query {
		viewer(id: ID!): Viewer
	}

	type Viewer {
		id: ID!
		name: String!
		email: String!
	}
		`
