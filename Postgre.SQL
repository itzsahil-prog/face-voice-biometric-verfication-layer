-- Users table (Identity Only)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone VARCHAR(20) UNIQUE,
    email VARCHAR(255) UNIQUE,
    device_fingerprint VARCHAR(255),
    trust_score INT DEFAULT 50, -- Dynamic trust score
    biometric_enabled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Biometric Sessions (Temporary Trust Tokens)
-- Stores the BTT hash, not the biometric itself
CREATE TABLE biometric_sessions (
    session_id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    challenge_hash VARCHAR(255), -- Hash of the phrase spoken
    expires_at TIMESTAMP,
    status VARCHAR(20), -- VERIFIED, FAILED, EXPIRED
    created_at TIMESTAMP DEFAULT NOW()
);

-- Transactions (High Security)
CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    amount DECIMAL(19, 4),
    risk_score INT,
    risk_factors JSONB, -- e.g., ["high_amount", "new_location"]
    status VARCHAR(20), -- PENDING, AUTHORIZED, DECLINED
    created_at TIMESTAMP DEFAULT NOW()
);