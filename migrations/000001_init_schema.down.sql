
DROP TRIGGER IF EXISTS update_timestamp ON users;
DROP FUNCTION IF EXISTS update_updated_at_column();

DROP TRIGGER IF EXISTS set_referral_code ON referral_codes;
DROP FUNCTION IF EXISTS generate_unique_referral_code();

DROP FUNCTION IF EXISTS add_referral;


DROP TABLE IF EXISTS referrals;
DROP TABLE IF EXISTS referral_codes;
DROP TABLE IF EXISTS users;
