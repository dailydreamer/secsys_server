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
  nick_name text DEFAULT "",
  phone text DEFAULT "",
  email text DEFAULT "",
  avator text DEFAULT "",
  created timestamp with time zone DEFAULT transaction_timestamp(),
  modified timestamp with time zone DEFAULT transaction_timestamp(),
  password text,
  com_name text,
  com_field text DEFAULT "",
  com_man text DEFAULT "",
  com_phone text DEFAULT "",
  com_regnum text DEFAULT "",
  com_regcap decimal DEFAULT 0,
  com_capreport decimal DEFAULT 0,
  com_batch text DEFAULT "",
  com_level text DEFAULT "" ,
  appli_date text DEFAULT "",
  appli_level text DEFAULT "",
  appli_result text DEFAULT "",
  certf_date text DEFAULT "",
  certf_num text DEFAULT "",
  verif_date text DEFAULT "",
  verif_result text DEFAULT "",
  com_turnover decimal DEFAULT 0,
  com_area decimal DEFAULT 0,
  police_num integer DEFAULT 0,
  police_duty integer DEFAULT 0,
  police_cancel integer DEFAULT 0,
  police_dutycancel decimal DEFAULT 0,
  list_duty integer DEFAULT 0,
  list_dutycancel decimal DEFAULT 0,
  emp_num integer DEFAULT 0,
  emp_contract integer DEFAULT 0,
  emp_lccr decimal DEFAULT 0,
  cont_num integer DEFAULT 0,
  cont_vac decimal DEFAULT 0,
  cont_samptnum integer DEFAULT 0,
  cont_sampfnum integer DEFAULT 0,
  cont_sampvac decimal DEFAULT 0,
  emp_sep integer DEFAULT 0,
  emp_seprate decimal DEFAULT 0,
  list_certrate decimal DEFAULT 0,
  list_sampcertrate decimal DEFAULT 0,
  emp_ssemanum integer DEFAULT 0,
  emp_ssemarate decimal DEFAULT 0,
  emp_semanum integer DEFAULT 0,
  emp_semarate decimal DEFAULT 0,
  emp_jsenum integer DEFAULT 0,
  emp_jserate decimal DEFAULT 0,
  train_period integer DEFAULT 0,
  com_salary decimal DEFAULT 0,
  train_funds decimal DEFAULT 0,
  train_fundsrate decimal DEFAULT 0,
  com_comins integer DEFAULT 0,
  com_sosec integer DEFAULT 0,
  com_sosecrate decimal DEFAULT 0,
  com_party text DEFAULT "",
  com_youth text DEFAULT "",
  com_union text DEFAULT "",
  com_crime text DEFAULT "",
  com_acc text DEFAULT "",
  com_mwgs text DEFAULT "",
  com_license text DEFAULT ""
);
CREATE UNIQUE INDEX users_phone_idx ON users (phone);
CREATE UNIQUE INDEX users_com_name_idx ON users (com_name);

CREATE TRIGGER update_users_modified
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE update_modified_column();


CREATE TABLE logs (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id uuid,
  created timestamp with time zone DEFAULT transaction_timestamp(),
  com_name text,
  ip text,
  address text DEFAULT "",
  status text DEFAULT ""
);
ALTER TABLE logs add CONSTRAINT fk_logs_user_id foreign key(user_id) references users(id);


CREATE TABLE messages (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id uuid,
  created timestamp with time zone DEFAULT transaction_timestamp(),
  com_name text,
  message text
);
ALTER TABLE messages add CONSTRAINT fk_messages_user_id foreign key(user_id) references users(id);


CREATE TABLE contracts (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id uuid,
  com_name text,
  contract_no text DEFAULT "",
  project_name text DEFAULT "",
  com_field text DEFAULT "",
  customer_name text DEFAULT "",
  customer_type text DEFAULT "",
  people_num integer DEFAULT 0,
  start_time timestamp with time zone DEFAULT NULL,
  end_time timestamp with time zone DEFAULT NULL,
  unit_price decimal DEFAULT 0,
  total_price decimal DEFAULT 0,
  income text DEFAULT "",
  created timestamp with time zone DEFAULT transaction_timestamp(),
  modified timestamp with time zone DEFAULT transaction_timestamp()
);
ALTER TABLE contracts add CONSTRAINT fk_contracts_user_id foreign key(user_id) references users(id);

CREATE TRIGGER update_contracts_modified
BEFORE UPDATE ON contracts
FOR EACH ROW EXECUTE PROCEDURE update_modified_column();


CREATE TABLE scores (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id uuid,
  com_name text,
  year text DEFAULT "",
  standard text DEFAULT "",
  score_no text DEFAULT "",
  score_type text DEFAULT "",
  satisfied text DEFAULT "",
  score decimal DEFAULT 0,
  reason text DEFAULT "",
  created timestamp with time zone DEFAULT transaction_timestamp(),
  modified timestamp with time zone DEFAULT transaction_timestamp()
);
ALTER TABLE scores add CONSTRAINT fk_scores_user_id foreign key(user_id) references users(id);

CREATE TRIGGER update_scores_modified
BEFORE UPDATE ON scores
FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
