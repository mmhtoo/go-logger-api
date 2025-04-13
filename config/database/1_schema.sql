
-- Table: projects
CREATE TABLE IF NOT EXISTS projects (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  description VARCHAR(255) NOT NULL,
  project_type VARCHAR(100) NOT NULL,
  created_user_id VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

-- Table: jwt_secrets
create table if not exists jwt_secrets (
	id varchar(255) primary key,
	key_name varchar(100) not null,
	private_key text not null,
	public_key text not null,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp ,
	created_user_id varchar(255) not null,
	updated_user_id varchar(255) not null,
	project_id varchar(255) not null,
	is_active boolean default false,
	constraint projects_to_jwt_secrets foreign key (project_id) 
	references projects (id) on delete set null
)