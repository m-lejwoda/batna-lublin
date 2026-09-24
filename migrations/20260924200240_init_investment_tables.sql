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
)

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
)

CREATE TABLE parking_place(
  id SERIAL PRIMARY KEY UNIQUE
  parking_place_external_id VARCHAR(255) UNIQUE
  parking_type VARCHAR(20)
  floor_area INT
  link VARCHAR(255)
  investment_id INT NOT NULL
)

CREATE TABLE flat_status(
  price BIGINT
  currency VARCHAR(10)
  status VARCHAR(20)
  flat_id INT NOT NULL
  created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
)

CREATE TABLE parking_place_status(
  price BIGINT
  currency VARCHAR(10)
  status VARCHAR(20)
  created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
  parking_place_id INT NOT NULL
)



-- +goose Down
