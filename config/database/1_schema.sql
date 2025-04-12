
-- Table: projects
CREATE TABLE IF NOT EXISTS projects (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  description VARCHAR(255) NOT NULL,
  project_type VARCHAR(100) NOT NULL,
  created_user_id VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)