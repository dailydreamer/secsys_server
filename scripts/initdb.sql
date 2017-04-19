CREATE EXTENSION "uuid-ossp";

CREATE OR REPLACE FUNCTION update_modified_column()	
RETURNS TRIGGER AS $$
BEGIN
    NEW.modified = transaction_timestamp();
    RETURN NEW;	
END;
$$ language 'plpgsql';

CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  is_admin boolean DEFAULT false,
  phone text,
  nickname text,
  avator text,
  created timestamp with time zone DEFAULT transaction_timestamp(),
  modified timestamp with time zone DEFAULT transaction_timestamp(),
  password text
);
CREATE UNIQUE INDEX users_phone_idx ON users (phone);

CREATE TRIGGER update_users_modified 
BEFORE UPDATE ON users 
FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

