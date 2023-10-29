CREATE TABLE IF NOT EXISTS users_have_roles (
  id SERIAL PRIMARY KEY,
  user_id SERIAL,
  role_id SERIAL,
  created_at TIMESTAMP,
  UNIQUE (user_id, role_id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
); 
