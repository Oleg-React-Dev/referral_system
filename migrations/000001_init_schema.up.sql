
CREATE TABLE IF NOT EXISTS users (
    user_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();


CREATE TABLE referral_codes (
    code TEXT PRIMARY KEY,
    user_id UUID REFERENCES users(user_id) UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


CREATE OR REPLACE FUNCTION generate_unique_referral_code() RETURNS TRIGGER AS $$
DECLARE
    generated_code TEXT;
BEGIN
    LOOP
        generated_code := substring(md5(random()::text), 1, 8);
        IF NOT EXISTS (SELECT 1 FROM referral_codes WHERE code = generated_code) THEN
            NEW.code := generated_code;
            EXIT;
        END IF;
    END LOOP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_referral_code
BEFORE INSERT ON referral_codes
FOR EACH ROW
WHEN (NEW.code IS NULL)
EXECUTE FUNCTION generate_unique_referral_code();


CREATE TABLE referrals (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    referrer_id UUID REFERENCES users(user_id),
    referral_id UUID REFERENCES users(user_id) UNIQUE,
    referred_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


CREATE OR REPLACE FUNCTION add_referral(p_referral_id UUID, p_referral_code TEXT)
RETURNS BOOLEAN AS $$
DECLARE
    v_referrer_id UUID;
BEGIN
    SELECT user_id INTO v_referrer_id
    FROM referral_codes
    WHERE code = p_referral_code
      AND expires_at > CURRENT_TIMESTAMP;

    IF v_referrer_id IS NULL THEN
        RETURN FALSE;
    END IF;

    INSERT INTO referrals (referral_id, referrer_id, referred_at)
    VALUES (p_referral_id, v_referrer_id, CURRENT_TIMESTAMP);

    RETURN TRUE;
EXCEPTION
    WHEN OTHERS THEN
        RETURN FALSE;
END;
$$ LANGUAGE plpgsql;
