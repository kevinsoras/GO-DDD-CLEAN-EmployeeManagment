-- Tabla principal: EMPLOYEES
CREATE TABLE employees (
    employee_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    person_id UUID NOT NULL REFERENCES persons(person_id) ON DELETE CASCADE,
    salary NUMERIC(12,2) NOT NULL,
    contract_type VARCHAR(20) NOT NULL,
    position VARCHAR(100) NOT NULL,
    work_schedule VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL,
    work_location VARCHAR(100),
    bank_account VARCHAR(50),
    afp VARCHAR(50) NOT NULL,
    eps VARCHAR(50) NOT NULL,
    start_date DATE NOT NULL,
    has_cts BOOLEAN DEFAULT false,
    has_gratification BOOLEAN DEFAULT false,
    has_vacation BOOLEAN DEFAULT false,
    cts NUMERIC(12,2) DEFAULT 0,
    gratification NUMERIC(12,2) DEFAULT 0,
    vacation_days INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
