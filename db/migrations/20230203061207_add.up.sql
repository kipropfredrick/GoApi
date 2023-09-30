CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "first_name" varchar(100) NOT NULL,
    "last_name" varchar(100) NOT NULL,
    "phone" varchar(100) NOT NULL,
    "id_number" varchar(100)  NULL,
    "email" varchar(100) NOT NULL,
    "password" text  NULL,
    "plate_number" varchar(100) NULL,
    "license_number" varchar(100) NULL,
    "image_caption" varchar(100) NULL,
    "is_verified" boolean DEFAULT false
)