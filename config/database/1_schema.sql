
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

-- Table: log_groups
create table if not exists log_groups (
	id varchar(255) primary key,
	log_type varchar(100) not null,
	logged_at timestamp default current_timestamp,
	logged_by varchar(255) not null,
	path_name varchar(255) not null,
	project_id varchar(255) not null,
	payload jsonb not null,
	constraint projects_log_groups foreign key (project_id)
	references projects(id) on delete cascade,
	constraint jwt_secrets_log_groups foreign key (logged_by)
	references jwt_secrets(id) on delete cascade
) 
CREATE INDEX idx_log_groups ON log_groups (path_name, log_type)
