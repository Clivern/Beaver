# Create database
CREATE KEYSPACE IF NOT EXISTS [DatabaseName] WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

# Create client table
CREATE TABLE IF NOT EXISTS [DatabaseName].client (id UUID PRIMARY KEY, node_id INT, status VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP);

# Create channel table
CREATE TABLE IF NOT EXISTS [DatabaseName].channel (id UUID PRIMARY KEY, slug VARCHAR, name VARCHAR, type VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP);

# Create message table
CREATE TABLE IF NOT EXISTS [DatabaseName].message (id UUID PRIMARY KEY, message TEXT, from_client_id INT, to_channel_id INT, to_client_id INT, created_at TIMESTAMP, updated_at TIMESTAMP);

# Create node table
CREATE TABLE IF NOT EXISTS [DatabaseName].node (id UUID PRIMARY KEY, address VARCHAR, status VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP);

# Create client_channel relationship table
CREATE TABLE IF NOT EXISTS [DatabaseName].client_channel (id UUID PRIMARY KEY, client_id INT, channel_id INT);
