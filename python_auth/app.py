from flask import Flask, request, jsonify
import requests

app = Flask(__name__)

OPA_URL = "http://localhost:8181/v1/data/authz/allow"


@app.route("/api/protected")
def protected():
    token = request.headers.get("Authorization")
    if not token:
        return jsonify({"error": "Missing token"}), 401

    payload = {"input": {"token": token}}
    response = requests.post(OPA_URL, json=payload)

    if response.json()["result"]:
        return jsonify({"message": "Access granted"})
    else:
        return jsonify({"error": "Access denied"}), 403


if __name__ == "__main__":
    app.run(debug=True)
