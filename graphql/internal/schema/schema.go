package schema

// Schema to be used in the graphql handler
var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	type Query {
		viewer: Viewer
		home: Home
	}

	type Mutation {
		createHome(
			name: String!,
			description: String!
		): Home!
	}

	type Viewer {
		id: ID!
		name: String!
		home_id: ID
		email: String!
		avatarURL: String
		createdAt: String!
		updatedAt: String!
	}

	type Home {
		id: ID!
		name: String!
  	description: String!
	  avatarURL: String
		createdAt: String!
		updatedAt: String!
	}

`
