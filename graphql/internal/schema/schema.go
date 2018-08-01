package schema

// Schema to be used in the graphql handler
var Schema = `
	schema {
		query: Query
	}
	type Query {
		viewer: Viewer
	}

	type Viewer {
		id: ID!
		name: String!
		email: String!
		avatarURL: String
		createdAt: String!
		updatedAt: String!
	}
`
