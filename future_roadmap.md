# Feature Development Roadmap

เอกสารนี้สรุปแผนการพัฒนาฟีเจอร์ใหม่สำหรับโปรเจกต์ Remove Background อ้างอิงจากไอเดียที่ได้เสนอ

## Phase 1: Core Background Enhancement

ต่อยอดจากความสามารถหลักที่มีอยู่ให้สมบูรณ์ยิ่งขึ้น

### 1. เปลี่ยนพื้นหลัง (Background Replacement)
*   **เป้าหมาย:** ให้ผู้ใช้สามารถเปลี่ยนพื้นหลังที่ถูกลบออกเป็นอย่างอื่นได้
*   **รายละเอียด:**
    *   [ ] **Frontend:** เพิ่ม UI สำหรับเลือกสี (Color Picker)
    *   [ ] **Frontend:** เพิ่ม UI สำหรับอัปโหลดรูปภาพพื้นหลังใหม่
    *   [ ] **Go Backend:** แก้ไข API ให้รับข้อมูล "พื้นหลังใหม่" (ค่าสี หรือไฟล์รูปภาพ)
    *   [ ] **Python Processor:** เพิ่ม Logic ในการนำภาพที่ลบพื้นหลังแล้วมาวางซ้อนบนพื้นหลังใหม่

## Phase 2: Advanced Editing & User Experience

เพิ่มเครื่องมือแก้ไขที่ซับซ้อนขึ้นและปรับปรุงประสบการณ์การใช้งาน

### 2. เลือกวัตถุที่จะคงไว้ (Object Selection)
*   **เป้าหมาย:** ให้ผู้ใช้สามารถเลือกวัตถุที่ต้องการเก็บไว้ในภาพได้ แทนที่จะเป็นการลบพื้นหลังเพียงอย่างเดียว
*   **รายละเอียด:**
    *   [ ] **Frontend:** สร้าง Interactive Canvas ให้ผู้ใช้สามารถ "ระบาย" หรือ "คลิก" เลือกวัตถุได้
    *   [ ] **Go Backend:** แก้ไข API ให้รับข้อมูล Mask (พื้นที่ที่ผู้ใช้เลือก) ส่งต่อไปยัง Processor
    *   [ ] **Python Processor:** เปลี่ยนไปใช้ AI Model ที่รองรับ Interactive Segmentation หรือ Object Detection

### 3. ประมวลผลแบบหลายไฟล์ (Batch Processing)
*   **เป้าหมาย:** รองรับการอัปโหลดและประมวลผลรูปภาพทีละหลายไฟล์
*   **รายละเอียด:**
    *   [ ] **Frontend:** แก้ไข UI ให้รองรับการเลือกหลายไฟล์ หรือ Drag-and-Drop ทั้งโฟลเดอร์
    *   [ ] **Frontend:** แสดงสถานะการประมวลผลของแต่ละไฟล์
    *   [ ] **Go Backend:** ออกแบบและติดตั้งระบบคิว (Queue) เช่น Redis หรือ RabbitMQ เพื่อจัดการลำดับงาน
    *   [ ] **Go Backend:** API สำหรับการอัปโหลดเป็นไฟล์ .zip และดาวน์โหลดผลลัพธ์เป็น .zip

### 4. เครื่องมือแก้ไขภาพเบื้องต้น (Simple Editor)
*   **เป้าหมาย:** เพิ่มเครื่องมือแก้ไขพื้นฐานหลังจากลบพื้นหลังเสร็จ
*   **รายละเอียด:**
    *   [ ] **Frontend:** เพิ่ม UI สำหรับ Crop, Resize, Add Shadow, Add Outline
    *   [ ] **Go Backend:** สร้าง API Endpoints ใหม่สำหรับแต่ละคำสั่ง (`/api/crop`, `/api/resize` เป็นต้น)
    *   [ ] **Python Processor:** เพิ่มฟังก์ชันสำหรับจัดการรูปภาพตามคำสั่งที่ได้รับ

## Phase 3: User-Centric Features

เพิ่มฟีเจอร์ที่ต้องมีการเก็บข้อมูลผู้ใช้

### 5. ดูประวัติการทำงาน (History)
*   **เป้าหมาย:** สร้างระบบสมาชิกเพื่อให้ผู้ใช้สามารถดูและดาวน์โหลดรูปภาพที่เคยทำไว้ได้
*   **รายละเอียด:**
    *   [ ] **Infrastructure:** ติดตั้งและตั้งค่าฐานข้อมูล (เช่น PostgreSQL, MongoDB)
    *   [ ] **Go Backend:** สร้าง API สำหรับการสมัครสมาชิก (Register), ล็อกอิน (Login), และการดึงข้อมูลประวัติ (Get History)
    *   [ ] **Frontend:** สร้างหน้า UI สำหรับ Login, Register, และหน้า Profile สำหรับแสดงประวัติ
