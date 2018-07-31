package schema

// Schema to be used in the graphql handler
var Schema = `
	schema {
		query: Query
	}

	type Query {
		viewer: Viewer
		tasks: [String!]!
	}

	# // type Mutation {
	# // 	createTask(input: CreateTaskInput!):CreateTaskPayload
	# // }

	# // input CreateTaskInput {

	# // }



	type Viewer {
		id: ID!
		name: String!
		email: String!
		avatarURL: String
		createdAt: String!
		updatedAt: String!
	}

	type Task {
		id: ID!
		user_id: ID!
		# title: String!
		# instructions: String!
		# createdAt: String!
		# updatedAt: String!
	}

`
