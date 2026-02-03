from flask import Flask, jsonify, request, redirect, render_template
import requests

app = Flask(__name__, template_folder='templates')


@app.route('/')
def home():
    return render_template('index.html')


@app.route('/send-url', methods=['POST'])
def send_url():
    Url = request.form.get('url')
    Tag = request.form.get('tag')
    data = {"Url": Url, "Tag": Tag}
    response = requests.post('http://localhost:3000/shorten', json=data)
    return jsonify({"status": "sent", "response": response.text})




if __name__ == '__main__':
    app.run(debug=True)






