CREATE DATABASE vehicle_db;

USE vehicle_db;

CREATE TABLE cars (
	license_plate VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    build_date DATE NOT NULL,
    odometer_value INT UNSIGNED NOT NULL DEFAULT (0),
    odometer_type ENUM('miles', 'kilometers') NOT NULL,
    transmission ENUM('automatic', 'manual') NOT NULL,
    engine_type VARCHAR(255) NOT NULL,
	retired ENUM('yes', 'no') NOT NULL,
    PRIMARY KEY (license_plate)
);

INSERT INTO cars 
		( license_plate, brand, model, build_date, odometer_value, odometer_type, transmission, engine_type, retired )
	VALUES
		( '43-PN-JK','Toyota','Corolla','1996-07-25', 87745, 'kilometers', 'manual', '2,0L L4 DOHC 16 valves VVT-iE', 'no'), 
        ( '21-VV-3B','Ford','Escort 95 (Mk6)','1995-01-01', 688937, 'miles', 'automatic', '1.8TD 1.753 cm³ 4-in-lijn 8v 90pk 178Nm', 'yes'),
        ( '7-BWE-8J','KIA','Picanto','2012-05-22', 87745, 'kilometers', 'manual', '1.0 CVVT R-SportbyKia', 'no');
        