# Supabase JWT

Create a Supabase JWT for testing FTC applications. Currently it only accepts email/password login.

## Configuration

Update the ```configs/config.yml``` file at the root of this repository. You'll need to enter:

```APP_NAME``` - The name of the FTC project (i.e. BeBooked)
```SUPABASE_PROJECT_URL``` - The Supbase project URL found in the Supabase Dashboard
```SUPABASE_API_KEY``` - The Supabase service role api key (not to be confused with public)

Example:

```yaml
APP_NAME: BeBooked
SUPABASE_PROJECT_URL: https://ldctpuxqdmlqvneoirss.supabase.co
SUPABASE_API_KEY: <api-key-here>
```

## How it works

Using the Supabase project url and api key, this application will attempt to generate a valid JWT for testing the API via Postman or anything else.
You will be prompted to enter the ```email``` and ```password``` of user. If the email/password combination is incorrect it will not
produce a valid JWT.


## How to run

```bash
./executables/apple-silicon/supabase-jwt.ftc.exe
```

or

```powershell
./executables/windows-86/supabase-jwt.ftc
```

Follow on-screen instructions

## How to compile and run for Windows

You can follow these steps to compile and execute for Windows.
Remember to change ```appName``` to the name of the app you're
compiling for.

Note: You ```MUST``` have the latest version of Go installed

1. Ensure you've updated the config.yml file (follow steps above)
2. go mod tidy
3. Compile:
```powershell
go build -o executables/windows-x86/supabase-jwt.ftc.exe ./cmd/supabase-jwt 
```
4. Execute:
```powershell
./executables/windows-x86/supabase-jwt.ftc.exe
```