# อัปเดตเวอร์ชัน Go ให้ตรงตามที่ go.mod ต้องการ
FROM golang:1.24.3

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
# คำสั่งนี้จะทำงานสำเร็จเมื่อเวอร์ชัน Go ถูกต้อง
RUN go mod download
COPY *.go ./

# สร้าง binary file
RUN go build -o /go-image-remover

EXPOSE 8080
CMD [ "/go-image-remover" ]