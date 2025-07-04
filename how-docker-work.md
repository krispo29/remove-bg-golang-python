# การทำงานของ Docker ในโปรเจกต์นี้ (ฉบับร้านอาหารมหัศจรรย์)

โปรเจกต์นี้เปรียบเสมือนร้านอาหารเปิดใหม่ที่มีเมนูเด็ดคือ "การลบพื้นหลังออกจากรูปภาพ" เพื่อให้เข้าใจง่ายขึ้น เรามาดูกันว่าแต่ละส่วนทำหน้าที่อะไร

### องค์ประกอบของร้านอาหาร

ในร้านนี้ มีพนักงาน 3 ส่วนหลักที่ทำงานร่วมกัน:

1.  **`frontend` (หน้าร้านและเมนู)**
    *   นี่คือส่วนที่ลูกค้า (ผู้ใช้งาน) มองเห็นและโต้ตอบด้วย เป็นหน้าตาของร้านที่มีเมนูให้เลือก
    *   ในโปรเจกต์นี้ มันคือหน้าเว็บไซต์ที่ให้ผู้ใช้กด "อัปโหลดรูปภาพ"

2.  **`go-backend` (พนักงานต้อนรับและรับออเดอร์)**
    *   พนักงานคนนี้คอยรับออเดอร์จากลูกค้า เมื่อลูกค้าอัปโหลดรูปภาพเข้ามา พนักงานคนนี้จะรับไฟล์รูปนั้นไว้ แล้วนำไปส่งต่อให้พ่อครัวในครัว
    *   ในโปรเจกต์นี้ คือเซิร์ฟเวอร์ที่เขียนด้วยภาษา Go ซึ่งทำหน้าที่รับคำขอ (request) จากหน้าเว็บ

3.  **`python-processor` (พ่อครัววิเศษในครัว)**
    *   นี่คือผู้เชี่ยวชาญของร้าน พ่อครัวคนนี้มีความสามารถพิเศษในการลบพื้นหลังรูปภาพ
    *   เมื่อพนักงานรับออเดอร์นำรูปมาให้ พ่อครัวจะใช้โปรแกรมภาษา Python เพื่อประมวลผลรูปภาพนั้น แล้วส่งผลลัพธ์กลับไปให้พนักงานต้อนรับเพื่อนำไปเสิร์ฟลูกค้าต่อไป

### Docker ทำหน้าที่อะไร?

**Docker คือ "ตู้คอนเทนเนอร์วิเศษ" ที่ใช้สร้างร้านอาหารของเรา**

แทนที่จะสร้างอาคารใหญ่ๆ หลังเดียว เราใช้ Docker เพื่อสร้าง "ห้องส่วนตัว" (Container) แยกสำหรับแต่ละแผนก:

*   **Container 1:** ห้องสำหรับ `frontend` (หน้าร้าน)
*   **Container 2:** ห้องสำหรับ `go-backend` (พนักงานรับออเดอร์)
*   **Container 3:** ห้องสำหรับ `python-processor` (ห้องครัว)

แต่ละห้องจะมีอุปกรณ์และเครื่องมือทุกอย่างที่จำเป็นสำหรับการทำงานของตัวเองครบถ้วน ทำให้ไม่เกิดความสับสนและไม่รบกวนการทำงานของแผนกอื่น

### `docker-compose.yml` คือ "คู่มือการประกอบร้าน"

ไฟล์นี้เป็นเหมือนพิมพ์เขียวที่บอก Docker ว่า:
1.  ต้องสร้างทั้งหมด 3 ห้อง (3 containers) ตามที่ระบุไว้
2.  ห้องของ "พนักงานรับออเดอร์" และ "พ่อครัว" จะต้องมีช่องทางสื่อสารพิเศษ (network) เพื่อให้ส่งงานกันได้
3.  ลูกค้าสามารถติดต่อร้านได้ผ่าน "ประตูหน้า" เท่านั้น (port 8080)

ด้วยวิธีนี้ เราสามารถยก "ร้านอาหารคอนเทนเนอร์" ทั้งชุดนี้ไปวางที่ไหนก็ได้ (รันบนคอมพิวเตอร์เครื่องใดก็ได้) และมันจะทำงานได้อย่างสมบูรณ์แบบเหมือนเดิมทุกครั้ง