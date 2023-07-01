USE recipeshelf;

CREATE TYPE diet AS ENUM
(
    'Unknown',
	'Diabetic',
	'GlutenFree',
	'Halal',
	'Hindu',
	'Kosher',
	'LowCalorie',
	'LowFat',
	'LowLactose',
	'LowSalt',
	'Vegan',
	'Vegetarian'
);

CREATE TABLE nutrition (
    id UUID primary key DEFAULT gen_random_uuid(), 
    calories FLOAT NULL, 
    carbohydrate_grams FLOAT NULL,
    cholesterol_milligrams FLOAT NULL,
    fat_grams FLOAT NULL,
    fiber_grams FLOAT NULL,
    protein_grams FLOAT NULL,
    saturated_fat_grams FLOAT NULL,
    serving_size STRING(1024) NULL,
    sodium_milligrams FLOAT NULL,
    sugar_grams FLOAT NULL,
    trans_fat_grams FLOAT NULL,
    unsaturated_fat_grams FLOAT NULL
);

CREATE TABLE recipes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    author STRING(1024) NULL,
    categories STRING(1024)[] NULL,
    cook_time INTERVAL NULL,
    cuisine STRING(1024)[] NULL,
    description STRING(4096) NULL,
    image_url STRING(2048) NULL,
    ingredients STRING(1024)[] NULL,
    instructions STRING(2048)[] NULL,
    language STRING(1024),
    name STRING(1024),
    nutrition_id UUID NULL REFERENCES nutrition,
    prep_time INTERVAL NULL,
    suitable_diets diet[] NULL,
    total_time INTERVAL NULL,
    yield STRING(1024) NULL,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE bookmarks (    
    url STRING(2048) PRIMARY KEY,
    recipe_id UUID NULL REFERENCES recipes,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name STRING(1024) UNIQUE NOT NULL,
    supertoken_key STRING(1024) UNIQUE NULL
);

CREATE TABLE user_recipes (
    user_id UUID REFERENCES users,
    recipe_id UUID REFERENCES recipes,
    view_count INT NOT NULL
);
