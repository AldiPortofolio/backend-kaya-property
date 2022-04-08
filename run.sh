#main:
export KAYA_API_SERVICE_PORT=0.0.0.0:5000

#token:
export SECRET_KEY="jwye99wbjskayafkjkdjserviceskdjkdrnmwrt"
export REFRESH_SECRET_KEY="jwye99wbjskayafkjkdjserviceskdsmdnssnrt"

#redis:
export REDIS_ENDPOINT="127.0.0.1:6379"

#routes/route:
export KAYA_API_SERVICE="KAYA_API_SERVICE"
export READ_TIMEOUT="120"
export WRITE_TIMEOUT="120"

#database/postgres:
#export DB_POSTGRES_USER="postgres"
#export DB_POSTGRES_PASS="admin123"
#export DB_POSTGRES_NAME="kaya"
#export DB_POSTGRES_HOST="147.139.192.236"
#export DB_POSTGRES_PORT="5432"
#export DB_POSTGRES_SSLMODE="disable"
#export DB_POSTGRES_TIMEOUT="5"

export DB_POSTGRES_USER="postgres"
export DB_POSTGRES_PASS="admin"
export DB_POSTGRES_NAME="kaya"
export DB_POSTGRES_HOST="localhost"
export DB_POSTGRES_PORT="5432"
export DB_POSTGRES_SSLMODE="disable"
export DB_POSTGRES_TIMEOUT="5"

#swagger:
export SWAGGER_HOST_LOCAL="0.0.0.0:4000"
export SWAGGER_HOST_DEV="0.0.0.0:4000"

#minio:
export MINIO_ENDPOINT="127.0.0.1:9000"
export MINIO_ACCESSKEY="AKIAIOSFODNN7KAYAACCESSKEY"
export MINIO_SECRETKEY="wJalrXUtnFEMI/K7MDENG/bPxRfiCYKAYASECRETKEY"
export MINIO_BUCKET="kaya-dev"
export MINIO_FILE_PATH="http://127.0.0.1:9000"

#zenziva:
export ZENZIVA_ENDPOINT="https://console.zenziva.net/wareguler/api/sendWA/"
export ZENZIVA_USER_KEY="d3d73f1714d3"
export ZENZIVA_PASS_KEY="0581b880f3373945c851e954"
export ZENZIVA_ON="1"

#upload directory:
export UPLOAD_DIRECTORY="files"

#email:
export SMTP_ADDRESSS="smtp-relay.gmail.com"
export SMTP_PORT="587"
export EMAIL_SENDER="hello@kaya.co.id"
export EMAIL_PASSWORD="kayaProperty123?"

#endpoint web:
export ENDPOINT_RESET_PASSWORD="https://kaya.co.id/reset-password/reset"

go run main.go

# nohup ./kaya-backend > nohup.out 2>&1 & echo $! > run.pid