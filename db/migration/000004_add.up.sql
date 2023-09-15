-- Equipment Type Enum
CREATE TYPE EquipmentType AS ENUM (
    'Barbell',
    'Dumbbell',
    'Machine',
    'Bodyweight',
    'Other'
);
--
-- Completion Type Enum
CREATE TYPE CompletionEnum AS ENUM (
  'Completed',
  'Incomplete',
  'NotStarted'
);

-- Weight Unit Enum
CREATE TYPE WeightUnit AS ENUM (
    'kg',  -- Kilograms
    'lb'   -- Pounds
);

-- Create the muscle_group_enum type if it doesn't exist
CREATE TYPE MuscleGroupEnum AS ENUM (
    'Chest',
    'Back',
    'Legs',
    'Shoulders',
    'Arms',
    'Abs',
    'Cardio'
    -- Add more muscle groups as needed
);

-- Create the muscle_group_enum type if it doesn't exist
CREATE TYPE WorkoutGoalEnum AS ENUM (
    'Build Muscle',
    'Lose Weight',
    'Improve Endurance',
    'Maintain Fitness',
    'Tone Body',
    'Custom' -- Add a "Custom" option for user-defined goals
);

-- Define the "Difficulty" enum type
CREATE TYPE Difficulty AS ENUM (
    'Beginner',
    'Intermediate',
    'Advanced'
);

-- Define the "Visibility" enum type
CREATE TYPE Visibility AS ENUM (
    'Public',
    'Private'
);

-- Define the "Rating" enum type
CREATE TYPE Rating AS ENUM (
    '1', '2', '3', '4', '5'
);

-- Define the "FatigueLevel" enum type
CREATE TYPE FatigueLevel AS ENUM (
    'Very Light',
    'Light',
    'Moderate',
    'Heavy',
    'Very Heavy'
);

-- User Table
CREATE TABLE IF NOT EXISTS "users" (
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    password_changed_at timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00z',
    "created_at" timestamptz NOT NULL DEFAULT (now()),

  -- Add a unique constraint to the email column
    CONSTRAINT unique_email UNIQUE (email)
);

-- UserProfile Table
CREATE TABLE IF NOT EXISTS UserProfile (
    user_profile_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES users(username) UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL,
    height_cm INT NOT NULL,
    height_ft_in VARCHAR(20) NOT NULL,
    preferred_unit WeightUnit NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- AvailableWorkoutPlans Table (catalog of workout plans for users to choose from)
CREATE TABLE IF NOT EXISTS AvailableWorkoutPlans (
    plan_id BIGSERIAL PRIMARY KEY,
    plan_name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL DEFAULT(''),
    goal WorkoutGoalEnum NOT NULL DEFAULT('Lose Weight'),
    difficulty Difficulty NOT NULL DEFAULT('Beginner'),
    is_public Visibility NOT NULL DEFAULT('Private'),
    created_at timestamptz NOT NULL DEFAULT(now()),
    updated_at timestamptz NOT NULL DEFAULT(now()),
    creator_username VARCHAR(255) REFERENCES users(username) NOT NULL
    -- Other metadata or information about the plan
    -- ...
);

-- WorkoutPlan Table
CREATE TABLE IF NOT EXISTS WorkoutPlan (
    plan_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES users(username) NOT NULL,
    plan_name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT(''),
    start_date timestamptz NOT NULL DEFAULT(now()),
    end_date timestamptz NOT NULL DEFAULT(now()),
    goal WorkoutGoalEnum NOT NULL DEFAULT('Lose Weight'), -- Fitness goal (e.g., build muscle, lose weight)
    difficulty Difficulty NOT NULL DEFAULT('Beginner'), -- Difficulty level (e.g., beginner, intermediate, advanced)
    is_public Visibility NOT NULL DEFAULT('Public'), -- Whether the plan is public
    "created_at" timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT unique_workout_plan_per_username UNIQUE (plan_name, username)
);
-- Workout Table
CREATE TABLE IF NOT EXISTS Workout (
    workout_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES users(username) NOT NULL,
    workout_date timestamptz NOT NULL DEFAULT (now()),
    workout_duration VARCHAR(8) NOT NULL DEFAULT (''),
    fatigue_level FatigueLevel NOT NULL DEFAULT('Very Light'), -- User's fatigue level (e.g., 1 to 5)
    notes VARCHAR(255) NOT NULL DEFAULT(''),
    total_calories_burned INT NOT NULL DEFAULT(0),
    total_distance INT NOT NULL DEFAULT(0), -- Distance covered in the workout
    total_repetitions INT NOT NULL DEFAULT(0), -- Total repetitions completed in all exercises
    total_sets INT NOT NULL DEFAULT(0), -- Total sets completed in all exercises
    total_weight_lifted INT NOT NULL DEFAULT(0), -- Total weight lifted in all exercises
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Exercise Table
CREATE TABLE IF NOT EXISTS Exercise (
    exercise_name VARCHAR(255) NOT NULL PRIMARY KEY,
    equipment_required EquipmentType NOT NULL DEFAULT('Barbell'),
    description VARCHAR(255) NOT NULL DEFAULT (''),
    muscle_group_name MuscleGroupEnum NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),

    CONSTRAINT unique_equipment_type_per_exercise UNIQUE (equipment_required, exercise_name)
);

-- AvailablePlanExercises Table (associating exercises with available workout plans)
CREATE TABLE IF NOT EXISTS AvailablePlanExercises (
    id BIGSERIAL PRIMARY KEY,
    plan_id BIGINT REFERENCES AvailableWorkoutPlans(plan_id) NOT NULL,
    exercise_name VARCHAR(255) REFERENCES Exercise(exercise_name) NOT NULL,
    sets INT NOT NULL NOT NULL DEFAULT(0),
    rest_duration VARCHAR NOT NULL DEFAULT(''),
    notes VARCHAR NOT NULL DEFAULT('')
    -- ...
);


-- OneOffWorkoutExercise Table (junction table for exercises in one-off workouts)
CREATE TABLE IF NOT EXISTS OneOffWorkoutExercise (
    id SERIAL PRIMARY KEY,
    workout_id BIGINT REFERENCES Workout(workout_id) NOT NULL,
    exercise_name VARCHAR(255) REFERENCES Exercise(exercise_name) NOT NULL,
    description VARCHAR(255) NOT NULL DEFAULT (''),
    muscle_group_name MuscleGroupEnum NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    -- ...
    CONSTRAINT unique_exercise_per_workout UNIQUE (workout_id, exercise_name)
);

-- WorkoutLog Table (advanced workout log)
CREATE TABLE IF NOT EXISTS WorkoutLog (
    log_id BIGSERIAL PRIMARY KEY,
    username VARCHAR REFERENCES users(username) NOT NULL,
    plan_id BIGINT REFERENCES WorkoutPlan(plan_id) NOT NULL,
    log_date timestamptz NOT NULL DEFAULT(now()),
    rating Rating NOT NULL DEFAULT('1'),
    fatigue_level FatigueLevel NOT NULL DEFAULT('Very Light'), -- User's fatigue level (e.g., 1 to 5)
    overall_feeling TEXT NOT NULL DEFAULT(''), -- User's overall feeling (e.g., great, tired)
    comments TEXT NOT NULL DEFAULT(''), -- General workout comments
    -- ...

    -- Additional fields for tracking overall workout performance
    workout_duration VARCHAR(10) NOT NULL DEFAULT('0m'), -- Duration of the entire workout
    total_calories_burned INT NOT NULL DEFAULT(0),
    total_distance INT NOT NULL DEFAULT(0), -- Distance covered in the workout
    total_repetitions INT NOT NULL DEFAULT(0), -- Total repetitions completed in all exercises
    total_sets INT NOT NULL DEFAULT(0), -- Total sets completed in all exercises
    total_weight_lifted INT NOT NULL DEFAULT(0), -- Total weight lifted in all exercises
    -- ...
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- ExerciseLog Table (advanced exercise log)
CREATE TABLE IF NOT EXISTS ExerciseLog (
    exercise_log_id BIGSERIAL PRIMARY KEY,
    log_id BIGINT REFERENCES WorkoutLog(log_id) NOT NULL,
    exercise_name VARCHAR(255) REFERENCES Exercise(exercise_name) NOT NULL,
    sets_completed INT NOT NULL DEFAULT(0),
    repetitions_completed INT NOT NULL DEFAULT(0),
    weight_lifted INT NOT NULL DEFAULT(0),
    notes VARCHAR(255) NOT NULL DEFAULT(''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- ExerciseSet Table (to track sets for individual exercises)
CREATE TABLE IF NOT EXISTS ExerciseSet (
    set_id BIGSERIAL PRIMARY KEY,
    exercise_log_id BIGINT REFERENCES ExerciseLog(exercise_log_id) NOT NULL,
    set_number INT NOT NULL,
    weight_lifted INT NOT NULL DEFAULT(0),
    repetitions_completed INT NOT NULL DEFAULT(0)
    -- ...
);

-- Set Table
CREATE TABLE IF NOT EXISTS Set (
    set_id BIGSERIAL PRIMARY KEY,
    exercise_name VARCHAR(255) REFERENCES Exercise(exercise_name) NOT NULL,
    set_number INT NOT NULL DEFAULT (1),
    weight int NOT NULL DEFAULT (1),
    rest_duration VARCHAR(8) NOT NULL DEFAULT (''),
    reps_completed INT NOT NULL,
    notes VARCHAR(255) NOT NULL DEFAULT(''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);


-- WeightEntry Table
CREATE TABLE IF NOT EXISTS WeightEntry (
    weight_entry_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES users(username) NOT NULL,
    entry_date timestamptz NOT NULL DEFAULT(now()),
    weight_kg INT NOT NULL DEFAULT(0),
    weight_lb INT NOT NULL DEFAULT(0),
    notes VARCHAR(255) NOT NULL DEFAULT (''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- MaxRepGoal Table
CREATE TABLE IF NOT EXISTS MaxRepGoal (
    goal_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES users(username) NOT NULL,
    exercise_name VARCHAR(255) REFERENCES Exercise(exercise_name) UNIQUE NOT NULL,
    goal_reps INT NOT NULL,
    notes VARCHAR(255) NOT NULL DEFAULT (''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- MaxWeightGoal Table
CREATE TABLE IF NOT EXISTS MaxWeightGoal (
    goal_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) REFERENCES users(username) NOT NULL,
    exercise_name VARCHAR(255) REFERENCES Exercise(exercise_name) UNIQUE NOT NULL,
    goal_weight INT NOT NULL,
    notes VARCHAR(255) NOT NULL DEFAULT (''),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES users ("username");
