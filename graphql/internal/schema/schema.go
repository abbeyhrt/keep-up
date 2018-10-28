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
		task(id: ID!): Task
		users(name: String!): [User]
	}

	type Mutation {
		createHome(
			name: String!,
			description: String!
		): Home!
		createTask(
			title: String!,
			description: String!,
			): Task!
		insertHomeID(
			homeID: String!,
			userID: String!,
			): Home!
	}

	type Viewer {
		id: ID!
		first_name: String!
		last_name: String!
		home: Home
		tasks: [Task!]!
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

	type Task {
		id: ID!
	  user_id: String!
		title: String!
		description: String!
		createdAt: String!
		updatedAt: String!
	}

	type User {
		id: ID!
		email: String!
		home_id: String
		first_name: String!
		last_name: String!
		avatarURL: String
	}

`
