from flask import Flask, request, send_file
from rembg import remove
from PIL import Image
import io
import requests # เพิ่มเข้ามาเพื่อดึงรูปจาก URL

app = Flask(__name__)

# ฟังก์ชันสำหรับรวมภาพ
def composite_images(fg_bytes, bg_bytes):
    foreground = Image.open(io.BytesIO(fg_bytes)).convert("RGBA")
    background = Image.open(io.BytesIO(bg_bytes)).convert("RGBA")

    # ปรับขนาดพื้นหลังให้เท่ากับภาพหลัก
    background = background.resize(foreground.size)
    
    # รวมภาพ
    background.paste(foreground, (0, 0), foreground)
    
    output_buffer = io.BytesIO()
    background.save(output_buffer, format='PNG')
    output_buffer.seek(0)
    return output_buffer

# ฟังก์ชันสำหรับใส่สีพื้นหลัง
def add_color_background(fg_bytes, bg_color):
    foreground = Image.open(io.BytesIO(fg_bytes)).convert("RGBA")
    
    # สร้างพื้นหลังสี
    background = Image.new('RGBA', foreground.size, bg_color)
    
    # รวมภาพ
    background.paste(foreground, (0, 0), foreground)

    output_buffer = io.BytesIO()
    background.save(output_buffer, format='PNG')
    output_buffer.seek(0)
    return output_buffer

@app.route('/remove-bg', methods=['POST'])
def remove_bg_endpoint():
    if 'image_fg' not in request.files:
        return 'No foreground image provided', 400

    # รับไฟล์ภาพหลัก
    fg_file = request.files['image_fg']
    fg_bytes = fg_file.read()

    # ลบพื้นหลังก่อน
    processed_fg_bytes = remove(fg_bytes)

    # --- ตรวจสอบเงื่อนไข ---
    # 1. ถ้ามีไฟล์พื้นหลังถูกส่งมาด้วย
    if 'image_bg' in request.files:
        bg_file = request.files['image_bg']
        if bg_file.filename != '':
            bg_bytes = bg_file.read()
            final_image_buffer = composite_images(processed_fg_bytes, bg_bytes)
            return send_file(final_image_buffer, mimetype='image/png')

    # 2. ถ้ามีโค้ดสีส่งมาด้วย
    bg_color = request.form.get('bg_color')
    if bg_color:
        final_image_buffer = add_color_background(processed_fg_bytes, bg_color)
        return send_file(final_image_buffer, mimetype='image/png')
    
    # 3. ถ้าไม่มีทั้งสองอย่าง (ลบพื้นหลังอย่างเดียว)
    return send_file(io.BytesIO(processed_fg_bytes), mimetype='image/png')