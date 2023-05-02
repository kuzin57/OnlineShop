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

CREATE OR REPLACE TRIGGER user_update_on_delete
AFTER DELETE ON bshop.user
FOR EACH ROW EXECUTE FUNCTION bshop.write_user_history_on_delete();

------------------------------------------------------------------------

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

CREATE OR REPLACE TRIGGER user_update_on_insert_or_update
AFTER INSERT OR UPDATE ON bshop.user
FOR EACH ROW EXECUTE FUNCTION bshop.write_user_history_on_insert_or_update();

---------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION bshop.update_product_rating_and_amount_rating()
RETURNS TRIGGER
AS $$
BEGIN
    UPDATE bshop.product
    SET rating = (rating * rating_amount + NEW.mark) / (rating_amount + 1),
        rating_amount = rating_amount + 1
    WHERE product_id = NEW.product_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER product_rating_update
AFTER INSERT ON bshop.review
FOR EACH ROW EXECUTE FUNCTION bshop.write_user_history_on_insert_or_update();

-------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION bshop.write_product_history_on_delete()
RETURNS TRIGGER
AS $$
BEGIN
    INSERT INTO bshop.user_history VALUES (
        OLD.product_id,
        OLD.category,
        OLD.name,
        OLD.brand,
        OLD.price,
        OLD.available,
        OLD.rating,
        OLD.rating_amount,
        NOW()
    );
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER product_update_on_delete
AFTER DELETE ON bshop.product
FOR EACH ROW EXECUTE FUNCTION bshop.write_product_history_on_delete();

---------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION bshop.write_product_history_on_insert_or_update()
RETURNS TRIGGER
AS $$
BEGIN
    INSERT INTO bshop.user_history VALUES (
        NEW.product_id,
        NEW.category,
        NEW.name,
        NEW.brand,
        NEW.price,
        NEW.available,
        NEW.rating,
        NEW.rating_amount,
        NOW()
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER product_update_on_delete
AFTER DELETE ON bshop.product
FOR EACH ROW EXECUTE FUNCTION bshop.write_product_history_on_insert_or_update();