-- Base structure for common elements
-- Note: You don't directly create a base table, but this is represented in each table using these columns.

-- Table for ObjectTypeGroup
CREATE TABLE ObjectTypeGroups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(7),  -- Assuming color is a hex code
    monogram VARCHAR(2), -- One or two letter monogram
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for ObjectType
CREATE TABLE ObjectTypes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(7),  -- Hex color
    monogram VARCHAR(2), -- One or two letter monogram
    object_type_group_id INT REFERENCES ObjectTypeGroups(id) ON DELETE SET NULL,
    workflow_id INT, -- Placeholder for Workflow reference
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Many-to-Many table for Fields associated with ObjectTypes
CREATE TABLE Fields (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    field_type VARCHAR(20) CHECK (field_type IN ('text', 'number', 'date', 'datetime', 'richtext')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ObjectTypeFields (
    object_type_id INT REFERENCES ObjectTypes(id) ON DELETE CASCADE,
    field_id INT REFERENCES Fields(id) ON DELETE CASCADE,
    PRIMARY KEY (object_type_id, field_id)
);

-- Table for Relations between ObjectTypes (can have multiple relations)
CREATE TABLE Relations (
    id SERIAL PRIMARY KEY,
    object_type_id INT REFERENCES ObjectTypes(id) ON DELETE CASCADE,
    related_object_type_group_id INT REFERENCES ObjectTypeGroups(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for Forms
CREATE TABLE Forms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    object_type_id INT REFERENCES ObjectTypes(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for Form Elements
CREATE TABLE FormElements (
    id SERIAL PRIMARY KEY,
    form_id INT REFERENCES Forms(id) ON DELETE CASCADE,
    field_id INT REFERENCES Fields(id),  -- Only needed if this element represents a field
    relation_id INT REFERENCES Relations(id), -- If this element represents a relation
    element_type VARCHAR(20) NOT NULL, -- E.g., 'field', 'relation', 'text_block', 'button'
    position INT, -- Order of the element within the form
    properties JSONB DEFAULT '{}', -- Additional properties like "read-only" stored as JSON
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Example index for sorting FormElements by position
CREATE INDEX idx_form_elements_position ON FormElements(form_id, position);

