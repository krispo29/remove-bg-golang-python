from flask import Flask, request, send_file
from rembg import remove
from PIL import Image
import io

app = Flask(__name__)

@app.route('/remove-bg', methods=['POST'])
def remove_bg():
    if 'image' not in request.files:
        return 'No image file provided', 400

    file = request.files['image']
    
    # Ensure it's a valid image file
    if file.filename == '':
        return 'No selected file', 400

    input_bytes = file.read()
    output_bytes = remove(input_bytes)
    
    return send_file(
        io.BytesIO(output_bytes),
        mimetype='image/png',
        as_attachment=True,
        download_name='no-bg.png'
    )

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)