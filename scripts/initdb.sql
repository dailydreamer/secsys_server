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
  email: text,
  nickname text,
  people: text,
  avator text,
  created timestamp with time zone DEFAULT transaction_timestamp(),
  modified timestamp with time zone DEFAULT transaction_timestamp(),
  password text,
  com_name: text,
  com_field: text,
  com_man: text,
  com_phone: text,
  com_regnum: text,
  com_regcap: decimal,
  com_capreport: decimal,
  com_batch: text,
  com_level: text,
  appli_date: text,
  appli_level: text,
  appli_result: text,
  certf_date: text,
  certf_num: text,
  verif_date: text,
  verif_result: text,
  com_turnover: decimal,
  com_area: decimal,
  police_num: integer,
  police_duty: integer,
  police_cancel: integer,
  police_dutycancel: decimal,
  list_duty: integer,
  list_dutycancel: decimal,
  emp_num: integer,
  emp_contract: integer,
  emp_lccr: decimal,
  cont_num: integer,
  cont_vac: decimal,
  cont_samptnum: integer,
  cont_sampfnum: integer,
  cont_sampvac: decimal,
  emp_sep: integer,
  emp_seprate: decimal,
  list_certrate: decimal,
  list_sampcertrate: decimal,
  emp_ssemanum: integer,
  emp_ssemarate: decimal,
  emp_semanum: integer,
  emp_semarate: decimal,
  emp_jsenum: integer,
  emp_jserate: decimal,
  train_period: integer,
  com_salary: decimal,
  train_funds: decimal,
  train_fundsrate: decimal,
  com_comins: integer,
  com_sosec: integer,
  com_sosecrate: decimal,
  com_party: text,
  com_youth: text,
  com_union: text,
  com_crime: text,
  com_acc: text,
  com_mwgs: text,
  com_license: text
);
CREATE UNIQUE INDEX users_phone_idx ON users (phone);

CREATE TRIGGER update_users_modified
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE update_modified_column();


CREATE TABLE logs (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  login_time timestamp with time zone DEFAULT transaction_timestamp(),
  ip text,
  address text,
  status text
);


CREATE TABLE messages (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  created timestamp with time zone DEFAULT transaction_timestamp(),
  name text,
  message text
);


CREATE TABLE contracts (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id uuid,
  com_name text,
  contract_no text,
  project_name: text,
  com_field: text,
  customer_name: text,
  customer_type: text,
  people_num: integer,
  start_time: timestamp with time zone,
  end_time: timestamp with time zone,
  unit_price: decimal,
  total_price: decimal,
  income: text
);


CREATE TABLE scores (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id uuid,
  com_name text,
  year: text,
  standard: text,
  score_no: text,
  score_type: text,
  satisfied: text,
  score: decimal,
  reason: text
);
