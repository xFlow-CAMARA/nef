from flask import Flask, request, jsonify
import logging


app = Flask(__name__)

@app.route('/callbacks', methods=['POST'])
def callbacks():
    data = request.get_json()
    if data:
        print("Received JSON:", data)
        return jsonify({"message": "Callback received"}), 200
    else:
        return jsonify({"error": "Invalid JSON"}), 400

if __name__ == '__main__':
    app.run(debug=True, port=5050)