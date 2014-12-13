CREATE TABLE item_group (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);
CREATE TABLE contact (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    vendor BOOLEAN NOT NULL,
    supplier BOOLEAN NOT NULL
);
CREATE TABLE unit_of_measure (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE tenant (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);



CREATE TABLE item (
    id SERIAL PRIMARY KEY,
    item_group_id INTEGER REFERENCES item_group NOT NULL,
    code VARCHAR(255) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    unit_of_measure_id INTEGER REFERENCES unit_of_measure NOT NULL,
    comments TEXT NOT NULL,
    component BOOLEAN NOT NULL,
    assembly BOOLEAN NOT NULL
);



CREATE TABLE bill_of_material (
    id SERIAL PRIMARY KEY,
    item_id INTEGER REFERENCES item NOT NULL
);

CREATE TABLE bill_of_material_line (
    id SERIAL PRIMARY KEY,
    bill_of_material_id INTEGER REFERENCES bill_of_material NOT NULL,
    item_id INTEGER REFERENCES item NOT NULL,
    used_qty INTEGER NOT NULL,
    waste_qty INTEGER NOT NULL
);


CREATE TABLE purchase (
    id SERIAL PRIMARY KEY,
    supplier_id INTEGER REFERENCES contact NOT NULL,
    date TIMESTAMP NOT NULL,
    approved BOOLEAN NOT NULL,
    approved_by INTEGER NOT NULL,
    approved_on TIMESTAMP NOT NULL

);

CREATE TABLE purchase_line (
    id SERIAL PRIMARY KEY,
    item_id INTEGER REFERENCES item NOT NULL,
    description text NOT NULL,
    qty INTEGER NOT NULL,
    unit_cost MONEY NOT NULL
);
CREATE TABLE receipt (
    id SERIAL PRIMARY KEY,
    purchase_id INTEGER REFERENCES purchase NOT NULL,
    date TIMESTAMP NOT NULL
);
CREATE TABLE receipt_line (
    id SERIAL PRIMARY KEY,
    receipt_id INTEGER REFERENCES receipt NOT NULL,
    purchase_line_id INTEGER REFERENCES purchase_line NOT NULL,
    item_id INTEGER REFERENCES item NOT NULL,
    qty INTEGER NOT NULL
);
CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    item_id INTEGER REFERENCES item NOT NULL,
    qty INTEGER NOT NULL,
    comments TEXT NOT NULL,
    date TIMESTAMP NOT NULL,
    type_id INTEGER NOT NULL
);
