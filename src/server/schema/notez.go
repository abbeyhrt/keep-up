type Mutation {
		# Create a new Business Unit
		createBusinessUnit(input: CreateBusinessUnitInput!):
			CreateBusinessUnitPayload

		# Update an existing Business Unit
		updateBusinessUnit(input: UpdateBusinessUnitInput!):
			UpdateBusinessUnitPayload

		# Destroy an existing Business Unit
		destroyBusinessUnit(input: DestroyBusinessUnitInput!):
			DestroyBusinessUnitPayload

		# Create a new Portfolio
		createPortfolio(input: CreatePortfolioInput!): CreatePortfolioPayload
	}

	# Input arguments to update a new Business Unit
	input UpdateBusinessUnitInput {
		# The id of the Business Unit
		id: ID!
		# The name of the Business Unit
		name: String
		# The description for the Business Unit
		description: String
	}

	# The payload to receive after making an update to a Business Unit
	type UpdateBusinessUnitPayload {
		# The id of the Business Unit
		id: ID!
		# The name of the Business Unit
		name: String!
		# The description for the Business Unit
		description: String!
		# The time when the resource was created
		createdAt: String!
		# The time when the resource was last updated
		updatedAt: String!
	}
