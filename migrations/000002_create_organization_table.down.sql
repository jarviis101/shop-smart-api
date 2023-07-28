ALTER TABLE users DROP CONSTRAINT fk_organization;
ALTER TABLE organizations DROP CONSTRAINT fk_user;
DROP TABLE IF EXISTS organizations;