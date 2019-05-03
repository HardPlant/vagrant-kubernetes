SET IMAGE=hp_build_bot
SET VERSION=1.0.0.3
SET DEPLOY_FILE=deployments.yml
powershell -Command "(gc %DEPLOY_FILE%) -replace 'image: %IMAGE%', 'image: %IMAGE%:%VERSION% #' | Out-File %DEPLOY_FILE%"
REM docker build -t %IMAGE%:%VERSION% .
REM kubectl apply -f deployments.yml