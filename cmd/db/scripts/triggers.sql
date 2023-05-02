CREATE OR REPLACE FUNCTION bshop.write_user_history_on_delete()
RETURNS TRIGGER
AS $$
BEGIN
    INSERT INTO bshop.user_history VALUES (
        OLD.user_id,
        OLD.firstname,
        OLD.surname,
        OLD.phone_number,
        OLD.email,
        OLD.birthdate,
        OLD.hashed_password,
        NOW()
    );
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER users_update_on_delete
AFTER DELETE ON bshop.user
FOR EACH ROW EXECUTE FUNCTION bshop.write_users_history_on_delete();


CREATE OR REPLACE FUNCTION bshop.write_user_history_on_insert_or_update()
RETURNS TRIGGER
AS $$
BEGIN
    INSERT INTO bshop.user_history VALUES (
        NEW.user_id,
        NEW.firstname,
        NEW.surname,
        NEW.phone_number,
        NEW.email,
        NEW.birthdate,
        NEW.hashed_password,
        NOW()
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER users_update_on_insert_or_update
AFTER INSERT OR UPDATE ON bshop.user
FOR EACH ROW EXECUTE FUNCTION bshop.write_users_history_on_insert_or_update();