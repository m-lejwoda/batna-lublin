-- +goose Up
CREATE TABLE investment (
  id SERIAL PRIMARY KEY UNIQUE
  name VARCHAR(255)
  office_location VARCHAR(255)
  phone VARCHAR(20)
  underground_parking_place BOOLEAN
  surface_parking_place BOOLEAN
  garage BOOLEAN
  finish_date DATE NOT NULL
);

CREATE TABLE flat(
  id SERIAL PRIMARY KEY UNIQUE
  flat_external_id VARCHAR(255) UNIQUE 
  floor_area INT
  layout VARCHAR(512)
  floor_num INT
  flat_number VARCHAR(20)
  building_number VARCHAR(20)
  rooms_number INT
  balcony BOOLEAN
  flat_link VARCHAR(512)
  investment_id INT NOT NULL  

  CONSTRAINT fk_flat_investment FOREIGN KEY (investment_id) REFERENCES investment(id)
);

CREATE TABLE parking_place(
  id SERIAL PRIMARY KEY UNIQUE
  parking_place_external_id VARCHAR(255) UNIQUE
  parking_type VARCHAR(20)
  floor_area INT
  link VARCHAR(255)
  investment_id INT NOT NULL

  CONSTRAINT fk_parking_place_investment FOREIGN KEY (investment_id) REFERENCES investment(id)
);

CREATE TABLE flat_status(
  price BIGINT
  currency VARCHAR(10)
  status VARCHAR(20) CHECK (status in ("free", "reserved", "sold"))
  flat_id INT NOT NULL
  created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL

  CONSTRAINT fk_flat_status_flat FOREIGN KEY (flat_id) REFERENCES flat(id)
);

CREATE TABLE parking_place_status(
  price BIGINT
  currency VARCHAR(10)
  status VARCHAR(20) CHECK (status in ("free", "reserved", "sold"))
  created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
  parking_place_id INT NOT NULL

  CONSTRAINT fk_parking_place_status_parking_place FOREIGN KEY (parking_place_id) REFERENCES parking_place(id)
);

CREATE TABLE flat_comment(
  comments TEXT
  importanancy INT CHECK (importancy <= 10)
  flat_id INT
  
  CONSTRAINT fk_flat_comments_flat FOREIGN (flat_id) REFERENCES flat(id)
);

CREATE TABLE investment_comment(
  comments Text 
  importanncy INT CHECK (importanncy <= 10)
  investment_id INT

  CONSTRAINT fk_investment_comment_investment FOREIGN KEY (investment_id) REFERENCES investment(id)
);



-- +goose Down
DROP TABLE IF EXISTS investment, flat, parking_place, parking_place_status, flat_status, flat_comment, investment_comment;
