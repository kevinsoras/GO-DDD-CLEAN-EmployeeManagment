-- 🔹 Definir ENUM para tipo de persona
CREATE TYPE person_type AS ENUM ('NATURAL', 'JURIDICAL'); -- NATURAL = Natural, JURIDICAL = Jurídica

-- 🔹 Tabla principal: PERSONS
CREATE TABLE persons (
    person_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    person_type person_type NOT NULL,         -- Usa ENUM en vez de CHAR(1)
    email VARCHAR(100),
    phone VARCHAR(20),
    address TEXT,
    country VARCHAR(100) DEFAULT 'PERU',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- 🔹 Tabla: PERSONAS NATURALES
CREATE TABLE natural_persons (
    person_id UUID PRIMARY KEY REFERENCES persons(person_id) ON DELETE CASCADE,
    document_number VARCHAR(20) NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name_paternal VARCHAR(100) NOT NULL,
    last_name_maternal VARCHAR(100),
    birth_date DATE,
    gender CHAR(1) CHECK (gender IN ('M','F','O')) -- M = Masculino, F = Femenino, O = Otro
);

-- 🔹 Tabla: PERSONAS JURÍDICAS
CREATE TABLE juridical_persons (
    person_id UUID PRIMARY KEY REFERENCES persons(person_id) ON DELETE CASCADE,
    document_number VARCHAR(20) NOT NULL UNIQUE,
    business_name VARCHAR(200) NOT NULL,       -- Razón social
    trade_name VARCHAR(200),                   -- Nombre comercial
    constitution_date DATE,                    -- Fecha de constitución
    representative_name VARCHAR(200),          -- Representante legal
    representative_document VARCHAR(20)        -- Documento representante
);
