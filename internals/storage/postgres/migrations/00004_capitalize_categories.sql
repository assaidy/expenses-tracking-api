-- +goose Up
UPDATE categories
SET name = INITCAP(name);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION capitalize_name_in_categories()
RETURNS TRIGGER
LANGUAGE plpgsql
AS $$
BEGIN
    NEW.name = INITCAP(NEW.name);
    RETURN NEW;
END
$$;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION capitalize_category_in_expenses()
RETURNS TRIGGER
LANGUAGE plpgsql
AS $$
BEGIN
    NEW.category = INITCAP(NEW.category);
    RETURN NEW;
END
$$;
-- +goose StatementEnd

CREATE TRIGGER tr_capitalize_name_in_categories
BEFORE INSERT OR UPDATE 
ON categories
FOR EACH ROW
EXECUTE FUNCTION capitalize_name_in_categories();

CREATE TRIGGER tr_capitalize_category_in_expenses
BEFORE INSERT OR UPDATE 
ON expenses
FOR EACH ROW
EXECUTE FUNCTION capitalize_category_in_expenses();

-- +goose Down
-- cascade will remove related triggers
DROP FUNCTION IF EXISTS capitalize_name_in_categories CASCADE; 
DROP FUNCTION IF EXISTS capitalize_category_in_expenses CASCADE; 

