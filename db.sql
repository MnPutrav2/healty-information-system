CREATE TYPE role_type AS ENUM ('Admin', 'User');
CREATE TYPE gender_type AS ENUM ('Laki - laki', 'Perempuan');
CREATE TYPE wedding_type AS ENUM ('Menikah', 'Belum menikah', 'Cerai', 'Cerai mati');
CREATE TYPE religion_type AS ENUM ('Islam', 'Kristen', 'Budha', 'Hindu', 'Khonghucu');
CREATE TYPE relationship_type AS ENUM ('Ibu', 'Ayah', 'Saudara', 'Teman');
CREATE TYPE awareness_type AS ENUM ('Compos Mentis', 'Somnolence', 'Sopor', 'Coma', 'Alert', 'Confusion', 'Voice', 'Pain', 'Unresponsive');
CREATE TYPE level_type AS ENUM ('INFO', 'WARN', 'ERROR');
CREATE TYPE payment_type AS ENUM ('UMUM', 'BPJS');
CREATE TYPE education_type AS ENUM ('Tidak/Belum sekolah', 'SD/Sederajat', 'SMP/Sederajat', 'SMA/Sederajat', 'Diploma (D1 - D4)', 'Sarjana (S1)', 'Magister (S2)', 'Doctor (S3)');
CREATE TYPE category_type AS ENUM ('Obat', 'BHP');
CREATE TYPE unit_type AS ENUM ('mg', 'ml');
CREATE TYPE recipe_type AS ENUM ('compound', 'common');
CREATE TYPE check_status AS ENUM ('Sudah', 'Belum');
CREATE TYPE services_type AS ENUM ('Ralan', 'Ranap');

CREATE TABLE policlinic (
    id VARCHAR(6) PRIMARY KEY,
    name VARCHAR(50),
    satu_sehat_id TEXT,
    poli_type VARCHAR(50)
);

CREATE TABLE employees (
    id VARCHAR(6) PRIMARY KEY,
    name VARCHAR(50),
    gender gender_type DEFAULT 'Laki - laki',
    birth_place VARCHAR(20),
    birth_date DATE,
    address TEXT,
    village INT,
    district INT,
    regencie INT,
    province INT,
    nik VARCHAR(16),
    bpjs VARCHAR(20),
    npwp VARCHAR(20),
    phone_number VARCHAR(13),
    email VARCHAR(50)
);

CREATE TABLE doctors (
    id VARCHAR(6) PRIMARY KEY,
    employee_id VARCHAR(6) REFERENCES employees(id),
    specialis BOOLEAN DEFAULT TRUE,
    name VARCHAR(20),
    type_specialis VARCHAR(50)
);

CREATE TABLE patients (
    medical_record VARCHAR(6) PRIMARY KEY,
    name VARCHAR(50),
    gender gender_type DEFAULT 'Laki - laki',
    wedding wedding_type DEFAULT 'Menikah',
    religion religion_type DEFAULT 'Islam',
    education education_type DEFAULT 'Sarjana (S1)',
    birth_place VARCHAR(20),
    birth_date DATE,
    work VARCHAR(20),
    address TEXT,
    village BIGINT,
    district INT,
    regencie INT,
    province INT,
    nik VARCHAR(16),
    bpjs VARCHAR(20),
    phone_number VARCHAR(13),
    parent_name VARCHAR(50),
    relationship relationship_type DEFAULT 'Ibu',
    parent_gender gender_type DEFAULT 'Perempuan'
);

CREATE TABLE ambulatory_care (
    care_number VARCHAR(14) REFERENCES registration(care_number) PRIMARY KEY,
    medical_record VARCHAR(6) REFERENCES patients(medical_record),
    date TIMESTAMP,
    body_temperature INT, 
    tension VARCHAR(7),
    pulse INT,
    respiration INT,
    height INT,
    weight INT,
    spo2 INT,
    gcs INT,
    awareness awareness_type DEFAULT 'Compos Mentis',
    complaint TEXT,
    examination TEXT,
    allergy TEXT,
    followup TEXT,
    assessment TEXT,
    instructions TEXT,
    evaluation TEXT,
    officer VARCHAR(6) REFERENCES employees(id)
);

CREATE TABLE users (
    id VARCHAR(6) PRIMARY KEY,
    employee_id VARCHAR(6) REFERENCES employees(id),
    username VARCHAR(20),
    role role_type DEFAULT 'User',
    password VARCHAR(100)
);

CREATE TABLE session_token (
    id SERIAL PRIMARY KEY,
    users_id VARCHAR(6) REFERENCES users(id),
    token VARCHAR(70),
    expired DATE
);

CREATE TABLE user_pages (
    id SERIAL PRIMARY KEY,
    users_id VARCHAR(6) REFERENCES users(id),
    path_group VARCHAR(20),
    name VARCHAR(30),
    path VARCHAR(100)
);

CREATE TABLE satu_sehat_token (
    id SERIAL PRIMARY KEY,
    token VARCHAR(50),
    expired TIMESTAMP
);

CREATE TABLE logs (
    id SERIAL PRIMARY KEY,
    users_id VARCHAR(6),
    level level_type DEFAULT 'INFO',
    message TEXT,
    path VARCHAR(100),
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE registration (
    care_number VARCHAR(14) PRIMARY KEY,
    register_number VARCHAR(3),
    register_date DATE,
    medical_record VARCHAR(6) REFERENCES patients(medical_record),
    payment_method payment_type DEFAULT 'BPJS',
    policlinic VARCHAR(3) REFERENCES policlinic(id),
    doctor VARCHAR(6) REFERENCES doctors(id),
    status check_status DEFAULT 'Belum'
);

CREATE TABLE distributor (
    id VARCHAR(6) PRIMARY KEY,
    name VARCHAR(50),
    address TEXT
);

CREATE TABLE drug_datas (
    id VARCHAR(6) PRIMARY KEY,
    name VARCHAR(50),
    distributor VARCHAR(6) REFERENCES distributor(id),
    capacity INT,
    fill INT,
    unit unit_type DEFAULT 'mg',
    composition VARCHAR(100),
    price INT,
    category category_type DEFAULT 'Obat',
    expired_date DATE
);

CREATE TABLE drug_stock (
    drug_id VARCHAR(6) REFERENCES drug_datas(id),
    stock INT
);

CREATE TABLE recipes (
    recipe_id VARCHAR(14) PRIMARY KEY,
    care_number VARCHAR(14) REFERENCES registration(care_number),
    date TIMESTAMP,
    validate TIMESTAMP,
    validate_status BOOLEAN,
    handover TIMESTAMP
);

CREATE TABLE detail_recipes (
    recipe_id VARCHAR(14) REFERENCES recipes(recipe_id) ON DELETE CASCADE,
    drug_id VARCHAR(6) REFERENCES drug_datas(id),
    validate_status BOOLEAN,
    compound_name VARCHAR(3),
    compound_value INT,
    recipe_type recipe_type DEFAULT 'common', 
    value INT,
    use VARCHAR(20),
    embalming INT,
    tuslah INT,
    total_price INT
);

CREATE TABLE labolatorium_datas (
    id VARCHAR(6) PRIMARY KEY,
    name VARCHAR(80),
    referral_fee INT,
    officer_fee INT,
    management INT,
    price INT
);

CREATE TABLE labolatorium_template (
    template_id INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    id VARCHAR(6) REFERENCES labolatorium_datas(id) ON DELETE CASCADE,
    name VARCHAR(10),
    unit VARCHAR(10),
    normal_value TEXT
);

CREATE TABLE labolatorium_request (
    lab_id VARCHAR(14) PRIMARY KEY,
    care_number VARCHAR(14) REFERENCES registration(care_number),
    referral_id VARCHAR(6) REFERENCES doctors(id),
    officer_id VARCHAR(6) REFERENCES employees(id),
    request_date TIMESTAMP,
    sample TIMESTAMP,
    validate BOOLEAN,
    validate_date TIMESTAMP
);

CREATE TABLE labolatorium_detail (
    lab_request VARCHAR(14) REFERENCES labolatorium_request(lab_id) ON DELETE CASCADE,
    lab_id VARCHAR(6) REFERENCES labolatorium_datas(id),
    referral_fee INT,
    officer_fee INT,
    management INT,
    price INT
);

CREATE TABLE labolatorium_detail_template (
    lab_request VARCHAR(14) REFERENCES labolatorium_request(lab_id) ON DELETE CASCADE,
    lab_id VARCHAR(6) REFERENCES labolatorium_datas(id),
    template_id VARCHAR(6),
    value VARCHAR(20)
);

CREATE TABLE patient_account (
    patient_id VARCHAR(6) REFERENCES patients(medical_record),
    username VARCHAR(20),
    password VARCHAR(100)
);

CREATE TABLE patient_session_token (
    id SERIAL PRIMARY KEY,
    patient_id VARCHAR(6) REFERENCES patients(medical_record),
    token VARCHAR(70),
    expired DATE
);

CREATE TABLE schedule (
    id SERIAL PRIMARY KEY,
    doctor_id VARCHAR(6) REFERENCES doctors(id),
    day VARCHAR(15),
    first_time VARCHAR(10),
    last_time VARCHAR(10)
);

CREATE TABLE examination_cost (
    id VARCHAR(6) PRIMARY KEY,
    examination_name VARCHAR(80),
    payment_method payment_type DEFAULT 'UMUM',
    doctor_cost INT,
    nurse_cost INT,
    management_cost INT
);

CREATE TABLE examination (
    id SERIAL PRIMARY KEY,
    care_number VARCHAR(14) REFERENCES registration(care_number),
    examination VARCHAR(6) REFERENCES examination_cost(id),
    doctor_id VARCHAR(6) REFERENCES doctors(id),
    nurse_id VARCHAR(6) REFERENCES employees(id),
    service_type services_type DEFAULT 'Ralan',
    date TIMESTAMP,
    total INT
);