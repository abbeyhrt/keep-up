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
	}

	type Mutation {
		createHome(
			name: String!,
			description: String!
			): Home!
		updateHome(
			home: HomeInput!
			): Home
		deleteHome(
			id: String!
			): Home
		createTask(
			title: String!,
			description: String!,
			): Task!
		updateTask(
			task: TaskInput!
			): Task!
		deleteTask(
			id: String!
			): Task
		updateUser(
			user: UserInput!
			): User!
		deleteUser(
			id: String!
			): User
	}

	type Viewer {
		id: ID!
		first_name: String!
		last_name: String!
		home: Home
		tasks: [Task!]
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

	input HomeInput {
		id: ID!
		name: String
		description: String
		avatarURL: String
	}

	type Task {
		id: ID!
	  user_id: String!
		title: String!
		description: String!
		createdAt: String!
		updatedAt: String!
	}

	input TaskInput {
		id: ID!
		user_id: String
		title: String
		description: String
	}

	type User {
		id: ID!
		email: String!
		home_id: String
		first_name: String!
		last_name: String!
		avatarURL: String
	}

	input UserInput {
		id: ID!
	  email: String
	  home_id: String
		first_name: String
	  last_name: String
	  avatarURL: String
	}


`
